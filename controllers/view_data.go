package controllers

import (
  "beego_sample_app/models"
  "github.com/astaxie/beego"
  "fmt"
)

type ViewDataController struct {
  beego.Controller
}

func (this *ViewDataController) ViewData() {
  fmt.Println("Data: ", models.GetData())

  length := len(models.GetData())
  var value string
  var values []string

  for i := 0; i < length; i++ {
    email := models.GetData()[i].Email
    role := models.GetData()[i].Role

    if i != length - 1 {
      value = email + ": " + role + ","
    } else {
      value = email + ": " + role
    }
    values = append(values, value)
  }

  this.Data["DataToView"] = values
  this.TplName = "view_data.tpl"
}
