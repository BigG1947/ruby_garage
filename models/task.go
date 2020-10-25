package models

import (
	"RubyGarage_v.2.0/utils"
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Id        int64  `json:"id"`
	Text      string `json:"text"`
	Priority  int    `json:"priority"`
	Deadline  string `json:"deadline"`
	Checked   bool   `json:"checked"`
	IdProject int64  `json:"id_project"`
}

type TaskList []Task

func (t *Task) Valid() (map[string]interface{}, bool) {
	if len(t.Text) == 0 {
		return utils.Message(false, "Task text is required"), false
	}

	if len(t.Deadline) == 0 {
		return utils.Message(false, "Task deadline is required"), false
	}

	if t.IdProject == 0 {
		return utils.Message(false, "Task project id is required"), false
	}

	return utils.Message(true, "Task is valid"), true
}

func (t *Task) Get(uid int64) (map[string]interface{}, bool) {
	if t.Id == 0 || t.IdProject == 0 {
		return utils.Message(false, "id and id_project is required parameters"), false
	}

	row := db.QueryRow("SELECT text, priority, deadline, checked FROM task t INNER JOIN project p ON p.id = t.id_project WHERE p.id_user = ? AND t.id = ?", uid, t.Id)

	if err := row.Scan(&t.Text, &t.Priority, &t.Deadline, &t.Checked); err != nil {
		if err == sql.ErrNoRows {
			return utils.Message(false, fmt.Sprintf("Task #%d not found", t.Id)), false
		}
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error"), false
	}

	response := utils.Message(true, "Ok!")
	response["task"] = t
	return response, true
}

func (t *Task) Create(uid int64) map[string]interface{} {
	if response, ok := t.Valid(); !ok {
		return response
	}

	if response, ok := isUserProject(uid, t.IdProject); !ok {
		return response
	}

	result, err := db.Exec("INSERT INTO task(text, priority, deadline, checked, id_project) VALUES (?, (SELECT IFNULL(MAX(priority), 0) + 1 FROM task WHERE id_project = ?), ?, ?, ?)", t.Text, t.IdProject, t.Deadline, t.Checked, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error")
	}

	if t.Id, err = result.LastInsertId(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please retry later")
	}

	if err := db.QueryRow("SELECT priority FROM task WHERE id = ?", t.Id).Scan(&t.Priority); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please retry later")
	}

	response := utils.Message(true, "Task has been created")
	response["task"] = t
	return response
}

func (t *Task) Edit(uid int64) map[string]interface{} {
	if t.Id == 0 {
		return utils.Message(false, "Id is required parameters")
	}

	if response, ok := t.Valid(); !ok {
		return response
	}

	if response, ok := isUserProject(uid, t.IdProject); !ok {
		return response
	}

	if _, err := db.Exec("UPDATE task SET text = ?, deadline = ? WHERE id = ? AND id_project = ?", t.Text, t.Deadline, t.Id, t.IdProject); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	response := utils.Message(true, "Update done!")
	response["task"] = t
	return response
}

func (t *Task) Delete(uid int64) map[string]interface{} {
	if t.Id == 0 || t.IdProject == 0 {
		return utils.Message(false, "id and id_project is required parameters")
	}

	if response, ok := t.Get(uid); !ok {
		return response
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	if _, err := tx.Exec("DELETE FROM task WHERE id = ? AND id_project = ?", t.Id, t.IdProject); err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	if _, err := tx.Exec("UPDATE task SET priority = priority - 1 WHERE id_project = ? AND priority > ? ", t.IdProject, t.Priority); err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	return utils.Message(true, "Task has been delete")
}

func (t *Task) Check(uid int64) map[string]interface{} {
	if t.Id == 0 || t.IdProject == 0 {
		return utils.Message(false, "Id and id_project is required parameters")
	}

	if response, ok := isUserProject(uid, t.IdProject); !ok {
		return response
	}

	if _, err := db.Exec("UPDATE task SET checked = ? WHERE id = ? AND id_project = ?", t.Checked, t.Id, t.IdProject); err != nil {
		return utils.Message(false, "DB connection error. Please try later")
	}

	return utils.Message(true, "Check status for this project has been update!")
}

func (t *Task) PriorityUp(uid int64) map[string]interface{} {
	if t.Id == 0 || t.IdProject == 0 {
		return utils.Message(false, "id and id_project is required parameters")
	}

	if response, ok := t.Get(uid); !ok {
		return response
	}

	if t.Priority == 1 {
		return utils.Message(false, "Task has maximal priority")
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = - (priority - 1) WHERE id = ? AND id_project = ?", t.Id, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = (priority + 1) WHERE priority = ? AND id_project = ?", t.Priority-1, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = - (priority) WHERE id = ? AND id_project = ?", t.Id, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	return utils.Message(true, "Priority up for task has done!")
}

func (t *Task) PriorityDown(uid int64) map[string]interface{} {
	if t.Id == 0 || t.IdProject == 0 {
		return utils.Message(false, "id and id_project is required parameters")
	}

	if response, ok := t.Get(uid); !ok {
		return response
	}

	var maxPriority int
	db.QueryRow("SELECT MAX(priority) FROM task WHERE id_project = ?", t.IdProject).Scan(&maxPriority)
	if t.Priority == maxPriority {
		return utils.Message(false, "Task has minimal priority")
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = - (priority + 1) WHERE id = ? AND id_project = ?", t.Id, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = (priority - 1) WHERE priority = ? AND id_project = ?", t.Priority+1, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	_, err = tx.Exec("UPDATE task SET priority = - (priority) WHERE id = ? AND id_project = ?", t.Id, t.IdProject)
	if err != nil {
		log.Printf("%s\n", err)
		_ = tx.Rollback()
		return utils.Message(false, "DB connection error. Please try later")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	return utils.Message(true, "Priority down for task has done!")
}

func (tl *TaskList) GetProjectTasks(pid int64) (map[string]interface{}, bool) {
	rows, err := db.Query("SELECT id, text, priority, deadline, checked, id_project FROM task WHERE id_project = ? ORDER BY priority", pid)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later"), false
	}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.Id, &t.Text, &t.Priority, &t.Deadline, &t.Checked, &t.IdProject); err != nil {
			log.Printf("%s\n", err)
			return utils.Message(false, "DB connection error. Please try later"), false
		}

		*tl = append(*tl, t)
	}

	return utils.Message(true, "Ok!"), true
}
