package data

import (
	"database/sql"
	"log"

	"github.com/CJHouser/tasklist/models"
	_ "github.com/go-sql-driver/mysql"
)

// DBTodoItem embeds TodoItem and adds DB-specific methods
// https://golang.org/doc/effective_go.html#embedding
type DBTodoItem struct {
	*models.TodoItem
}

func (todoItem *DBTodoItem) scan(rows *sql.Rows) {
	err := rows.Scan(
		&todoItem.Id,
		&todoItem.Title)
	if err != nil {
		log.Fatalln(err)
	}
}

// DBTodoItems embeds DBTodoItem and adds DB-specific methods
type DBTodoItems []*DBTodoItem

func (todoItems *DBTodoItems) scan(rows *sql.Rows) error {
	for rows.Next() {
		todoItem := &DBTodoItem{&models.TodoItem{}}
		todoItem.scan(rows)
		*todoItems = append(*todoItems, todoItem)
	}
	return rows.Err()
}

// TodoDB provides methods for accessing DB data
type TodoDB struct {
	DB *sql.DB
}

// OpenDb opens a MySQL database with the specified dbname and dbuser
func (todoDb *TodoDB) OpenDb(dbname string, dbuser string) error {
	db, err := sql.Open("mysql", dbuser+":pass@/"+dbname)
	todoDb.DB = db
	return err
}

// CreateTablesIfNotExists creates any MySQL tables that do not exist
func (todoDb *TodoDB) CreateTablesIfNotExists() error {
	_, err := todoDb.DB.Exec(sqlCreateTodoItemsTable)
	return err
}

// SelectAllTodoItems returns all rows from the DB as DBTodoItems
func (todoDb *TodoDB) SelectAllTodoItems() (items *DBTodoItems, err error) {
	todoItems := &DBTodoItems{}
	rows, err := todoDb.DB.Query(sqlSelectAllTodoItems)
	if err != nil {
		return nil, err
	}
	todoItems.scan(rows)
	return todoItems, nil
}

func (todoDb *TodoDB) SelectSingleTodoItem(item *DBTodoItem) (todoItem *DBTodoItem, err error) {
	row := todoDb.DB.QueryRow(
		sqlSelectSingleTodoItem,
		item.Id)
	if err != nil {
		return nil, err
	}
	var id int
	var title string
	err = row.Scan(&id, &title)
	if err != nil {
		return nil, err
	}
	todoItem = &DBTodoItem{&models.TodoItem{Id: id, Title: title}}
	return todoItem, nil
}

// InsertTodoItem inserts a single DBTodoItem into the DB
func (todoDb *TodoDB) InsertTodoItem(item *DBTodoItem) error {
	_, err := todoDb.DB.Exec(
		sqlInsertTodoItem,
		item.Title)
	return err
}

// UpdateTodoItem updates a single DBTodoItem within the DB
func (todoDb *TodoDB) UpdateTodoItem(item *DBTodoItem) error {
	_, err := todoDb.DB.Exec(
		sqlUpdateTodoItem,
		item.Title,
		item.Id)
	return err
}

// UpdateTodoItem updates a single DBTodoItem within the DB
func (todoDb *TodoDB) DeleteTodoItem(item *DBTodoItem) error {
	_, err := todoDb.DB.Exec(
		sqlDeleteTodoItem,
		item.Id)
	return err
}
