package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/lib/pq"

	"github.com/ramil600/sensors2/business/user/db"
)

func TestCreate(t *testing.T) {

	usr := db.User{
		Name:         "Ramil Mirnov",
		Email:        "s@s1.com",
		Roles:        pq.StringArray([]string{"admin"}),
		PasswordHash: []byte("hello"),
		DateCreated:  time.Now(),
		DateUpdated:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dbconn, err := db.Open(db.DBcfg)
	if err != nil {
		t.Fatal(err)
	}
	store := db.NewStore(dbconn)
	err = store.Create(ctx, usr)
	if err != nil {
		t.Fatal(err)
	}

}
