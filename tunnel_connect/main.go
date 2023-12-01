package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (dial *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return dial.client.Dial("tcp", addr)
}

func main() {

	sshHost := "" // SSH Server Hostname/IP
	sshPort := 22 // SSH Port
	sshUser := "" // SSH Username
	sshPass := "" // Empty string for no password
	dbUser := ""  // DB username
	dbPass := ""  // DB Password
	dbHost := ""  // DB Hostname/IP
	dbName := ""  // Database name

	var agentClient agent.Agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		agentClient = agent.NewClient(conn)
	}
	sshConfig := &ssh.ClientConfig{
		User:            sshUser,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	if sshPass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPass, nil
		}))
	}
	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig)
	if err == nil {
		defer sshcon.Close()
		fmt.Println("connect ssh server")
		dial := (&ViaSSHDialer{sshcon}).Dial
		mysql.RegisterDialContext("mysql+tcp", func(_ context.Context, addr string) (net.Conn, error) {
			return dial(addr)
		})
		fmt.Println("mysql register dial context")
		if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {

			fmt.Printf("Successfully connected to the db\n")

			if rows, err := db.Query("show tables"); err == nil {
				for rows.Next() {
					var name string
					rows.Scan(&name)
					fmt.Printf("table name :%s", name)
				}
				rows.Close()
			} else {
				fmt.Printf("Failure: %s", err.Error())
			}

			db.Close()

		} else {

			fmt.Printf("Failed to connect to the db: %s\n", err.Error())
		}

	}
}
