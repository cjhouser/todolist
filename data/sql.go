package data

const (
	sqlCreateTodoItemsTable = `CREATE TABLE IF NOT EXISTS todoitems (id BIGINT PRIMARY KEY AUTO_INCREMENT, title TEXT NOT NULL)`
	sqlInsertTodoItem       = `INSERT INTO todoitems (title) VALUES(?)`
	sqlSelectAllTodoItems   = `SELECT id, title FROM todoitems`
	sqlUpdateTodoItem       = `UPDATE todoitems SET title=? WHERE id=?`
	sqlDeleteTodoItem       = `DELETE FROM todoitems WHERE id=?`
)
