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

package fabcore

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func Client(ip_port, user, password string, command map[string]string) {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, err := ssh.Dial("tcp", ip_port, &Conf)
	if err == nil {
		fmt.Println(ip_port, "connect status:success")
	} else {
		fmt.Print(ip_port, "connect error:", err)
		return
	}
	defer Client.Close()
	for _, cmd := range command {
		if session, err := Client.NewSession(); err == nil {
			defer session.Close()
			var Result bytes.Buffer
			session.Stderr = &Result
			session.Stdout = &Result
			err = session.Run(cmd)
			if err == nil {
				fmt.Println(ip_port, "run command:", cmd, "run status:Ok")
			} else {
				fmt.Println(ip_port, "run error:", err)
			}
			fmt.Println(ip_port, "run result:\n", Result.String())
		}
	}
}
