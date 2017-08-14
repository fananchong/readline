package readline

import (
	"fmt"
	"net"
	"net/rpc"
	"strings"
)

type Args struct {
	A []string
}

type Reply struct {
	R int
}

type Scanner struct {
	line   string
	err    error
	server *rpc.Server
	client *rpc.Client
	addr   string
}

func (this *Scanner) Scan(prompt string) bool {

	fmt.Print(prompt)

	var c byte
	var err error
	var b []byte
	for {
		_, err = fmt.Scanf("%c", &c)
		if err == nil && c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	this.line = string(b)
	this.err = err
	return err == nil
}

func (this *Scanner) Text() string {
	return this.line
}

func (this *Scanner) Err() error {
	return this.err
}

func (this *Scanner) initServer() {
	if this.server == nil {
		server := rpc.NewServer()
		var l net.Listener
		l, e := net.Listen("tcp", "127.0.0.1:0")
		if e != nil {
			fmt.Errorf("net.Listen tcp :0: %v", e)
			return
		}
		fmt.Println("NewServer test RPC server listening on", l.Addr().String())

		go func() {
			this.server.Accept(l)
		}()

		this.server = server
		this.addr = l.Addr().String()
	}
}

func (this *Scanner) Register(rcvr interface{}) {
	this.initServer()
	this.server.Register(rcvr)
}

func (this *Scanner) RegisterName(name string, rcvr interface{}) {
	this.initServer()
	this.server.RegisterName(name, rcvr)
}

func (this *Scanner) Callback(clas string, cmd string) {
	if this.client == nil {
		client, err := rpc.Dial("tcp", this.addr)
		if err != nil {
			fmt.Errorf("rpc.Dial tcp :0: %v", err)
		}
		this.client = client
	}
	var err error
	params := strings.Split(cmd, " ")
	if len(params) <= 0 {
		return
	} else if len(params) == 1 {
		err = this.client.Call(clas+"."+strings.Title(params[0]), &Args{[]string{}}, &Reply{})
	} else {
		err = this.client.Call(clas+"."+strings.Title(params[0]), &Args{params[1:]}, &Reply{})
	}
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
