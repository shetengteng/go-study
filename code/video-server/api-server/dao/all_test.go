package dao

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 先初始化，清空数据库，然后测试，清空数据库
// init(dblogin,truncate tables) -> run tests -> clear data(truncate tables)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	// clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("delete", testDeleteUser)
	t.Run("reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("stt", "123")
	if notNil(err) {
		t.Errorf("error of addUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("stt")
	if pwd != "123" || notNil(err) {
		t.Errorf("error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("stt", "123")
	if notNil(err) {
		t.Errorf("error of deleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("stt")
	if notNil(err) {
		t.Errorf("error of regetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("preUser", testAddUser)
	t.Run("addVideo", testAddVideo)
	t.Run("getVideo", testGetVideo)
	t.Run("deleteVideo", testDeleteVideo)
	t.Run("regetVideo", testRegetVideoInfo)
}

var tempVid string

// 添加video，然后获取id
func testAddVideo(t *testing.T) {
	video, err := AddVideo(1, "myvideo")
	if notNil(err) {
		t.Errorf("error of addVideo ： %v", err)
	}
	tempVid = video.Id
}

func testGetVideo(t *testing.T) {
	_, err := GetVideo(tempVid)
	if notNil(err) {
		t.Errorf("error of getVideo: %v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempVid)
	if notNil(err) {
		t.Errorf("error of deleteVideo %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	video, err := GetVideo(tempVid)
	if err != nil {
		t.Errorf("error of reget video :%v", err)
	}
	if video != nil {
		t.Errorf("delete step is not corrent execution")
	}
}

func TestCommentWorkFlow(t *testing.T) {
	clearTables()
	t.Run("addUser", testAddUser) // 添加id为1的user
	t.Run("addComment", testAddComments)
	t.Run("getComments", testGetComments)
}

func testAddComments(t *testing.T) {
	err := AddComment("123", 1, "i like this")
	if notNil(err) {
		t.Errorf("error of addComments :%v", err)
	}
}

func testGetComments(t *testing.T) {
	vid := "123"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := GetComments(vid, from, to)
	if notNil(err) {
		t.Errorf("error of getComments %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment:%d %v \n", i, ele)
	}
}

func TestSessionWorkFlow(t *testing.T) {
	clearTables()
	t.Run("addSession", testAddSession)

}

func testAddSession(t *testing.T) {

}

func notNil(err error) bool {
	return err != nil
}
