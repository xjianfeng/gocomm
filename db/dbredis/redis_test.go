package dbredis

import (
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

// redis 操作
// https://www.cnblogs.com/woshimrf/p/5198361.html

// 官网
//https://godoc.org/github.com/gomodule/redigo/redis
// 示例
//https://blog.csdn.net/xcl168/article/details/44622013

/*
func TestRedis(t *testing.T) {
	red := GetRedis()
	defer red.Close()

	red.Do("SELECT", 1)

	red.Do("SET", "sText", "go_reids_test")
	red.Do("SETNX", "text_go_reids", "test")
	red.Do("INCR", "iGlobalUid")
	cnt, _ := redis.Int(red.Do("HGET", "hInfo", "UserCnt"))
	gr, _ := redis.String(red.Do("GET", "sText"))
	uid, _ := redis.Int(red.Do("GET", "iGlobalUid"))
	roomId, _ := redis.Int(red.Do("LRANGE", "lRoomListFriend", 0, 0))

	serveAdress, _ := redis.StringMap(red.Do("HGETALL", "hServerList"))

	fmt.Println(gr, uid, serveAdress, "cnt", cnt, "roomId", roomId)
}

func TestPipeLine(t *testing.T) {
	c := GetRedis()
	defer c.Close()

	info := make(map[int]string)
	info[100001] = "120.2.1.2:9000"
	info[100002] = "121.2.1.2:9001"
	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Send("SET", "foo1", "bar1")
	c.Send("GET", "foo1")
	c.Send("HSET", "hInfo", "value", info)
	c.Flush()
	v, _ := c.Receive() // reply from SET
	fmt.Println(v)
	v, _ = redis.String(c.Receive()) // reply from GET
	fmt.Println(v)
	v, _ = c.Receive() // reply from SET
	fmt.Println(v)
	v, _ = redis.String(c.Receive()) // reply from GET
	fmt.Println(v)
}

func TestSetList(t *testing.T) {
	c := GetRedis()
	defer c.Close()

	s, e := redis.Strings(c.Do("SMEMBERS", "RoomType1---"))
	if e != nil {
		fmt.Println("error %s", e.Error())
	}
	fmt.Printf("%v", s)
}
*/

func TestRedisCURD(t *testing.T) {
	SetUp(&RedisConf{
		Host:        "127.0.0.1:6379",
		Password:    "",
		MaxIdle:     10,
		MaxActive:   500,
		DefaultDb:   0,
		IdleTimeout: time.Duration(200),
	})
	r := GetRedis()
	// Create Data
	ret, err := r.Do("INCR", "T_iIncrData")
	t.Logf("CREATE INCR ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("SET", "T_sData", "hello World!")
	t.Logf("CREATE SET ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("HSET", "T_Hdata", "cKey", "Key")
	t.Logf("CREATE HSET ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("HSET", "T_Hdata", "cValue", "Value")
	t.Logf("CREATE HSET ret:%+v, err:%+v", ret, err)

	// READ DATA
	ret, err = redis.Int(r.Do("GET", "T_iIncrData"))
	t.Logf("READ INCR ret:%+v, err:%+v", ret, err)

	ret, err = redis.String(r.Do("GET", "T_sData"))
	t.Logf("READ SET ret:%+v, err:%+v", ret, err)

	ret, err = redis.StringMap(r.Do("HGETALL", "T_Hdata"))
	t.Logf("READ HGETALL ret:%+v, err:%+v", ret, err)

	// UPDATE Data
	ret, err = r.Do("INCR", "T_iIncrData")
	t.Logf("UPDATE INCR ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("SET", "T_sData", "Change Data")
	t.Logf("UPDATE SET ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("HSET", "T_Hdata", "cKey", "1")
	t.Logf("UPDATE HSET ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("HSET", "T_Hdata", "cValue", "2")
	t.Logf("UPDATE HSET ret:%+v, err:%+v", ret, err)

	// READ DATA
	ret, err = redis.Int(r.Do("GET", "T_iIncrData"))
	t.Logf("READ INCR ret:%+v, err:%+v", ret, err)

	ret, err = redis.String(r.Do("GET", "T_sData"))
	t.Logf("READ SET ret:%+v, err:%+v", ret, err)

	ret, err = redis.StringMap(r.Do("HGETALL", "T_Hdata"))
	t.Logf("READ HGETALL ret:%+v, err:%+v", ret, err)

	// DELETE DATA
	ret, err = r.Do("DEL", "T_iIncrData")
	t.Logf("DEL INCR ret:%+v, err:%+v", ret, err)

	ret, err = r.Do("DEL", "T_sData")
	t.Logf("DEL SET ret:%+v, err:%+v", ret, err)
	ret, err = r.Do("HDEL", "T_Hdata", "cKey")
	t.Logf("DEL HDEL ret:%+v, err:%+v", ret, err)

	// READ DATA
	ret, err = redis.Int(r.Do("GET", "T_iIncrData"))
	t.Logf("READ INCR ret:%+v, err:%+v", ret, err)

	ret, err = redis.String(r.Do("GET", "T_sData"))
	t.Logf("READ SET ret:%+v, err:%+v", ret, err)

	ret, err = redis.StringMap(r.Do("HGETALL", "T_Hdata"))
	t.Logf("READ HGETALL ret:%+v, err:%+v", ret, err)

	//CLEAN DATA
	r.Do("DEL", "T_Hdata")
}
