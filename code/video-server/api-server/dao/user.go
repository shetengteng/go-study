package dao

import (
	"database/sql"
	"log"
	"video-server/api-server/domain/entity"
)

func AddUser(user *entity.User) error {
	return AddUserCredential(user.Username, user.Pwd)
}

func AddUserCredential(loginName string, pwd string) error {
	// 不能每次都openConnection，否则会有很多半连接出现
	// 对sql进行预编译
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
