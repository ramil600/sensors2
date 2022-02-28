package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/lib/pq"

	"github.com/ramil600/sensors2/business/user/db"
)

func TestCreate(t *testing.T) {

	nu := db.NewUser{
		Name:         "Ramil Mirhasanov",
		Email:        "s@s.com",
		Roles:        pq.StringArray([]string{"admin"}),
		PasswordHash: "hello",
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
	err = store.Create(ctx, nu)
	if err != nil {
		t.Fatal(err)
	}

}
