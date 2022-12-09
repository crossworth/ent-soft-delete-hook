package bug

import (
	"context"
	"testing"

	"entgo.io/bug/ent/enttest"
	_ "entgo.io/bug/ent/runtime"
	"entgo.io/bug/ent/schema"
	"entgo.io/bug/ent/user"
	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestNode(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()

	u := client.User.Create().SetName("Ariel").SetAge(30).SaveX(ctx)

	if n := client.User.Query().CountX(ctx); n != 1 {
		t.Errorf("unexpected number of users: %d", n)
	}

	// driver.Query: query=SELECT `users`.`id` FROM `users` WHERE `users`.`name` = ? AND `users`.`deleted_time` IS NULL LIMIT 1 args=[Ariel]
	exists := client.Debug().User.Query().Where(user.Name("Ariel")).ExistX(ctx)
	require.True(t, exists)

	// driver.Exec: query=UPDATE `users` SET `deleted_time` = ? WHERE `users`.`id` IN (?) args=[2022-12-08 14:23:35.446387627 -0300 -03 m=+0.003711003 1]
	client.Debug().User.DeleteOne(u).ExecX(ctx)

	// driver.Query: query=SELECT `users`.`id` FROM `users` WHERE `users`.`name` = ? AND `users`.`deleted_time` IS NULL LIMIT 1 args=[Ariel]
	exists = client.Debug().User.Query().Where(user.Name("Ariel")).ExistX(ctx)
	require.False(t, exists)

	// skip
	// driver.Query: query=SELECT `users`.`id` FROM `users` WHERE `users`.`name` = ? LIMIT 1 args=[Ariel]
	exists = client.Debug().User.Query().Where(user.Name("Ariel")).ExistX(schema.SkipSoftDelete(ctx))
	require.True(t, exists)

	// skip delete
	// driver.Exec: query=DELETE FROM `users` WHERE `users`.`id` = ? args=[1]
	client.Debug().User.DeleteOne(u).ExecX(schema.SkipSoftDelete(ctx))

	// driver.Query: query=SELECT `users`.`id` FROM `users` WHERE `users`.`name` = ? LIMIT 1 args=[Ariel]
	exists = client.Debug().User.Query().Where(user.Name("Ariel")).ExistX(schema.SkipSoftDelete(ctx))
	require.False(t, exists)
}

func TestO2M(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()

	u := client.User.Create().SetName("Ariel").SetAge(30).SaveX(ctx)
	td1 := client.Todo.Create().SetName("T1").SetCreator(u).SaveX(ctx)
	client.Todo.Create().SetName("T2").SetCreator(u).SaveX(ctx)

	users := client.Debug().User.Query().Where(user.Name("Ariel")).WithTodos().AllX(ctx)
	require.Len(t, users, 1)
	require.Len(t, users[0].Edges.Todos, 2)
	require.Equal(t, "T1", users[0].Edges.Todos[0].Name)
	require.Equal(t, "T2", users[0].Edges.Todos[1].Name)

	client.Debug().Todo.DeleteOne(td1).ExecX(ctx)

	users = client.Debug().User.Query().Where(user.Name("Ariel")).WithTodos().AllX(ctx)
	require.Len(t, users, 1)
	require.Len(t, users[0].Edges.Todos, 1)
	require.Equal(t, "T2", users[0].Edges.Todos[0].Name)

	// skip
	users = client.Debug().User.Query().Where(user.Name("Ariel")).WithTodos().AllX(schema.SkipSoftDelete(ctx))
	require.Len(t, users, 1)
	require.Len(t, users[0].Edges.Todos, 2)
	require.Equal(t, "T1", users[0].Edges.Todos[0].Name)
	require.Equal(t, "T2", users[0].Edges.Todos[1].Name)
}

func TestM2M(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()

	u1 := client.User.Create().SetName("u1").SetAge(10).SaveX(ctx)
	u2 := client.User.Create().SetName("u2").SetAge(20).SaveX(ctx)

	client.Group.Create().SetName("group").AddUsers(u1, u2).SaveX(ctx)

	g1 := client.Group.Query().WithUsers().OnlyX(ctx)
	require.Len(t, g1.Edges.Users, 2)
	require.Equal(t, u1.ID, g1.Edges.Users[0].ID)
	require.Equal(t, u2.ID, g1.Edges.Users[1].ID)

	client.User.DeleteOne(u1).ExecX(ctx)

	g1 = client.Group.Query().WithUsers().OnlyX(ctx)
	require.Len(t, g1.Edges.Users, 1)
	require.Equal(t, u2.ID, g1.Edges.Users[0].ID)
}
