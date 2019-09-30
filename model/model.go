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

package model

// Connection model interface used to hold the connection key
type Connection struct {
	ConnectionKey string `json:"ConnectionKey"`
}

// Auth interface model used to hold the connection key
type Auth struct {
	AccountName         string `json:"AccountName"`
	AuthenticationToken string `json:"AuthenticationToken"`
	Endpoint            string `json:"Endpoint"`
}
