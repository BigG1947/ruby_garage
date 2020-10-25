package models

import (
	"RubyGarage_v.2.0/utils"
	"database/sql"
	"log"
)

type Project struct {
	Id     int64    `json:"id"`
	Name   string   `json:"name"`
	IdUser int64    `json:"id_user"`
	Tasks  TaskList `json:"tasks"`
}

type ProjectList []Project

func (p *Project) Valid() (map[string]interface{}, bool) {
	if len(p.Name) < 4 {
		return utils.Message(false, "Project name must be a least 4 characters"), false
	}

	if p.IdUser == 0 {
		return utils.Message(false, "User not authenticate"), false
	}

	var id int64

	err := db.QueryRow("SELECT id FROM project WHERE name = ? AND id_user = ?", p.Name, p.IdUser).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error"), false
	}

	if id != 0 {
		return utils.Message(false, "Project already exist"), false
	}

	return utils.Message(true, "Validate passed"), true
}

func (p *Project) Create() map[string]interface{} {
	if response, ok := p.Valid(); !ok {
		return response
	}

	result, err := db.Exec("INSERT INTO project(name, id_user) VALUES (?, ?)", p.Name, p.IdUser)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Try please later")
	}

	if p.Id, err = result.LastInsertId(); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Try please later")
	}

	response := utils.Message(true, "Project has been created")
	response["project"] = p
	return response
}

func (p *Project) Edit() map[string]interface{} {
	if p.Id == 0 {
		return utils.Message(false, "Id is required parameters")
	}

	if response, ok := p.Valid(); !ok {
		return response
	}

	if _, err := db.Exec("UPDATE project SET name = ? WHERE id = ? AND id_user = ?", p.Name, p.Id, p.IdUser); err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error")
	}

	response := utils.Message(true, "Update done!")
	response["project"] = p
	return response
}

func (p *Project) Delete() map[string]interface{} {
	if p.Id == 0 {
		return utils.Message(false, "Id is required parameters")
	}

	_, err := db.Exec("DELETE FROM project WHERE id = ? AND id_user = ?", p.Id, p.IdUser)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error")
	}

	return utils.Message(true, "Delete done!")
}

func (pl *ProjectList) Get(uid int64) map[string]interface{} {
	rows, err := db.Query("SELECT id, name, id_user FROM project WHERE id_user = ?", uid)
	if err != nil {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error. Please try later")
	}

	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.Id, &p.Name, &p.IdUser); err != nil {
			log.Printf("%s\n", err)
			return utils.Message(false, "DB connection error. Please try later")
		}

		if response, ok := p.Tasks.GetProjectTasks(p.Id); !ok {
			return response
		}

		*pl = append(*pl, p)
	}

	response := utils.Message(true, "success")
	response["projects"] = pl
	return response
}

func isUserProject(uid int64, pid int64) (map[string]interface{}, bool) {
	var id int64
	err := db.QueryRow("SELECT id FROM project WHERE id = ? AND id_user = ?", pid, uid).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return utils.Message(false, "DB connection error"), false
	}

	if err == sql.ErrNoRows {
		return utils.Message(false, "Project does not belong to this user"), false
	}

	return utils.Message(true, "Project is belong to this user"), true
}
