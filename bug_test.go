package bug

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"testing"

	"entgo.io/bug/ent"
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

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
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
