package infraestructure

import (
	"log"

	"example.com/m/Users/domain"
	"example.com/m/core"
)

type UserRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewUserRepoMySQL() *UserRepoMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return &UserRepoMySQL{Connection: *conn}
}

func (r *UserRepoMySQL) SaveUserFunction(user domain.User)error{
	query := "INSERT INTO User (username, email, password) VALUES (?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Username, user.Email, user.Password)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
    return err
}

func (r *UserRepoMySQL) GetUserFunction(id int32)([]domain.User, error){
	query := "SELECT id, username, email FROM User WHERE id = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var users []domain.User
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int32
		var username string
		var email string

		if err := rows.Scan(&id, &username, &email); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		user := domain.User{ID: id, Username: username, Email: email, Password: ""}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepoMySQL) EditUserFunction(user domain.User)error{
	query := "UPDATE User SET username = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
        log.Fatalf("Error al editar info. dl usuario:", err)
    }
    return err
}

func (r *UserRepoMySQL) DeleteUserFunction(id int32)error{
	_, err := r.Connection.ExecPreparedQuerys("DELETE FROM User WHERE id = ?", id)
	if err != nil {
        log.Fatalf("Error al eliminar usuario:", err)
    }
	return err
}

func (r *UserRepoMySQL) ComprobateQuantityOfUsers()(int, error){
	result, err := r.Connection.FetchRows("SELECT * FROM User")
	var quantity_of_users int
	if err != nil {
		log.Fatalf("Error al obtener cantidad de usuarios: ", err)
	}
	defer result.Close()
	for result.Next() {
		quantity_of_users++
	}
	return quantity_of_users, err
}