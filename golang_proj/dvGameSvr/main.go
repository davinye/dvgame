// dvGameSvr project main.go
/*package main

import (
	"conn"
	"fmt"
)

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	Room  int    "room"
}

func main() {
	player := &Player{"a", 1, 1, 1}
	player2 := &Player{"b", 2, 2, 2}

	fmt.Println(player2.Name, player2.Level)
	fmt.Println(player.Name, player.Level)

	fmt.Println(player.Level + player2.Level)

	var d *conn.Request = conn.GetRequest()
	d.Method = "aaaaaa"
	d.Params = "bbbbbbb"
	fmt.Println(d.Method)
}
*/

package main

import (
	_ "conn"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"socketest"
	_ "zxhd/domain"
	_ "zxhd/fbprototest"
)

var (
	Port        = flag.String("i", "127.0.0.1:9989", "IP port to listen on")
	logFileName = flag.String("log", "cServer.log", "Log file name")
	//configFileName = flag.String("configfile", "config.ini", "General configuration file")
	//configFile     = flag.String("configfile", "config.ini", "General configuration file")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	//	conn.RedigoTest()

	//	var d *domain.User = domain.CreatUser(domain.SYSTEM_USER)
	//	print(d.UserID)

	//	conn.DvMysql()
	//fbprototest.TestFBPrototest()

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//set logfile Stdout End
	//start listen
	listenErr := socketest.StartListen(*Port)
	if listenErr != nil {
		log.Fatalf("Server abort! Cause:%v \n", listenErr)
	}
}

/*
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
/*
/*

package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)
    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    os.Exit(0)
}
func checkError(err os.Error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.String())
        os.Exit(1)
    }
}

*/

/*

package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    service := ":7777"
    tcpAddr, err := net.ResolveTCPAddr("ip4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        daytime := time.LocalTime().String()
        conn.Write([]byte(daytime)) // don't care about return value
        conn.Close()                // we're finished with this client
    }
}
func checkError(err os.Error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.String())
        os.Exit(1)
    }
}

*/

/*

package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("ip4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handlerClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    daytime := time.LocalTime().String()
    conn.Write([]byte(daytime)) // don't care about return value
    // we're finished with this client
}
func checkError(err os.Error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.String())
        os.Exit(1)
    }
}

*/
