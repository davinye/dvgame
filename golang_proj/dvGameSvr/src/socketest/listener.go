package socketest

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
	"uuid"
	/*"goCS/consistent"
	"code.google.com/p/go-uuid/uuid"*/)

const (
	ConnectionMax = 100 //CS max connect
)

//conn Pool info
var (
	poolLock sync.RWMutex
	poolCli  [ConnectionMax]*CliInfo
)

//Web user info
type UserInfo struct {
	WS_ID       int
	WS_Name     string
	ServiceName string
}

//Cli info
type CliInfo struct {
	AssignID   int            //cs assign ID
	Conn       net.Conn       // The TCP/IP connectin to the player.
	ConnTime   time.Time      //连接时间
	VerifyKey  string         //连接验证KEY
	ConnVerify bool           //是否验证
	ServerType int32          //服务器类型（1DB服务器，2WEB服务器）
	NodeStat   int32          //服务器当前状态（0、宕机；1、正常；2、数据导入中；3、准备中;4、数据迁出中
	Address    string         //服务地址
	Port       int            //服务端口
	BackupSer  map[string]int //备份服务器列表map(ip:port)
	sync.RWMutex
}

type hashPool struct {
	Version  int
	Circle   map[uint32]string //hash圈节点分布
	Replicas map[string]int    //hash圈节点范围
}

// Client disconnect
func (cli *CliInfo) disconnect(clientID int) {
	poolLock.Lock()
	defer poolLock.Unlock()
	cli.Conn.Close()
	log.Printf("Client: %s quit\n", cli.VerifyKey)
	if cli.ServerType == 1 {
		println("移除clientID = ", clientID)

		poolCli[clientID] = nil

		//todo:掉线处理
		/*	if ok := cli.removeDBS(); ok {
				poolCli[clientID] = nil
			}
		*/
	} else {

	}

}

//listen handle
func (cli *CliInfo) listenHandle(clientID int) {

	headBuff := make([]byte, 50) // set read stream size
	defer cli.Conn.Close()
	cli.Conn.Read(headBuff)
	println(string(headBuff))

	//send verify Key：
	b := []byte(cli.VerifyKey)
	cli.Conn.Write(b)
	println("listenHandle  cli.VerifyKey---- ", cli.VerifyKey)
	// fmt.Println("cli-IP:", cli.Conn.RemoteAddr().String())

	//await 10 second verify
	cli.Conn.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))
	/*
		forControl := true
		for forControl {
			var headNum int
			for headNum < cap(headBuff) {
				readHeadNum, readHeadErr := cli.Conn.Read(headBuff[headNum:])
				if readHeadErr != nil {
					log.Println("errHead:", readHeadErr)
					forControl = false
					break
				}
				headNum += readHeadNum
			}
			if headNum == cap(headBuff) {

					//pack head Handle
					packHead := packHeadAnalysis(headBuff)
					bodyBuff := make([]byte, packHead.PackLen)
					var bodyNum int
					for bodyNum < cap(bodyBuff) {
						readBodyNum, readBodyErr := cli.Conn.Read(bodyBuff[bodyNum:])
						if readBodyErr != nil {
							log.Println("errBody:", readBodyErr)
							forControl = false
							break
						}
						bodyNum += readBodyNum
					}
					if bodyNum == int(packHead.PackLen) {
						//pack body Handle
						cli.packBodyAnalysis(clientID, packHead, bodyBuff)
						// fmt.Printf("packType:%d;packOther:%d;packLen:%d\n", packHead.PackType, packHead.PackOther, packHead.PackLen)
					}
			}
		}*/
	cli.disconnect(clientID)
}

//Check or assign new conn
func NewConnection_CS(conn net.Conn) (ok bool, index int, info *CliInfo) {
	poolLock.Lock()
	defer poolLock.Unlock()

	//Assign ID for client
	var i int
	for i = 0; i < ConnectionMax; i++ {
		if poolCli[i] == nil {
			break
		}
	}

	//Too many connections
	if i > ConnectionMax {
		log.Printf("Too many connections! Active Denial %s\n", conn.RemoteAddr().String())
		conn.Close()
		return false, 0, nil
	}

	println("poolCli lenght ", i)

	//Create client base info
	Cli := new(CliInfo)
	Cli.Conn = conn
	Cli.ConnTime = time.Now()
	Cli.VerifyKey = uuid.New()
	Cli.BackupSer = make(map[string]int)

	//Update Pool info
	poolCli[i] = Cli
	log.Println("Cli ID assign ok:", i)
	return true, i, Cli
}

//start listens
func StartListen(addr string) error {

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	log.Printf("start listener")
	// if Errors accept arrive 100 .listener stop.
	for failures := 0; failures < 100; {
		conn, listenErr := listener.Accept()
		println(conn.RemoteAddr().String())

		if listenErr != nil {
			log.Printf("number:%d,failed listening:%v\n", failures, listenErr)
			failures++
		}
		conn.Write([]byte("欢迎加入聊天组~"))
		if ok, index, Cli := NewConnection_CS(conn); ok {
			// A new connection is established. Spawn a new gorouting to handle that Client.
			go Cli.listenHandle(index)
		}
	}
	return errors.New("Too many listener.Accept() errors,listener stop")
}
