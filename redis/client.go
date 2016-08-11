package redis

import (
	"errors"
	"github.com/ftchao/meego"
	r "github.com/garyburd/redigo/redis"
	"time"
)

type RClient struct {
	pool *r.Pool
	// 	servers        []string
	// 	connectTimeout int
	// 	readTimeout    int
	// 	writeTimeout   int
	// 	fruits         []string
}

const (
	// 连接尝试次数 = RANDOM_TRY_MULTIPLE x 服务器数量
	RANDOM_TRY_MULTIPLE = 2
)

func Connect(node string) *RClient {

	servers := meego.ConfGet("redis." + node + ".server").([]interface{})
	connectTimeout := meego.ConfGet("redis." + node + ".connect.timeout").(int64)
	readTimeout := meego.ConfGet("redis." + node + ".write.timeout").(int64)
	writeTimeout := meego.ConfGet("redis." + node + ".read.timeout").(int64)

	return &RClient{
		pool: PoolConnect(servers, connectTimeout, readTimeout, writeTimeout),
	}
}

func (rc *RClient) Get(key string) (string, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	reply, err := r.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func PoolConnect(server []interface{}, connectTimeout, readTimeout, writeTimeout int64) *r.Pool {
	// todo
	//  1. 打乱连接服务器顺序, 平均分配负载
	//  2. 使用权重计算负载, 更能均衡的分担负债
	return &r.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (r.Conn, error) {
			for _, s := range server {
				// fmt.Printf("connect redis %s ... \n", s)
				c, err := r.DialTimeout("tcp", s.(string),
					time.Duration(connectTimeout)*time.Millisecond,
					time.Duration(readTimeout)*time.Millisecond,
					time.Duration(writeTimeout)*time.Millisecond)
				if err == nil && c != nil {
					return c, err
				}
			}
			return nil, errors.New("redispool: cannot connect to any redis server")
		},
	}
}
