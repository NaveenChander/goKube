package core

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func setCache(key string, val string, expirationInSeconds int) error {

	if pool == nil {
		initPool()
	}

	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	_, err = conn.Do("EXPIRE", key, expirationInSeconds)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	return nil
}

func setCacheInBytes(key string, val []byte, expirationInSeconds int) error {

	if pool == nil {
		initPool()
	}

	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	_, err = conn.Do("EXPIRE", key, expirationInSeconds)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	return nil
}

func getCache(key string) (string, error) {
	if pool == nil {
		initPool()
	}

	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", err
	}

	return s, nil
}

func getCacheInBytes(key string) ([]byte, error) {
	if pool == nil {
		initPool()
	}

	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	var s []byte
	s, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return nil, err
	}

	return s, nil
}
