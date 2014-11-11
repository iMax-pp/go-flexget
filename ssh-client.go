// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package main

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

// Execute a given command via SSH on server "flexget.ssh.server" with user "flexget.ssh.user"
func ExecSSHCmd(cmd string) (string, error) {
	// Init SSH config with user and private key
	config := &ssh.ClientConfig{
		User: props["flexget.ssh.user"],
		Auth: []ssh.AuthMethod{ssh.PublicKeys(getPrivateKey())},
	}

	// Create a SSH client
	client, err := ssh.Dial("tcp", props["flexget.ssh.server"], config)
	if err != nil {
		newErr := errors.New("Failed to dial: " + err.Error())
		logger.Error(newErr)
		return "", newErr
	}
	defer client.Close()
	logger.Debug("Logged on server", client.RemoteAddr().String())

	// Open a session, to launch the command
	session, err := client.NewSession()
	if err != nil {
		newErr := errors.New("Failed to create session: " + err.Error())
		logger.Error(newErr)
		return "", newErr
	}
	defer session.Close()
	logger.Debug("Session opened on", client.RemoteAddr().String())

	// Execute command retrieve console output
	var body bytes.Buffer
	session.Stdout = &body
	if err := session.Run(cmd); err != nil {
		newErr := errors.New("Failed to run: " + err.Error())
		logger.Error(newErr)
		return "", newErr
	}
	logger.Debug("Command '", cmd, "' executed")

	return body.String(), nil
}

// Retrieve Signer (private key), from path defined in property "flexget.ssh.privatekey"
func getPrivateKey() ssh.Signer {
	keyFile, err := os.Open(props["flexget.ssh.privatekey"])
	if err != nil {
		logger.Fatal(err)
	}
	defer keyFile.Close()

	buf, _ := ioutil.ReadAll(keyFile)
	signer, _ := ssh.ParsePrivateKey(buf)

	return signer
}
