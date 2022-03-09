package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"server/models"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// DB connection string
// for localhost sqlserver
const connectionString = "odbc:server=localhost\\MSSQLSERVER01;database=ToDoDB;user id=test;password=test123"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	fmt.Println("new task name ", task.Task)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	task.ID = id
	updateOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)

}

// get all task from the DB and return it
func getAllTask() []models.ToDoList {
	var results []models.ToDoList
	ctx := context.Background()
	e := db.PingContext(ctx)
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Printf("Connected!\n")

	cur, err := db.QueryContext(ctx, "getAllTask")

	if err != nil {
		fmt.Println("there was an error executing the sp", err)
		return results
	}

	defer cur.Close()

	for cur.Next() {
		fmt.Println("getting one row")
		var id int64
		var task string
		var status bool
		var createdDateTime, lastupdateddatetime time.Time
		var result1 models.ToDoList
		e := cur.Scan(&id, &task, &status, &createdDateTime, &lastupdateddatetime)

		if e != nil {
			fmt.Println("Error ", e)
			return results
		}

		result1.ID = id
		result1.Task = task
		result1.Status = status
		result1.CreatedDateTime = createdDateTime
		result1.LastUpdatedDateTime = lastupdateddatetime
		fmt.Println("status ", status)

		results = append(results, result1)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("closing result set")
	cur.Close()
	return results
}

// Insert one task in the DB
func insertOneTask(todoItem models.ToDoList) {
	ctx := context.Background()
	fmt.Println("Task name", todoItem.Task)
	var rowid int64
	_, err := db.ExecContext(ctx, "dbo.InsertOneTask", sql.Named("task", todoItem.Task), sql.Out{Dest: &rowid})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rowid is %d", rowid)

	fmt.Println("Inserted a Single Record ", rowid, err)
}

func updateOneTask(todoItem models.ToDoList) {
	ctx := context.Background()
	fmt.Println("Task name", todoItem.Task)
	insertResult, err := db.ExecContext(ctx, "updateTask",
		sql.Named("id", todoItem.ID),
		sql.Named("status", todoItem.Status),
		sql.Named("task", todoItem.Task),
	)

	if err != nil {
		log.Fatal(err)
	}

	numberOfRows, lasterror := insertResult.RowsAffected()

	fmt.Println("modified count: ", numberOfRows, lasterror)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := strconv.ParseInt(task, 10, 64)
	ctx := context.Background()
	numberofRows, err := db.ExecContext(ctx, "deleteOneTask",
		sql.Named("id", id),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", numberofRows)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	ctx := context.Background()
	numberofRows, err := db.ExecContext(ctx, "deleteAllTasks")
	if err != nil {
		log.Fatal(err)
	}

	deletedRows, error := numberofRows.RowsAffected()
	fmt.Println("Deleted Document", deletedRows, error)
	return deletedRows
}
