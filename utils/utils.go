//
// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
//

package utils

import (
	"bytes"
	"github.com/tableapi-golang/auth"
	"github.com/tableapi-golang/config"
	"github.com/tableapi-golang/constants"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// UtcNow gets time.Now in RFC7231 (HTTP-date) format for x-ms-date
func UtcNow() string {
	return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

// BuildHTTPRequests processes requests and returns the response data
func BuildHTTPRequests(accountName string, authenticationToken string, methodType string, ops string,
	url string, tableName string, buffer *bytes.Buffer) ([]byte, error) {

	request, _ := http.NewRequest(methodType, url, nil)

	if buffer != nil {
		request, _ = http.NewRequest(methodType, url, buffer)
	}

	utcString := UtcNow()

	config.BuildRequestHeaders(request, UtcNow(), ops)

	request.Header.Set(constants.AUTHORIZATION, auth.BuildAuthorizationToken(methodType,
		constants.ContentType, utcString, accountName, tableName, authenticationToken))
	client := &http.Client{}

	response, err := client.Do(request)

	var data []byte

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Response error %s\n", err)
		}
	}

	return data, err
}
