package main

import (
	"database/sql"
	"fmt"
)

type GlobalValuesInstance struct {
	ID    int
	Count int
}

const GlobalValuesTableKey string = "global_values"

func InitGlobalValuesTable(dbClient *sql.DB) error {
	query := fmt.Sprintf("SELECT table_name FROM information_schema.tables where table_name = %s;", GlobalValuesTableKey)
	rows, err := dbClient.Query(query)
	if err != nil {
		fmt.Println("Teste")
		return err
	}
	defer rows.Close()

	//if !rows.Next() {
	//	query := fmt.Sprintf(`
	//        CREATE TABLE %s (
	//		id SERIAL PRIMARY KEY,
	//		count INT NOT NULL
	//		);
	//    `, GlobalValuesTableKey)
	//
	//	_, err := dbClient.Exec(query)
	//	return err
	//}
	return nil
}

func (i *GlobalValuesInstance) Create(dbClient *sql.DB) error {
	if i.ID != 0 {
		query := fmt.Sprintf(`
            INSERT INTO %s (id, count)
VALUES ($1, $2)
ON CONFLICT (id) 
DO NOTHING;
        `, GlobalValuesTableKey)
		_, err := dbClient.Exec(query, i.ID, i.Count)
		if err != nil {
			fmt.Println("Erro create")
			return err
		}
		return nil
	}

	query := fmt.Sprintf(`
    INSERT INTO %s (count)
    VALUES ($1);
    `, GlobalValuesTableKey)

	_, err := dbClient.Exec(query, i.Count)
	if err != nil {
		fmt.Println("Erro create2")
		return err
	}
	return nil
}

func (i *GlobalValuesInstance) Read(dbClient *sql.DB) error {
	query := fmt.Sprintf(`
    SELECT count FROM %s WHERE id=$1;
    `, GlobalValuesTableKey)

	err := dbClient.QueryRow(query, i.ID).Scan(&i.Count)
	if err != nil {
		fmt.Println("Erro read")
		return err
	}

	return nil
}

func (i *GlobalValuesInstance) Update(dbClient *sql.DB) error {
	i.Count++
	query := fmt.Sprintf(`
        UPDATE %s 
        SET count=$1
        WHERE id=$2;
    `, GlobalValuesTableKey)

	_, err := dbClient.Exec(query, i.Count, i.ID)
	if err != nil {
		return err
	}
	return nil
}

func (i *GlobalValuesInstance) Delete(dbClient *sql.DB) error {
	query := fmt.Sprintf(`
        DELETE FROM %s
        WHERE id=$1;
    `, GlobalValuesTableKey)

	_, err := dbClient.Exec(query, i.ID)
	if err != nil {
		return err
	}
	return nil
}
