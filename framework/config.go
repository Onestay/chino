package framework

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"errors"
	"github.com/onestay/chino/framework/models"
)

type MongoSession struct {
	s *mgo.Session
	c *mgo.Collection
}

func InitDB(db, collection string) (*MongoSession, error) {
	s, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)

	c := s.DB(db).C(collection)

	return &MongoSession{s, c}, nil
}

func (ms MongoSession) CreateGuildConfig(guildID string) error {
	c, err := ms.c.Find(bson.M{"guild_id": guildID}).Count()
	if err != nil {
		return err
	}

	if c > 0 {
		return errors.New("Already a guild with that ID. That shouldn't happen.")
	}

	err = ms.c.Insert(&models.GuildConfig{
		ID: bson.NewObjectId(),
		GuildID: guildID,
		DisabledCommands: []string{},
		Prefix: "",
	})
	if err != nil {
		return err
	}

	return nil
}

func (ms MongoSession) GetPrefix(guildID string) (string, error) {
	gc, err := ms.queryWithGuildID(guildID)
	if err != nil {
		return "", err
	}

	log.Printf("Successfully created a new config for %v", guildID)
	return gc.Prefix, nil
}

func (ms MongoSession) GetDisabledCommands(guildID string) ([]string, error) {
	gc, err := ms.queryWithGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return gc.DisabledCommands, nil
}

func (ms MongoSession) queryWithGuildID(guildID string) (*models.GuildConfig, error) {
	res := models.GuildConfig{}

	err := ms.c.Find(bson.M{"guild_id": guildID}).One(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
