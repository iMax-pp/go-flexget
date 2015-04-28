// Copyright (c) 2014-2015 Maxime SIMON. All rights reserved.

package common

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

const (
	confSSHServer  = "flexget.ssh.server"
	confSSHUser    = "flexget.ssh.user"
	confSSHPrivKey = "flexget.ssh.privatekey"
)

// ExecSSHCmd executes a given command via SSH on server "flexget.ssh.server" with user "flexget.ssh.user"
func ExecSSHCmd(cmd string) (string, error) {
	// Init SSH config with user and private key
	config := &ssh.ClientConfig{
		User: Props()[confSSHUser],
		Auth: []ssh.AuthMethod{ssh.PublicKeys(getPrivateKey())},
	}

	// Create a SSH client
	client, err := ssh.Dial("tcp", Props()[confSSHServer], config)
	if err != nil {
		newErr := errors.New("Failed to dial: " + err.Error())
		glog.Error(newErr)
		return "", newErr
	}
	defer client.Close()
	glog.Info("Logged on server ", client.RemoteAddr().String())

	// Open a session, to launch the command
	session, err := client.NewSession()
	if err != nil {
		newErr := errors.New("Failed to create session: " + err.Error())
		glog.Error(newErr)
		return "", newErr
	}
	defer session.Close()
	glog.Info("Session opened on ", client.RemoteAddr().String())

	// Execute command retrieve console output
	var body bytes.Buffer
	session.Stdout = &body
	if err := session.Run(cmd); err != nil {
		newErr := errors.New("Failed to run: " + err.Error())
		glog.Error(newErr)
		return "", newErr
	}
	glog.Info("Command '", cmd, "' executed")

	return body.String(), nil
}

// Retrieve Signer (private key), from path defined in property "flexget.ssh.privatekey"
func getPrivateKey() ssh.Signer {
	keyFile, err := os.Open(Props()[confSSHPrivKey])
	if err != nil {
		fmt.Println(err)
		glog.Fatal(err)
	}
	defer keyFile.Close()

	buf, _ := ioutil.ReadAll(keyFile)
	signer, _ := ssh.ParsePrivateKey(buf)

	return signer
}
