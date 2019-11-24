package main

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	rpch "github.com/tidusant/chadmin-repo/cuahang"
	"github.com/tidusant/chadmin-repo/models"

	//	"c3m/common/inflect"
	//	"c3m/log" test dev branch 3 dev3

	"flag"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

type Arith int

func (t *Arith) Run(data string, result *models.RequestResult) error {
	*result = models.RequestResult{}
	//parse args
	args := strings.Split(data, "|")

	if len(args) < 3 {
		return nil
	}
	var usex models.UserSession
	usex.Session = args[0]
	usex.Action = args[2]
	info := strings.Split(args[1], "[+]")
	if len(info) < 2 {
		return nil
	}
	usex.UserID = info[0]
	ShopID := info[1]
	usex.Params = ""
	if len(args) > 3 {
		usex.Params = args[3]
	}
	//check shop permission
	shop := rpch.GetShopById(usex.UserID, ShopID)
	if shop.Status == 0 {
		*result = c3mcommon.ReturnJsonMessage("-4", "Shop is disabled.", "", "")
		return nil
	}
	usex.Shop = shop

	if usex.Action == "s" {
		*result = Save(usex)
	} else if usex.Action == "la" {
		*result = LoadAll(usex)
	}

	return nil
}

func main() {
	var port int
	var debug bool
	flag.IntVar(&port, "port", 9892, "help message for flagname")
	flag.BoolVar(&debug, "debug", false, "Indicates if debug messages should be printed in log files")
	flag.Parse()

	logLevel := log.DebugLevel
	if !debug {
		logLevel = log.InfoLevel

	}

	log.SetOutputFile(fmt.Sprintf("resource-"+strconv.Itoa(port)), logLevel)
	defer log.CloseOutputFile()
	log.RedirectStdOut()

	//init db
	arith := new(Arith)
	rpc.Register(arith)
	log.Infof("running with port:" + strconv.Itoa(port))

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	c3mcommon.CheckError("rpc dail:", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	c3mcommon.CheckError("rpc init listen", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
