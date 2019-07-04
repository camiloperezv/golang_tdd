package dbuser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name string `json:"name"`
}

func checkName(name string) bool {
	if len(name) < 2 {
		return false
	}
	return true
}

func NewUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := gorm.Open("postgres", "postgres://sql_example@localhost:5432/sql_example?sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var user User
	raw, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(raw, &user)
	if checkName(user.Name) == false {
		s, _ := json.Marshal("{error:\"name is too short\"}")
		fmt.Fprint(w, string(s))
		return
	}
	db.Create(&user)
	s, _ := json.Marshal(user)
	fmt.Fprint(w, string(s))
}
