package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todoapp/dbconn"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Todo struct {
	Id           int           `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Category     string        `json:"category"`
	Progress     string        `json:"progress"`
	Status       string        `json:"status"`
	Deadline     sql.NullTime  `json:"deadline"`
	CreatedTime  sql.NullTime  `json:"createdTime"`
	UpdatedTime  sql.NullTime  `json:"updatedTime"`
	RemainingDay sql.NullInt64 `json:"remainingDay"`
}

type Category struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

func main() {

	e := echo.New()

	e.GET("/listall", listAll)
	e.GET("/listcategory", listCategory)
	e.GET("/getitem/:data", getItem)

	e.POST("/additem/:data", addItem)
	e.POST("/addcategory", addCategory)
	e.POST("/deletecategory/:data", deleteCategory)
	e.POST("/deleteitem/:data", deleteItem)

	e.PUT("/updateitem/:data", updateItem)
	e.PUT("/jobdone/:data", jobDone)

	e.Start(":8080")

}

func listAll(c echo.Context) error {

	var rowcount int
	conn := dbconn.Connection()
	defer conn.Close()

	select2, err := conn.Query("SELECT COUNT(*) From todo.todo")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select2.Next() {
		err1 := select2.Scan(&rowcount)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
	}

	items := make([]Todo, rowcount)

	select1, err := conn.Query("SELECT * From todo.todo")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	counter := 0
	for select1.Next() {
		err1 := select1.Scan(&items[counter].Id, &items[counter].Title, &items[counter].Description, &items[counter].Category, &items[counter].Progress, &items[counter].Deadline, &items[counter].Status, &items[counter].CreatedTime, &items[counter].UpdatedTime, &items[counter].RemainingDay)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
		counter++
	}

	for i := 0; i < len(items); i++ {

		_, _, day, _, _, _ := remainingDayCheck(time.Now(), items[i].Deadline.Time)
		//items[i].RemainingDay.Int64 = int64(day)
		if items[i].Progress == "overdue" || items[i].Progress == "done" {
			day = 0
		}

		update, err := conn.Query("UPDATE todo.todo SET remainingday=? WHERE id=?", int64(day), items[i].Id)
		if err != nil {
			log.Fatalf("Error: %+v\n", err)
		}
		defer update.Close()

		if items[i].Progress != "done" {

			var newProgress string

			progress := progressCheck(time.Now(), items[i].Deadline.Time)

			select1, err := conn.Query("SELECT progress From todo.progress WHERE id=?", progress)
			if err != nil {
				log.Fatalf("Error: %+v\n", err)
			}

			for select1.Next() {
				err1 := select1.Scan(&newProgress)
				if err1 != nil {
					log.Fatalf("Error: %+v\n", err1)
				}
			}

			update, err := conn.Query("UPDATE todo.todo SET progress=? WHERE id=?", newProgress, items[i].Id)
			if err != nil {
				log.Fatalf("Error: %+v\n", err)
			}
			defer update.Close()

		}
	}

	select2, err2 := conn.Query("SELECT * From todo.todo")
	if err2 != nil {
		log.Fatalf("Error: %+v\n", err2)
	}

	counter = 0
	for select2.Next() {
		err1 := select2.Scan(&items[counter].Id, &items[counter].Title, &items[counter].Description, &items[counter].Category, &items[counter].Progress, &items[counter].Deadline, &items[counter].Status, &items[counter].CreatedTime, &items[counter].UpdatedTime, &items[counter].RemainingDay)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
		counter++
	}

	return c.JSON(http.StatusOK, items)
}

func getItem(c echo.Context) error {

	item := Todo{}
	id := c.QueryParam("id")

	conn := dbconn.Connection()
	defer conn.Close()

	select1, err := conn.Query("SELECT * From todo.todo WHERE id=?", id)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select1.Next() {
		err1 := select1.Scan(&item.Id, &item.Title, &item.Description, &item.Category, &item.Progress, &item.Deadline, &item.Status, &item.CreatedTime, &item.UpdatedTime, &item.RemainingDay)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
	}

	return c.JSON(http.StatusOK, item)
}

func addItem(c echo.Context) error {

	item := Todo{}
	categoryid, _ := strconv.Atoi(c.QueryParam("categoryid"))

	respBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &item)
	if err != nil {
		return err
	}

	item.Category = categoryCheck(categoryid)
	item.CreatedTime.Time = time.Now()

	status := statusCheck(item.Description)

	progress := progressCheck(item.CreatedTime.Time, item.Deadline.Time)

	_, _, day, _, _, _ := remainingDayCheck(item.CreatedTime.Time, item.Deadline.Time)
	item.RemainingDay.Int64 = int64(day)
	if progress == 2|3 {
		item.RemainingDay.Int64 = 0
	}

	conn := dbconn.Connection()
	defer conn.Close()

	select1, err := conn.Query("SELECT progress From todo.progress WHERE id=?", progress)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select1.Next() {
		err1 := select1.Scan(&item.Progress)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
	}

	if status {
		conn := dbconn.Connection()
		defer conn.Close()

		select1, err := conn.Query("SELECT status From todo.status WHERE id=2")
		if err != nil {
			log.Fatalf("Error: %+v\n", err)
		}

		for select1.Next() {
			err1 := select1.Scan(&item.Status)
			if err1 != nil {
				log.Fatalf("Error: %+v\n", err)
			}
		}

	} else {
		conn := dbconn.Connection()
		defer conn.Close()

		select1, err := conn.Query("SELECT status From todo.status WHERE id=1")
		if err != nil {
			log.Fatalf("Error: %+v\n", err)
		}

		for select1.Next() {
			err1 := select1.Scan(&item.Status)
			if err1 != nil {
				log.Fatalf("Error: %+v\n", err1)
			}
		}
	}

	insert1, err := conn.Query("INSERT INTO todo.todo (title,description,category,progress,status,deadline,createdtime,remainingday) VALUES(?,?,?,?,?,?,?,?)", item.Title, item.Description, item.Category, item.Progress, item.Status, item.Deadline.Time, item.CreatedTime.Time, item.RemainingDay.Int64)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	defer insert1.Close()

	return c.String(http.StatusOK, "Item Added.")
}

func deleteItem(c echo.Context) error {
	id := c.QueryParam("id")

	conn := dbconn.Connection()
	defer conn.Close()

	_, err := conn.Query("DELETE FROM todo.todo WHERE id=?", id)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	return c.String(http.StatusOK, "Item Deleted.")
}

func addCategory(c echo.Context) error {

	category := Category{}

	conn := dbconn.Connection()
	defer conn.Close()

	respBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &category)
	if err != nil {
		return err
	}

	insert1, err := conn.Query("INSERT INTO todo.category (category) VALUES(?)", category.Category)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	defer insert1.Close()

	return c.String(http.StatusOK, "Category Added.")
}

func deleteCategory(c echo.Context) error {
	id := c.QueryParam("id")

	conn := dbconn.Connection()
	defer conn.Close()

	_, err := conn.Query("DELETE FROM todo.category WHERE id=?", id)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	return c.String(http.StatusOK, "Category Deleted.")
}

func listCategory(c echo.Context) error {

	var rowcount int
	conn := dbconn.Connection()
	defer conn.Close()

	select2, err := conn.Query("SELECT COUNT(*) From todo.category")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select2.Next() {
		err1 := select2.Scan(&rowcount)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err)
		}
	}

	categories := make([]Category, rowcount)

	select1, err := conn.Query("SELECT * From todo.category")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	counter := 0
	for select1.Next() {
		err1 := select1.Scan(&categories[counter].Id, &categories[counter].Category)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
		counter++
	}

	return c.JSON(http.StatusOK, categories)
}

func updateItem(c echo.Context) error {

	item := Todo{}
	categoryid, _ := strconv.Atoi(c.QueryParam("categoryid"))

	respBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &item)
	if err != nil {
		return err
	}

	item.Category = categoryCheck(categoryid)
	item.UpdatedTime.Time = time.Now()

	conn := dbconn.Connection()
	defer conn.Close()

	status := statusCheck(item.Description)

	if status {
		conn := dbconn.Connection()
		defer conn.Close()

		select1, err := conn.Query("SELECT status From todo.status WHERE id=2")
		if err != nil {
			log.Fatalf("Error: %+v\n", err)
		}

		for select1.Next() {
			err1 := select1.Scan(&item.Status)
			if err1 != nil {
				log.Fatalf("Error: %+v\n", err1)
			}
		}

	} else {
		conn := dbconn.Connection()
		defer conn.Close()

		select1, err := conn.Query("SELECT status From todo.status WHERE id=1")
		if err != nil {
			log.Fatalf("Error: %+v\n", err)
		}

		for select1.Next() {
			err1 := select1.Scan(&item.Status)
			if err1 != nil {
				log.Fatalf("Error: %+v\n", err1)
			}
		}
	}

	var newProgress string

	progress := progressCheck(item.UpdatedTime.Time, item.Deadline.Time)

	select1, err := conn.Query("SELECT progress From todo.progress WHERE id=?", progress)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select1.Next() {
		err1 := select1.Scan(&newProgress)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
	}

	update, err := conn.Query("UPDATE todo.todo SET progress=? WHERE id=?", newProgress, item.Id)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	defer update.Close()

	update1, err1 := conn.Query("UPDATE todo.todo SET title=?, description=?, category=?, deadline=?, updatedtime=? ,status=? WHERE id=?", item.Title, item.Description, item.Category, item.Deadline.Time, item.UpdatedTime.Time, item.Status, item.Id)
	if err1 != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	defer update1.Close()

	return c.String(http.StatusOK, "Update done.")
}

func statusCheck(description string) bool {

	statusbool := strings.Contains(description, "acil")

	return statusbool
}

func categoryCheck(catid int) string {
	var category string
	conn := dbconn.Connection()
	defer conn.Close()

	select1, err := conn.Query("SELECT category From todo.category WHERE id=?", catid)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select1.Next() {
		err1 := select1.Scan(&category)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err)
		}
	}
	return category
}

func progressCheck(createdTime time.Time, deadline time.Time) int {

	if createdTime.Before(deadline) {
		return 1
	}

	if createdTime.After(deadline) {
		return 3
	}

	return 0
}

func remainingDayCheck(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {

		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func jobDone(c echo.Context) error {

	var progress string
	id := c.QueryParam("id")

	conn := dbconn.Connection()
	defer conn.Close()

	select1, err := conn.Query("SELECT progress From todo.progress WHERE id=2")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	for select1.Next() {
		err1 := select1.Scan(&progress)
		if err1 != nil {
			log.Fatalf("Error: %+v\n", err1)
		}
	}

	update, err := conn.Query("UPDATE todo.todo SET progress=? WHERE id=?", progress, id)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	defer update.Close()

	return c.String(http.StatusOK, "Job Done.")
}
