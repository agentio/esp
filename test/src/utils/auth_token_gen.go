// Copyright (C) Endpoints Server Proxy Authors
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
////////////////////////////////////////////////////////////////////////////////
//
package utils

import (
	"os/exec"
	"log"
	"regexp"
	"fmt"
)


func GenAuthToken(secretFile string, audience string) (string, error) {
	var testBinRoot, err = GetTestBinRootPath()
	if err != nil {
		log.Printf("Failed to get the test env path. %s", err)
		return "", err
	}
	var binPath = testBinRoot + "/src/tools/auth_token_gen"
	output, err := exec.Command(binPath, secretFile, audience).Output()

	if err != nil {
		log.Printf("Failed to generate auth token. %s", err)
		return "", err
	}

	re := regexp.MustCompile("Auth token:?")
	outputStr := string(output)
	fmt.Println(outputStr)

	loc := re.FindStringIndex(outputStr)
	authToken := outputStr[loc[1]+1 : len(outputStr)-1]

	return authToken, nil
}
