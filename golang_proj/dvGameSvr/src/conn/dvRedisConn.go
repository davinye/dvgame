package conn

/*
import (
	"fmt"
	//"io"
	"log"
	//"net/http"
	//"runtime"
	//"time"



	"github.com/garyburd/redigo/redis"
)

// 连接池大小
var MAX_POOL_SIZE = 20
var redisPoll chan redis.Conn

func putRedis(conn redis.Conn) {
	// 基于函数和接口间互不信任原则，这里再判断一次，养成这个好习惯哦
	if redisPoll == nil {
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
	}
	if len(redisPoll) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	redisPoll <- conn
}
func InitRedis(network, address string) redis.Conn {
	// 缓冲机制，相当于消息队列
	if len(redisPoll) == 0 {
		// 如果长度为0，就定义一个redis.Conn类型长度为MAX_POOL_SIZE的channel
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				c, err := redis.Dial(network, address)
				if err != nil {
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPoll
}
func redisServer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	c := InitRedis("tcp", "127.0.0.1:6379")
	dbkey := "netgame:info"
	if ok, err := redis.Bool(c.Do("LPUSH", dbkey, "yanetao")); ok {
	} else {
		log.Print(err)
	}
	msg := fmt.Sprintf("用时：%s", time.Now().Sub(startTime))
	io.WriteString(w, msg+"\n\n")
}
func main() {
//	fmt.Println("start")
//	// 利用cpu多核来处理http请求，这个没有用go默认就是单核处理http的，这个压测过了，请一定要相信我
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	http.HandleFunc("/", redisServer)
//	http.ListenAndServe(":9527", nil)

	conn, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn)
	}

	size, err := conn.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)

}

*/

import (
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
)

func RedigoTest() {
	// connect to localhost, make sure to have redis-server running on the default port
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// add some keys
	if _, err = conn.Do("SET", "k1", "a2"); err != nil {
		log.Fatal(err)
	}
	if _, err = conn.Do("SET", "k2", "b2"); err != nil {
		log.Fatal(err)
	}
	// for fun, let's leave k3 non-existing

	// get many keys in a single MGET, ask redigo for []string result
	strs, err := redis.Strings(conn.Do("MGET", "k1", "k2", "k3"))
	if err != nil {
		log.Fatal(err)
	}

	// prints [a b ]
	fmt.Println(strs)

	// now what if we want some integers instead?
	if _, err = conn.Do("SET", "k4", "1"); err != nil {
		log.Fatal(err)
	}
	if _, err = conn.Do("SET", "k5", "2"); err != nil {
		log.Fatal(err)
	}

	// get the keys, but ask redigo to give us a []interface{}
	// (it doesn't have a redis.Ints helper).
	vals, err := redis.Values(conn.Do("MGET", "k4", "k5", "k6"))
	if err != nil {
		log.Fatal(err)
	}

	// scan the []interface{} slice into a []int slice
	var ints []int
	if err = redis.ScanSlice(vals, &ints); err != nil {
		log.Fatal(err)
	}

	// prints [1 2 0]
	fmt.Println(ints)

}
