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
	"github.com/tableapi-golang/constants"
	"github.com/tableapi-golang/model"
	"strings"
)

// GetAuth extracts the account name, endpoint and authentication token
func GetAuth(connectionString string) *model.Auth {
	auth := &model.Auth{}
	accountName, endpoint, authenticationToken := fetchCredentials(connectionString)
	auth.AccountName = accountName
	auth.Endpoint = endpoint
	auth.AuthenticationToken = authenticationToken
	return auth
}

// fetchCredentials processes the connection string
func fetchCredentials(connectionString string) (string, string, string) {

	endpoint := ""
	authenticationToken := ""
	accountName := ""

	connTokens := strings.Split(connectionString, ";")
	for _, val := range connTokens {
		if strings.HasPrefix(val, constants.AccountName) {
			accountName = strings.Replace(val, constants.AccountNamePrefix, "", -1)
		}
		if strings.HasPrefix(val, constants.AccountKey) {
			authenticationToken = strings.Replace(val, constants.AccountKeyPrefix, "", -1)
		}
		if strings.HasPrefix(val, constants.Endpoint) {
			endpoint = strings.Replace(val, constants.EndpointPrefix, "", -1)
		}
	}
	return accountName, endpoint, authenticationToken
}
