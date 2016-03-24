// +build go1.5.1

// Copyright 2014 fabux authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"
	"os"
	"runtime"

	"github.com/benlightning/fabux/fabcore"
)

var CurrentDir, _ = os.Getwd() // 当前操作目录
var (
	timeout int         = 0
	TimeOut int         = 10
	Ok      chan string = make(chan string)
	Err     chan string = make(chan string)
	Res     chan string = make(chan string)
	Sta     chan string = make(chan string)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 配置文件地址，公众命令global，服务器选择1,2,..，是否log
	c, g, h, l := fabcore.Flag()
	config := fabcore.GetConfig(*c)
	host := fabcore.GetHost(config)
	local := fabcore.GetLocal(config)
	var global map[string]string
	//can use global cmd
	if *g {
		global = fabcore.GetGlobal(config)
		//fabcore.Client("127.0.0.1:2222", "root", "vagrant", global)
		//log.Println(global)
	}
	if len(local) > 0 && len(host) > 0 {
		fabcore.Scpsend(h, local, host, Ok, Err, Sta)
	}

	//can print log
	if *l {
		log.Println(len(global))
		log.Println(local)
		log.Println(*l)
		log.Println(*h, host)
	}
}
