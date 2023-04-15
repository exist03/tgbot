package mysql

import "database/sql"

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(id int64, title, content string) error {
	stmt := `INSERT INTO snippets (id, title, content) VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, id, title, content)
	if err != nil {
		return err
	}
	return nil
}
