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

package custom

import "time"

// Data model interface used when fetching multiple entities
type Data struct {
	PartitionKey string    `json:"PartitionKey"`
	RowKey       string    `json:"RowKey"`
	Timestamp    time.Time `json:"Timestamp"`
}

// Datas model interface used when fetching multiple entities
type Datas struct {
	Value []struct {
		PartitionKey string    `json:"PartitionKey"`
		RowKey       string    `json:"RowKey"`
		Timestamp    time.Time `json:"Timestamp"`
	} `json:"value"`
}
