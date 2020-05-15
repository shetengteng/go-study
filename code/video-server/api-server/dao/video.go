package dao

import (
	"database/sql"
	"time"
	"video-server/api-server/domain/entity"
	"video-server/api-server/utils"
)

func AddVideo(aid int, name string) (*entity.Video, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	// 必须指定该时间
	ctime := t.Format("2016-01-02 15:04:06")
	stmtIns, err := dbConn.Prepare(
		`INSERT INTO video_info(id ,author_id,name,display_ctime)
		 VALUES(?,?,?,?);`,
	)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &entity.Video{Id: vid, AuthorId: aid, Name: name, DisplayCreateTime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideo(vid string) (*entity.Video, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		return nil, err
	}
	var (
		aid               int
		displayCreateTime string
		name              string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &displayCreateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &entity.Video{
		Id:                vid,
		AuthorId:          aid,
		DisplayCreateTime: displayCreateTime,
		Name:              name,
	}
	return res, nil
}

func DeleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
