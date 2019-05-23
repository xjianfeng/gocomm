package mongo

import (
	"fmt"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// 参考链接
// https://blog.csdn.net/yangzhengyi68/article/details/21518719
// 官网
// https://godoc.org/labix.org/v2/mgo

func TestMongo(t *testing.T) {
	var result interface{}
	mongo := GetMongo("userinfo", "user")
	c := mongo.Collection
	defer mongo.Close()

	c.Find(nil).One(&result)
	fmt.Println("table user", result)

	c = mongo.UseCollect("test")
	c.Find(nil).One(&result)
	fmt.Println("talbe test", result)
}

type MongoInfo struct {
	Id    bson.ObjectId `bson:"_id"`
	Mongo string        `bson:"mongo"`
	Port  int           `bson:"port"`
}

//test官方使用
func TestGlobalMongo(t *testing.T) {
	var result interface{}
	session := MONGO_SESSION

	db := session.DB("userinfo")
	c := db.C("user")
	defer session.Close()

	c.Find(nil).One(&result)
	fmt.Println("table user", result)

	info := []MongoInfo{}
	c = db.C("test")
	c.Find(nil).All(&info)
	fmt.Println("talbe test", info)

	c.Remove(bson.D{{"_id", "5b1f45ba4b668c8a8892d403"}})
	c.Find(nil).All(&info)
	fmt.Println("talbe test", info)
	//按规则插入数据
	//c.Insert(bson.M{"mongo": "sssss", "port": 3422})
	//c.Find(nil).All(&info)
	//fmt.Println("talbe test", info)

	//随便插入数据
	//c.Insert(bson.M{"mongo_abc": "121221", "port_cdd": 3232})
	//c.Find(nil).All(&info)
	//fmt.Println("talbe test", info)
}
