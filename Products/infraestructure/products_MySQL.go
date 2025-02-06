package infraestructure

import (
    "log"

    "example.com/m/Products/domain"
    "example.com/m/core"
)

type ProductRepoMySQL struct {
    connection core.ConectionMySQL
}

func NewCreateProductRepoMySQL() *ProductRepoMySQL {
    conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
    return &ProductRepoMySQL{connection: *conn}
}

func (r *ProductRepoMySQL) Save(product *domain.Product) error {
    query := "INSERT INTO Product (name, price) VALUES (?, ?)"
    _, err := r.connection.ExecPreparedQuerys(query, product.Name, product.Price)
    if err != nil {
        log.Fatalf("Error al registrar producto:", err)
    }
    return err
}

func (r *ProductRepoMySQL) GetAll() ([]domain.Product,error) {
    query := "SELECT * FROM Product"
    rows, err := r.connection.FetchRows(query)
	var products []domain.Product
    if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var price float32

		
        if err := rows.Scan(&id, &name, &price); err != nil {
            log.Println("Error al escanear la fila:", err)
        }
        log.Printf("ID: %d, Name: %s, Price: %f\n", id, name, price)
		product := domain.Product{ID: int32(id), Name: name, Price: price}
		products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error al iterar sobre las filas:", err)
    }
    return products, err
}

func (r *ProductRepoMySQL) Edit(name string, price float32, id int32) error {
    query := "UPDATE Product SET name = ?, price = ? WHERE id = ?"
    _, err := r.connection.ExecPreparedQuerys(query, name, price, id)
    if err != nil {
        log.Fatalf("Error al editar info. de producto:", err)
    }
    return err
}

func (r *ProductRepoMySQL) Delete(id int32) error {
    query := "DELETE FROM Product WHERE id = ?"
    _, err := r.connection.ExecPreparedQuerys(query, id)
    if err != nil {
        log.Fatalf("Error al eliminar producto:", err)
    }
    return err
}

func (r *ProductRepoMySQL) UpdatePriceProduct(id int32, price float32) error {
    _, err := r.connection.ExecPreparedQuerys("UPDATE Product SET price = ? WHERE id = ?", price, id)
    if err != nil {
        log.Fatalf("Error al actualizar precios: ", err)
    }
    return err
}

func (r *ProductRepoMySQL)GetLastAddedProduct()(domain.Product, error){
    rows, err := r.connection.FetchRows("SELECT * FROM Product ORDER BY id DESC LIMIT 1")
    if err != nil {
        log.Fatalf("Error al obtener último producto registrado: ", err)
    }
    defer rows.Close()

    var product domain.Product

    for rows.Next() {
        var id int32
        var name string
        var price float32

        if err := rows.Scan(&id, &name, &price); err != nil {
            log.Println("Error al escanear la fila:", err)
        }
        log.Printf("ID: %d, Name: %s, Price: %f\n", id, name, price)
		new_product := &domain.Product{ID: id, Name: name, Price: price}
        product = *new_product
    }

    if err := rows.Err(); err != nil {
        log.Println("Erro al iterar sobre las filas: ", err)
    }
    return product, err
}