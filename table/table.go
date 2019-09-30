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

package table

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tableapi-golang/auth"
	"github.com/tableapi-golang/constants"
	"github.com/tableapi-golang/custom"
	"github.com/tableapi-golang/model"
	"github.com/tableapi-golang/utils"
	"log"
)

var connectionString = ""
var authn = &model.Auth{}

// TableService model interface contains AccountName, AuthenticationToken, Endpoint
type TableService struct {
	AccountName         string `json:"AccountName"`
	AuthenticationToken string `json:"AuthenticationToken"`
	Endpoint            string `json:"Endpoint"`
}

// NewTableService object instantiated with AccountName, AuthenticationToken, Endpoint
func NewTableService(connectionString string) *TableService {
	authn = auth.GetAuth(connectionString)
	return &TableService{authn.AccountName,
		authn.AuthenticationToken, authn.Endpoint}
}

// CreateTable a Table Service method creates a table
func (ts TableService) CreateTable(tableNameParam string) ([]byte, error) {

	log.Println("*Creating Table.")

	jsonData := map[string]string{"TableName": "" + tableNameParam + ""}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		log.Println(err)
	}

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_POST, constants.CREATE, ts.Endpoint+"Tables", "Tables", bytes.NewBuffer(jsonValue))

	log.Println("Done!")

	return data, err
}

// InsertEntity method inserts entity to the cosmosdb database
func (ts TableService) InsertEntity(tableNameParam string, jsonData interface{}) ([]byte, error) {

	log.Println("*Inserting Entity")

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		log.Println(err)
	}

	tableName := fmt.Sprintf("%s()", tableNameParam)

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_POST, constants.INSERT, ts.Endpoint+tableName, tableName, bytes.NewBuffer(jsonValue))

	log.Println("Done!")

	return data, err
}

// FetchEntity fetches entity from cosmosdb database
func (ts TableService) FetchEntity(tableNameParam string, partitionKey string, rowKey string) ([]byte, error) {

	log.Println("*Fetching Entity")

	tableName := fmt.Sprintf("%s(PartitionKey='%s',RowKey='%s')", tableNameParam, partitionKey, rowKey)

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_GET, constants.FETCH, ts.Endpoint+tableName, tableName, nil)

	log.Println("Done!")

	return data, err
}

// FetchEntities fetches entities from database
func (ts TableService) FetchEntities(tableNameParam string, partitionKey string) ([]byte, error) {

	log.Println("*Fetching Entities")

	tableName := fmt.Sprintf("%s()", tableNameParam)
	tableURI := fmt.Sprintf("%s()?$filter=PartitionKey%seq%s'%s'", tableNameParam, "%20", "%20", partitionKey)

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_GET, constants.FETCH, ts.Endpoint+tableURI, tableName, nil)

	log.Println("Done!")

	return data, err
}

// Delete removes entity from the database
func (ts TableService) Delete(tableNameParam string, partitionKey string, rowKey string) ([]byte, error) {

	log.Println("*Deleting Entity")

	tableName := fmt.Sprintf("%s(PartitionKey='%s',RowKey='%s')", tableNameParam, partitionKey, rowKey)

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_DELETE, constants.DELETE, ts.Endpoint+tableName, tableName, nil)

	log.Println("Done!")

	return data, err
}

// DeleteTable removes Table from the database
func (ts TableService) DeleteTable(tableNameParam string) ([]byte, error) {

	log.Println("*Deleting Table.")

	tableName := fmt.Sprintf("Tables('%s')", tableNameParam)

	data, err := utils.BuildHTTPRequests(ts.AccountName, ts.AuthenticationToken,
		constants.HTTP_DELETE, constants.DELETE, ts.Endpoint+tableName, tableName, nil)

	log.Println("Done!")

	return data, err
}

// DeleteEntities removes entities from the database
func (ts TableService) DeleteEntities(tableName string, partitionKey string, result *custom.Datas) error {
	log.Printf("*Deleting Entities of table '%s' for PartitionKey '%s'", tableName, partitionKey)
	data, _ := ts.FetchEntities(tableName, partitionKey)
	entitiesJson := string(data)
	err := json.Unmarshal([]byte(entitiesJson), &result)
	if err != nil {
		log.Printf("JSON Marshall Error:*%s", err)
		return err
	}
	for _, value := range result.Value {
		_, err = ts.Delete(tableName, partitionKey, value.RowKey)
		if err != nil {
			fmt.Printf("Error Occurred:*%s", err)
			break
		}
	}
	return err
}
