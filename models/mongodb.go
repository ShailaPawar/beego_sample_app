package models

import (
  "gopkg.in/mgo.v2"
  "github.com/astaxie/beego"
)
const kPreSalesDB string = "presales_huddle"
const kUsersTable string = "users"

var session *mgo.Session

func Conn() *mgo.Session {
  return session.Copy()
}

func init() {
  url := beego.AppConfig.String("mongodb::url")
  sess, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }

  session = sess
  session.SetMode(mgo.Monotonic, true)
}
