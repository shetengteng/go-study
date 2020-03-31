package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

// 将对象转换为json后在从json转换为指定的对象
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)

	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err

}
