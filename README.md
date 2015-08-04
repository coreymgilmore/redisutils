# redisutils.go
Helper Functions for Using Redis with Golang

This package provides some helper functions used to interact with Redis and wraps around the radix.v2 driver.  It is not meant for power users for for production.
Basically, this package can be used for simplifying the connection, disconnecting, getting, and setting functions of Redis using the radix.v2 driver.

---

###Functions

####Connect(server string, maxNumConnections int)
- Connect to Redis and store the pool data in the global variable.
- `server` is a "localhost:port" string of the Redis server to connect to.
- `maxNumConnections` is the size of the connection pool you want to hold open.

####Get(key string)
Find data for an associated key in Redis.  Returns the data found as a string and an error if the key does not exist.

####Set(key string, value interface{})
Store data in Redis by the key.  The key's value is an interface which because it could be a struct of any type that is flattened into a JSON string before storage in Redis.  When getting the data back, you would need to Unmarshal the JSON back into a struct as needed.
