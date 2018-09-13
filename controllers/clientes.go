package controllers

import (
  "github.com/gin-gonic/gin"
  "sait.mx/MVC/model"
  "io/ioutil"
  "encoding/json"
)



func ListClientes(c *gin.Context){
  clientes, err := model.ListClientes()
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  c.JSON(200, clientes)

}

func GetCliente(c *gin.Context){
  id := c.Param("id")
  cliente, err := model.GetCliente(id)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  c.JSON(200, cliente)
}


func InsertCliente(c *gin.Context){
  body, err := ioutil.ReadAll(c.Request.Body)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  var cliente model.Cliente
  err = json.Unmarshal(body, &cliente)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  err = model.InsertCliente(cliente)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  c.JSON(200, "Insertado correctamente")
}


func UpdateCliente(c *gin.Context){
  id := c.Param("id")
  body, err := ioutil.ReadAll(c.Request.Body)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  var cliente model.Cliente
  err = json.Unmarshal(body, &cliente)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  cliente.ID = id
  err = model.UpdateCliente(cliente)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  c.JSON(200, "actualizado correctamente")
}

func DeleteCliente(c *gin.Context){
  id := c.Param("id")
  err := model.DeleteCliente(id)
  if err!= nil {
    c.JSON(500, err.Error())
    return
  }
  c.JSON(200, "eliminado correctamente")
}
