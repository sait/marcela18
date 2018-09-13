package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cliente struct {
	ID       string `db:"id"          json:"id"`      // ETIQUETAS(campo) DE LA BASE DE DATOS id
	Nombre   string `db:"nombre"      json:"nombre"`
	Apellido string `db:"apellido"    json:"apellido"`
}

var DB *sqlx.DB // declara de manera global aqui se generaran todas las operaciones a la BD


  func OpenDB(){
    db, err := sqlx.Open("mysql", "root:@tcp(localhost:3307)/pruebasgo")
    if err != nil{

      fmt.Println(err)
      return
    }

    DB = db
    err = DB.Ping()
    if err != nil {
      fmt.Println(err)
      return
    }
}



func InsertCliente(cliente Cliente) (err error){
	_, err = DB.NamedExec(`insert into clientes (nombre, apellido)
		VALUES (:nombre, :apellido)`, &cliente)
	return
}

func DeleteCliente(clienteID string) (err error){
	statement, err := DB.Prepare(`DELETE  FROM clientes
	WHERE id=?`) // EL ? ES EL PARAMETRO ENVIADO
	statement.Exec(clienteID)

	return
}
/*func UpdateCliente(clienteID string, nombreID string, apellidoID string) (err error){
	statement, err := DB.Prepare(`UPDATE clientes SET nombre=?, apellido=? WHERE id=?`) // EL ? ES EL PARAMETRO ENVIADO
	statement.Exec(clienteID, nombreID, apellidoID)
	return
}*/

func UpdateCliente(cliente Cliente) (err error){
	_, err = DB.NamedExec(`UPDATE clientes SET nombre:=nombre, apellido=:apellido WHERE id:=id`, cliente)
	return
}

func GetCliente(clientID string) (cliente Cliente, err error) {
	err = DB.Get(&cliente, `SELECT id, nombre, apellido FROM clientes
		WHERE id=?`, clientID) // EL ? ES EL PARAMETRO ENVIADO
	return
}

func SelectCliente() (err error) {
	rows, err := DB.Query(`SELECT * FROM clientes`) // EL ? ES EL PARAMETRO ENVIADO
	for rows.Next() {
	            var id int
	            var nombre string
	            var apellido string

	            err = rows.Scan(&id, &nombre, &apellido)
	            fmt.Println(id)
	            fmt.Println(nombre)
	            fmt.Println(apellido)
	        }
	return
}

func ListClientes() (clientes []Cliente, err error){
  err = DB.Select(&clientes, `SELECT id, nombre, apellido FROM clientes`)
  return
}
