// 两种方式，一种是读取zip文件到缓冲(连接管道StdinPipe())，然后通过在服务器端执行scp -qrt /dir/存储
// 第二种是本地使用scp或pscp(win)命令进行文件传输
package fabcore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Scpsend(hostlevel *string, local map[string]string, host map[string][]string, Ok, Err, Sta chan string) {
	OldPath, _ := os.Getwd()
	//var file os.FileInfo
	var e error
	var SendFilePath string
	var SendName string
	for name, dir := range local {
		if _, e = os.Stat(dir); e == nil {
			Files2Zip(dir, name)
			SendFilePath = fmt.Sprintf("%s%s%s.zip", OldPath, string(filepath.Separator), name)
			SendName = fmt.Sprintf("%s.zip", name)
			if *hostlevel == "" {
				for i, _ := range host {
					if i != "0" {
						go connect(host[i][0], host[i][1], host[i][2], host[i][3], SendFilePath, SendName, Ok, Err, Sta)
					}
				}
			} else {
				FlagList := strings.Split(*hostlevel, ",")
				for _, i := range FlagList {
					go connect(host[i][0], host[i][1], host[i][2], host[i][3], SendFilePath, SendName, Ok, Err, Sta)
				}
			}
		} else {
			fmt.Println(dir, ":not exists!")
		}
	}
}

func connect(ip_port, user, password, FilePath, Name, RemoteDir string, Ok, Err, Sta chan string) {
	Auth := []ssh.AuthMethod{ssh.Password(password)}
	conf := ssh.ClientConfig{Auth: Auth, User: user}
	Client, err := ssh.Dial("tcp", ip_port, &conf)
	if err == nil {
		Ok <- fmt.Sprint(ip_port, "connect status:success")
	} else {
		Err <- fmt.Sprint(ip_port, "connect error:", err)
	}
	defer Client.Close()

	conn, err1 := Client.NewSession()
	if err1 == nil {
		Ok <- fmt.Sprint(ip_port, "session status:Ok")
	} else {
		Err <- fmt.Sprint(ip_port, "session error:", err1)
	}
	defer conn.Close()
	go func() {
		Buf := make([]byte, 1024)
		w, _ := conn.StdinPipe()
		defer w.Close()
		File, _ := os.Open(FilePath)
		defer File.Close()
		info, _ := File.Stat()
		fmt.Fprintln(w, "C0644", info.Size(), Name)
		for {
			n, err := File.Read(Buf)
			fmt.Fprint(w, string(Buf[:n]))
			if err != nil {
				if err == io.EOF {
					return
				} else {
					panic(err)
				}
			}
		}
	}()
	err3 := conn.Run("scp -qrt " + RemoteDir) //upload to remote directory
	if err3.Error() == "Process exited with: 1. Reason was:  ()" {
		Ok <- fmt.Sprint(ip_port, "execute status:Ok")
		Sta <- "Ok"
	} else {
		Err <- fmt.Sprint(ip_port, "execute error:", err3)
	}
}
