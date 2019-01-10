package data

import (
	"log"

	"github.com/c0z0/go-shrt/app/model"
	"github.com/zemirco/uid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	m *mgo.Database
}

func NewDB(uri, dbName string) *DB {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to db")

	db := session.DB(dbName)

	return &DB{m: db}
}

func (db *DB) Create(urlString string) model.Url {
	url := model.Url{Url: urlString, Id: uid.New(5)}

	q := db.m.C("urls").Find(bson.M{"url": urlString})
	if c, _ := q.Count(); c != 0 {
		mUrl := map[string]string{}
		q.One(&mUrl)
		return model.Url{Url: mUrl["url"], Id: mUrl["_id"]}
	}

	err := db.m.C("urls").Insert(map[string]string{"_id": url.Id, "url": url.Url})

	if err != nil {
		log.Fatal(err)
	}

	return url
}

func (db *DB) Find(id string) *model.Url {
	q := db.m.C("urls").Find(bson.M{"_id": id})
	if c, _ := q.Count(); c == 0 {
		return nil
	}
	mUrl := map[string]string{}

	q.One(&mUrl)

	return &model.Url{Url: mUrl["url"], Id: mUrl["_id"]}
}
