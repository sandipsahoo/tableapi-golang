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

package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/tableapi-golang/constants"
)

// BuildAuthorizationToken generates Authentication Token
func BuildAuthorizationToken(methodType string,
	contentType string,
	utcString string,
	accountName string,
	tableName string,
	authToken string) string {

	canonicalString := fmt.Sprintf(constants.CanonicalStringFormat,
		methodType,
		contentType,
		utcString,
		accountName,
		tableName,
	)

	decodedKey, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		fmt.Println(err)
	}
	hmacSha256 := hmac.New(sha256.New, decodedKey)

	_, err = hmacSha256.Write([]byte(canonicalString))
	if err != nil {
		fmt.Println(err)
	}

	hashPayload := hmacSha256.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(hashPayload)

	authorization := fmt.Sprintf(constants.SharedKeyFormat,
		accountName,
		signature,
	)
	return authorization
}
