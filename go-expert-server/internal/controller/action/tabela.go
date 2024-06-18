package controller

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CriarTabela() {
	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Comando SQL para criar a tabela cotacoes
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS cotacoes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        valor REAL NOT NULL,
        data TEXT NOT NULL
    );`

	// Executar o comando SQL
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabela cotacoes criada com sucesso!")
}
