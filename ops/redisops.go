package ops

import(
	"github.com/gomodule/redigo/redis"
	"fmt"
	"strconv"
	"encoding/json"
)


func getRedisPool() *redis.Pool{
	return &redis.Pool{
		// Max number of idle connections in the pool.
		MaxIdle: 80,
		// Max number of connections.
		MaxActive: 12000,
		// Config a connection.
		Dial: func() (redis.Conn, error){
			c, err := redis.Dial("tcp", ":6379")
			if err != nil{
				panic(err.Error())
			}
			return c, err
		},

	}
}

func SaveToRedis(con redis.Conn, arg interface{}, prefix string, index uint32) error{
	jsonstr, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	_, err = con.Do("JSON.SET", prefix+strconv.FormatUint(uint64(index), 10), ".", jsonstr)
	if err != nil {
		return err
	}
	return nil
}

func GetKeys(con redis.Conn, pattern string) ([]string, error){
	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(con.Do("SCAN", iter, "MATCH", pattern))
		if err != nil{
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0{
			break
		}
	}
	return keys, nil
}

func GetJsonValue(con redis.Conn, key string)(string, error){
	json_content, err := redis.String(con.Do("JSON.GET", key))
	if err != nil{
		return json_content, err
	}
	return json_content, err
}

func SeqNextVal(con redis.Conn, key string)(uint32, error){
	var seqval string
	is_seq_exist, err := redis.String(con.Do("EXIST", key))
	if is_seq_exist == "0" {
		_, err = con.Do("SET", key, 0)
		if err != nil{
			return 6, err
		}
	}
	_, err = redis.String(con.Do("INCR", key))
	seqval, err = redis.String(con.Do("GET", key))
	seqvalu32, err := strconv.ParseUint(seqval, 10, 32)
	return uint32(seqvalu32), nil
}
