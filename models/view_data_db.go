package models

import (
  "fmt"
)

type User struct {
  Email string `bson:"Email"`
  Role  string `bson:"Role"`
}

func GetData()  (users []User){
  mConn := Conn()
  defer mConn.Close()

  collection := mConn.DB(kPreSalesDB).C(kUsersTable)

  err := collection.Find(nil).All(&users)

  if err != nil {
    fmt.Printf("Error while fectching data from DB: ", err)
    return
  }
  return users
}
