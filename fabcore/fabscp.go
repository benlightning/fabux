package fabcore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

func Scpsend(hlevel, Path *string, config map[string][]string, Ok, Err, Sta chan string) {
	FlagList := strings.Split(*hlevel, ",")
	Name := filepath.Base(*Path)
	SendFilePath := *Path
	OldPath, _ := os.Getwd()
	if File, e := os.Stat(SendFilePath); e == nil {
		if File.IsDir() {
			//Zip_File.AddFilesToZip(SendFilePath, Name)
			SendFilePath = fmt.Sprintf("%s\\%s.zip", OldPath, Name)
			Name = fmt.Sprintf("%s.zip", Name)
		}
	} else {
		fmt.Println(SendFilePath, ":not exists!")
		os.Exit(2)
	}
	fmt.Println("send the file:", Name)
	if *hlevel == "" {
		for i, _ := range config {
			if i != "0" {
				go connect(config[i][0], config[i][1], config[i][2], SendFilePath, Name, Ok, Err, Sta)
			}
		}
	} else {
		for _, i := range FlagList {
			go connect(config[i][0], config[i][1], config[i][2], SendFilePath, Name, Ok, Err, Sta)
		}
	}
}

func connect(ip, user, password, FilePath, Name string, Ok, Err, Sta chan string) {
	Auth := []ssh.AuthMethod{ssh.Password(password)}
	conf := ssh.ClientConfig{Auth: Auth, User: user}
	Client, err := ssh.Dial("tcp", ip, &conf)
	if err == nil {
		Ok <- fmt.Sprint(ip, "connect status:success")
	} else {
		Err <- fmt.Sprint(ip, "connect error:", err)
	}
	defer Client.Close()

	conn, err1 := Client.NewSession()
	if err1 == nil {
		Ok <- fmt.Sprint(ip, "session status:Ok")
	} else {
		Err <- fmt.Sprint(ip, "session error:", err1)
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
	err3 := conn.Run("scp -qrt /mnt/") //upload to /mnt/ directory
	if err3.Error() == "Process exited with: 1. Reason was:  ()" {
		Ok <- fmt.Sprint(ip, "execute status:Ok")
		Sta <- "Ok"
	} else {
		Err <- fmt.Sprint(ip, "execute error:", err3)
	}
}
