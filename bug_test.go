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
	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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
	client = client.Debug()
	ctx := context.Background()

	// driver.Query: query=INSERT INTO `users` (`age`, `name`) VALUES (?, ?) RETURNING `id` args=[30 Ariel]
	u := client.User.Create().SetName("Ariel").SetAge(30).SaveX(ctx)

	// driver.Query: query=SELECT COUNT(DISTINCT `users`.`id`) FROM `users` args=[]
	if n := client.User.Query().CountX(ctx); n != 1 {
		t.Errorf("unexpected number of users: %d", n)
	}

	// driver.Exec: query=UPDATE `users` SET `deleted_time` = ? WHERE `users`.`id` IN (?) args=[2022-08-13 20:36:57.943277 -0300 -03 m=+0.007749764 1]
	client.User.DeleteOne(u).ExecX(ctx)

	// do a real delete op
	// query=DELETE FROM `users` args=[]
	_, err := client.User.Delete().Exec(schema.WithSkipDeletedTimeHook(ctx))
	if err != nil {
		t.Errorf("could not delete the users: %v", err)
	}

	// query=INSERT INTO `users` (`age`, `name`) VALUES (?, ?) RETURNING `id` args=[30 Ariel]
	client.User.Create().SetName("Ariel").SetAge(30).SaveX(ctx)

	// query=INSERT INTO `users` (`age`, `name`) VALUES (?, ?) RETURNING `id` args=[28 Pedro]
	client.User.Create().SetName("Pedro").SetAge(28).SaveX(ctx)

	// Delete many
	// query=SELECT `users`.`id` FROM `users` args=[]
	// query=UPDATE `users` SET `deleted_time` = ? WHERE `users`.`id` IN (?, ?) args=[2022-08-13 20:42:34.613814 -0300 -03 m=+0.009210090 2 3]
	client.User.Delete().ExecX(ctx)
}
