package dao

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video-server/api-server/domain/entity"
)

// 添加session
func AddSession(sid string, ttl int64, username string) error {

	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIn, err := dbConn.Prepare("INSERT INTO sessions(session_id,TTL,login_name) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmtIn.Close()
	_, err = stmtIn.Exec(sid, ttlStr, username)
	if err != nil {
		return err
	}
	return nil
}

// 获取session
func GetSession(sid string) (*entity.Session, error) {
	re := &entity.Session{}

	stmtOut, err := dbConn.Prepare("SELECT TTL,login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmtOut.Close()
	var (
		ttl      string
		username string
	)
	err = stmtOut.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ttlint, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		re.TTL = ttlint
		re.Username = username
		return re, nil
	}
	return nil, err
}

func GetSessionsMap() (*sync.Map, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	m := &sync.Map{}
	for rows.Next() {
		var (
			id       string
			ttlstr   string
			username string
		)
		if err = rows.Scan(&id, &ttlstr, &username); err != nil {
			log.Printf("get sessions map err: %s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			se := &entity.Session{
				Username: username,
				TTL:      ttl,
			}
			m.Store(id, se)
			log.Printf("session id :%s ttl: %d \n", id, ttl)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	defer stmtOut.Close()
	if _, err = stmtOut.Exec(sid); err != nil {
		return err
	}
	return nil
}
