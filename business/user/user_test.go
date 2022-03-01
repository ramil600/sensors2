package user

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ramil600/sensors2/business/user/db"
	"golang.org/x/crypto/bcrypt"
)

var core Core

func TestMain(m *testing.M) {
	dbc, err := db.Open(db.DBcfg)
	if err != nil {
		log.Fatal(err)
	}
	core = NewCore(dbc)
	os.Exit(m.Run())
}

func TestCoreCreate(t *testing.T) {

	now := time.Now()

	nu := NewUser{
		Name:            "ramil",
		Email:           "ramil@yahoo.com",
		Roles:           []string{"admin"},
		Password:        "ramil",
		PasswordConfirm: "ramil",
	}

	dbUsr, err := core.Create(context.TODO(), nu, now)
	if err != nil {
		t.Fatal(err)
	}

	if bcrypt.CompareHashAndPassword(dbUsr.PasswordHash, []byte(nu.Password)) != nil {
		t.Log(string(dbUsr.PasswordHash))
		t.Log("Password hash is not what is expected")
		t.Fail()
	}

	if dbUsr.DateCreated != now || dbUsr.DateUpdated != now {
		t.Log("Date created or updated mismatch")
		t.Fail()
	}

}
