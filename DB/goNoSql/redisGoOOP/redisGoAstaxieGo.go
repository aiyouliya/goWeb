package redisGoOOP

import (
	"fmt"

	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	// setting port is redis port
	client.Addr = "127.0.0.1:6379"
	//charactor operation
	client.Set("name", []byte("yezi")) //set key:value
	val, _ := client.Get("name")       //get key:value
	fmt.Println(string(val))           //print value
	client.Del("name")                 // del key:value

	//list operation
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("1", []byte(v))
	}
	dbvals, _ := client.Lrange("1", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("1")

}
