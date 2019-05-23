package mongo

import (
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// 参考链接
// https://blog.csdn.net/yangzhengyi68/article/details/21518719
// 官网
// https://godoc.org/labix.org/v2/mgo

func TestMongo(t *testing.T) {
	mongo := GetMongo("userinfo", "user")
	c := mongo.Collection
	defer mongo.Close()

	var result []interface{}
	c.Insert(bson.M{"mongo": "sssss", "port": 3422})
	c.Find(nil).All(&result)
	t.Logf("table test %+v", result)

	//随便插入数据
	c.Insert(bson.M{"mongo_abc": "121221", "port_cdd": 3232})
	c.Find(nil).All(&result)
	t.Logf("table test %+v", result)

	var oneResult interface{}
	c.Find(nil).One(&oneResult)
	t.Logf("table test FineOne %+v", oneResult)

	//删除数据
	c.DropCollection()

	var otherResult interface{}
	c = mongo.UseCollect("test")
	c.Find(nil).One(&otherResult)
	t.Logf("table test Change Collection %+v", otherResult)
}

func init() {
	SetUp("localhost:27017/?maxPoolSize=8")
}
