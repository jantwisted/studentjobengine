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

func SaveToRedis(con redis.Conn, job Job, index uint32) error{
	const prefix string = "JOB"
	jsonstr, err := json.Marshal(job)
	if err != nil {
		return err
	}
	_, err = con.Do("JSON.SET", prefix+strconv.FormatUint(uint64(index), 10), ".", jsonstr)
	if err != nil {
		return err
	}
	return nil
}

func GetFromRedis(con redis.Conn, key string) (string, error){
	s, err := redis.String(con.Do("JSON.GET", key))
	if err == redis.ErrNil{
		fmt.Println("Key doesn't exist!")
	} else if err != nil{
		return s, err
	}
	return s, nil
}

func SeqNextVal(con redis.Conn, key string)(uint32, error){
	seqval, err := redis.String(con.Do("GET", key))
	if err == redis.ErrNil{
		_, err = con.Do("SET", key, 0)
		if err != nil{
			return 0, err
		}
		seqval, err = redis.String(con.Do("GET", key))
	} else if err != nil{
		return 0,  err
	}
	seqvalu32, err := strconv.ParseUint(seqval, 10, 32)
	if err != nil{
		fmt.Println(err)
	}
	_, err = con.Do("SET", key, uint32(seqvalu32)+1)
	if err != nil{
		return 0, err
	}
	return uint32(seqvalu32), nil
}
