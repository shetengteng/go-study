package persist

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// rpc 接口
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	_, err := persist.Save(s.Client, item, s.Index)
	log.Printf("item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Println("error saving item %v :%v", item, err)
	}
	return err
}
