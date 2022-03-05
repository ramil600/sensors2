package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/lib/pq"
)

var usr User
var store Store

func TestMain(m *testing.M) {

	dbconn, err := Open(DBcfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()
	store = NewStore(dbconn)

	os.Exit(m.Run())

}
func TestCreate(t *testing.T) {

	usr = User{
		Name:         "Ramil Mirnov",
		Email:        "s@s1.com",
		Roles:        pq.StringArray([]string{"admin"}),
		PasswordHash: []byte("hello"),
		DateCreated:  time.Now(),
		DateUpdated:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := store.Create(ctx, usr)
	if err != nil {
		t.Fatal(err)
	}

}

func TestQuery(t *testing.T) {

	q := `SELECT user_id from users WHERE email=$1`

	err := store.DB.Get(&usr, q, usr.Email)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	usr, err := store.Query(ctx, usr.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("user found by query", usr.Name, usr.Email)

	if usr.Name == "" {
		t.Error("User not found")
	}

}
