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
var Email = "ramil@yahoo.com"

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
		Email:           Email,
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

func TestUpdate(t *testing.T) {
	name := "Some New Name"
	roles := []string{"admin", "user"}
	pwd := "password"
	pwdConfirm := "password"

	upd := UserUpdate{
		Name:            &name,
		Roles:           roles,
		Password:        &pwd,
		PasswordConfirm: &pwdConfirm,
	}
	rows, err := core.store.DB.Query("Select user_id from users where email=$1", Email)

	if err != nil {
		t.Fatal(err)
	}
	if !rows.Next() {
		t.Fatal(rows.Err())
	}
	var user_id string
	rows.Scan(&user_id)
	dbUsr, err := core.Update(context.TODO(), upd, user_id, time.Now())

	if err != nil {
		t.Fatal(err)
	}

	if dbUsr.Name != name {
		t.Log("Name returned is not what is expected")
		t.Fatal(err)
	}

}

func TestDelete(t *testing.T) {
	rows, err := core.store.DB.Query("SELECT user_id FROM users WHERE email=$1", Email)
	if err != nil {
		t.Fatal(err)
	}
	if !rows.Next() {
		t.Fatal(rows.Err())
	}
	var userId string
	err = rows.Scan(&userId)
	if err != nil {
		t.Fatal(err)
	}
	err = core.Delete(context.TODO(), userId)
	if err != nil {
		t.Fatal(err)
	}
}
