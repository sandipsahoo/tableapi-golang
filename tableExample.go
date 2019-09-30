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

package main

import (
	"encoding/json"
	"github.com/tableapi-golang/custom"
	"github.com/tableapi-golang/table"
	"log"
)

func main() {

	connectionString := "" // use your azure subscription cosmosdb connection string
	tableName := ""        //Enter your table name here

	//Create a Table Service
	ts := table.NewTableService(connectionString)

	//Create Table
	ts.CreateTable(tableName)

	//Insert Entity : create user data model or modify custom data model here
	data := &custom.Data{PartitionKey: "PK", RowKey: "RK"}
	ts.InsertEntity(tableName, data)

	//Fetch Entities
	//{"value":[{"PartitionKey":"PK","RowKey":"RK","Timestamp":"2019-09-23T15:32:45.8355720Z"}]}
	//Use Custom UnMarshal to fetch the custom model entities
	fdata, _ := ts.FetchEntities(tableName, "PK")
	datas := custom.Datas{}
	entitiesJson := string(fdata)
	err := json.Unmarshal([]byte(entitiesJson), &datas)
	if err != nil {
		log.Printf("JSON Marshall Error:*%s", err)
	}
	for _, value := range datas.Value {
		log.Printf("PartitionKey:%s, RowKey:%s", value.PartitionKey, value.RowKey)
	}

	//Fetch an Entity
	fdata, err = ts.FetchEntity(tableName, "PK", "RK")
	if err == nil {
		log.Println(string(fdata))
	} else {
		log.Println(err)
	}

	//Delete Entity
	ts.Delete(tableName, "PK", "RK")

	//Delete Entities
	datas = custom.Datas{}
	ts.DeleteEntities(tableName, "PK", &datas)

	//Delete Table
	ts.DeleteTable(tableName)

}
