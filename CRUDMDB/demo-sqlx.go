package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cliente struct {
	ID       string `db:"id"` // ETIQUETAS(campo) DE LA BASE DE DATOS id
	Nombre   string `db:"nombre"`
	Apellido string `db:"apellido"`
}

var DB *sqlx.DB // declara de manera global aqui se generaran todas las operaciones a la BD

func main() {  // EMPIEZA MAIN
	fmt.Println("Pruebas con base de datos")
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3307)/pruebasgo")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db //asignando valor a la variable DB
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Conexion exitosa")

	// INSERT CLIENTE
	var cli Cliente
	cli.Nombre = "Marcela"
	cli.Apellido = "Dorame"
	err = InsertCliente(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
		fmt.Println("Alta Exitosa")

	// DELETE CLIENTE
	cliente, err := DeleteCliente("4")  //ENVIANDO PARAMETRO y mllamando la funcion GetCliente
	 	if err != nil {
	 		fmt.Println(err)
	 		return
	 	}
	 	fmt.Println(cliente)
	 	fmt.Println("Eliminacion Exitosa")

	// UPDATE CLIENTE
	err = UpdateCliente("Juan", "Perez", "1")  //ENVIANDO PARAMETRO y mllamando la funcion GetCliente
		 	if err != nil {
		 		fmt.Println(err)
		 		return
		 	}
		 	fmt.Println("Correccion Exitosa")

	//SELECT CLIENTE
	cliente, err = GetCliente("1")  //envio parametro y llamo funcion GetCliente
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cliente)
	fmt.Println("Consulta Exitosa")
	// SELECT * CLIENTE
	err = SelectCliente()  //envio parametro y llamo funcion SelectClient
	if err != nil {
	fmt.Println(err)
	 return
 }

 fmt.Println("Query Exitoso")

}  // TERMINA MAIN

func InsertCliente(cliente Cliente) (err error){
	_, err = DB.NamedExec(`insert into clientes (nombre, apellido)
		VALUES (:nombre, :apellido)`, &cliente)
	return
}

func DeleteCliente(clienteID string) (cliente Cliente, err error){
	statement, err := DB.Prepare(`DELETE  FROM clientes
	WHERE id=?`) // EL ? ES EL PARAMETRO ENVIADO
	statement.Exec(clienteID)

	return
}
func UpdateCliente(clienteID string, nombreID string, apellidoID string) (err error){
	statement, err := DB.Prepare(`UPDATE clientes SET nombre=?, apellido=? WHERE id=?`) // EL ? ES EL PARAMETRO ENVIADO
	statement.Exec(clienteID, nombreID, apellidoID)
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
