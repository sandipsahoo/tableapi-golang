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

package config

import (
	"github.com/tableapi-golang/constants"
	"net/http"
)

// BuildRequestHeaders builds Request Headers for http request using utc string and operation type
func BuildRequestHeaders(request *http.Request, utcString string, ops string) {
	request.Header.Set("Content-Type", constants.ContentType)
	request.Header.Add("MaxDataServiceVersion", "3.0;NetFx")
	request.Header.Add("DataServiceVersion", "3.0;")
	request.Header.Add("x-ms-date", utcString)

	switch ops {
	case constants.CREATE:
		request.Header.Add("Prefer", "return-no-content")
		request.Header.Add("Accept", "application/json;odata=minimalmetadata")
		request.Header.Add("Content-Length", "1135")
		break
	case constants.INSERT:
		request.Header.Add("Prefer", "return-no-content")
		break
	case constants.FETCH:
		request.Header.Add("Prefer", "return-content")
		request.Header.Add("Accept", "application/json;odata=nometadata")
		break
	case constants.DELETE:
		request.Header.Add("Prefer", "return-no-content")
		request.Header.Add("Accept", "application/json;odata=nometadata")
		request.Header.Add("If-Match", "*")
		break
	}

}
