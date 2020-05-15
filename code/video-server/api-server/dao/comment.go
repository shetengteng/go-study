package dao

import (
	"video-server/api-server/domain/entity"
	"video-server/api-server/utils"
)

func AddComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments(id,video_id,author_id,content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 使用from 表示起始时间,使用to 表示结束时间
func GetComments(vid string, from, to int) ([]*entity.CommentDTO, error) {
	stmtOut, err := dbConn.Prepare(`
		SELECT 
			comments.id,
			users.login_name,
			comments.content
		FROM comments
		INNER JOIN users
		ON users.id = comments.author_id
		WHERE comments.video_id = ? 
		AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	if err != nil {
		return nil, err
	}

	var res []*entity.CommentDTO

	rows, err := stmtOut.Query(vid, from, to)
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &entity.CommentDTO{
			Id:      id,
			VideoId: vid,
			Author:  name,
			Content: content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
