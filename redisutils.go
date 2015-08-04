/*
Package redisutils is used to simplify some usage of the radix.v2 Redis driver.

Use this package to connect, disconnect, host the redis pool data (global variable), and perform some Get and Set functions.
Basically, this library helps clean up your code base elsewhere.

This package uses and requires the radix.v2(pool/redis) drivers for Redis.  No other drivers are supported.

When connecting to a Redis server, this package will save a connection pool to a global variable.
Include this file wherever you need to use Redis.

It is highly suggested that you create another file for storing your Redis server and maxNumConnections as constants.

Note: this package is not meant to meant for production environments.
*/

package redisutils

import (
	"log"
	"errors"
	"encoding/json"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	//GLOBAL CONNECTION POOL
	POOL *pool.Pool

	//ERROR MESSAGES
	ErrKeyNotSet = 		errors.New("keyNotSet")
)

//*********************************************************************************************************************************
//CONNECT & DISCONNECT

//CONNECT TO REDIS SERVER
//only connects to one redis server
//creates a pool of connections
//server is a  "localhost:port" string
//maxNumConnections is the size of the pool you want to open
func Connect(server string, maxNumConnections int) {
	connPool, err := pool.new("tcp", )
	if err != nil {
		log.Println("redisutils.go-Connect error")
		log.Panicln(err)
		return
	}

	//store pool in global variable
	//access the pool by importing this file and "getting" a connection from the pool
	log.Println("redisutils.go-Connect okay")
	POOL = connPool
	return
}

//CLOSE ALL POOL CONNECTIONS
//using the radix Empty() function
func Close () {
	PPOL.Empty()
	return
}

//*********************************************************************************************************************************
//GETTERS

//GET DATA FROM REDIS BY KEY
//check if the key exists
func Get (key string) (string, error) {
	//get a connection from the pool
	c, err := PPOL.Get()
	if err != nil {
		return "", err
	}
	defer PPOL.Put(c)

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
//value is an interface{} because it can by any type struct
//this value is then flattened into JSON.
//when getting this key, you will need to Unmarshal the json back into a struct
func Set (key string, value interface{}) error {
	//get a connection from the pool
	c, err := PPOL.Get()
	if err != nil {
		return err
	}
	defer PPOL.Put(c)

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
