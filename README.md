# redisutils.go
Helper Functions for Using Redis with Golang

This package provides some helper functions used to interact with Redis and wraps around the radix.v2 driver.

---

###Usage

This library will connect to your Redis DB(s) and store the session connection data in a global variable "REDIS_POOL".  Anytime you want to use Redis, you will need to import this file and grab a connection from the pool.

There is some required setup for your environment. Please see below.

###Setup

This setup is done inside this library so that all Redis configuration is in one place.  This makes quick development easier.  Production deployments would not use this library.

- REDIS_SERVER:
	- The IP or DNS name of the server running Redis.

- REDIS_PORT:
	- The port on which Redis is running.  The default is port 6379

- MAX_CONNS:
	- The number of connections to Redis that will be held open for use.  Default is 10.

---

##Functions

###Connect()
Connect to Redis and store the pool data in the global variable.

###Get(key string)
Find data for an associated key in Redis.  Returns the data found as a string and an error if the key does not exist.

###Set(key string, value interface{})
Store data in Redis by the key.  The key's value is an interface which because it could be a struct of any type that is flattened into a JSON string before storage in Redis.  When getting the data back, you would need to Unmarshal the JSON back into a struct as needed.
