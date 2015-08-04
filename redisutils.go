package redisutils

import (
	"log"
	"errors"
	"encoding/json"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

const (
	//DB CONFIG
	REDIS_SERVER = 	"127.0.0.1"
	REDIS_PORT = 	"6379"
	REDIS_URL = 	REDIS_SERVER + ":" + REDIS_PORT
	MAX_CONNS = 	10
)

var (
	//GLOBAL CONNECTION POOL
	REDIS_POOL 	*pool.Pool

	//ERROR MESSAGES
	ErrKeyNotSet = 		errors.New("keyNotSet")
)

//*********************************************************************************************************************************
//CONNECT & DISCONNECT

//CONNECT TO REDIS
func Connect () {
	p, err := pool.New("tcp", REDIS_URL, MAX_CONNS)
	if err != nil {
		log.Println("redisConnectError")
		log.Panicln(err)
		return
	}

	//store connection in global variable
	log.Println("RedisDB - Connected")
	REDIS_POOL = p
	return
}

//CLOSE ALL POOL CONNECTIONS
func Close () {
	REDIS_POOL.Empty()
	return
}

//*********************************************************************************************************************************
//GETTERS

//GET DATA FROM REDIS BY KEY
//check if the key exists
//passing in pool every time this func is called so this package has no dependencies on global vars
func Get (key string) (string, error) {
	//get a connection from the pool
	c, err := REDIS_POOL.Get()
	if err != nil {
		return "", err
	}
	defer REDIS_POOL.Put(c)

	//check if this key is stored in redis
	resp := 	c.Cmd("GET", key)
	
	//key is not in redis
	respType := resp.IsType(redis.Nil)
	if respType == true {
		return "", ErrKeyNotSet
	}

	//read value for key from redis
	value, err := resp.Str()

	//error while reading value
	if err != nil {
		return "", err
	}

	//return key value
	log.Println("REDIS - cache hit")
	return value, nil
}

//*********************************************************************************************************************************
//SETTERS

//SET DATA IN REDIS
//passing in pool every time this func is called so this package has no dependencies on global vars
//value is a struct of many possible types (User, Devices, Schedule,...)
func Set (key string, value interface{}) error {
	//get a connection from the pool
	c, err := REDIS_POOL.Get()
	if err != nil {
		return err
	}
	defer REDIS_POOL.Put(c)

	//convert value into string
	json, _ := 	json.Marshal(value)
	str := 		string(json)

	//save data to redis
	err = c.Cmd("SET", key, str).Err 
	
	//error while saving to redis
	if err != nil {
		return err
	}

	//data saved to redis
	return nil
}
