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

package main

import (
	"github.com/tableapi-golang/custom"
	"github.com/tableapi-golang/table"
	"testing"
)

var connectionString = "" // use your azure subscription cosmosdb connection string
var tableName = ""        //Enter your table name here

// TestCreateTable tests Create Table
func TestCreateTable(t *testing.T) {
	ts := table.NewTableService(connectionString)

	expected := true
	actual := true

	response, err := ts.CreateTable(tableName)
	if err != nil || len(response) > 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table creation of %s to be %t but instead got %t!", tableName, expected, actual)
	}

}

// TestInsertEntity tests Insert Entity
func TestInsertEntity(t *testing.T) {
	ts := table.NewTableService(connectionString)

	data := &custom.Data{PartitionKey: "PK", RowKey: "RK"}

	expected := true
	actual := true

	response, err := ts.InsertEntity(tableName, data)
	if err != nil || len(response) > 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table Data Insertion of %s to be %t but instead got %t!", tableName, expected, actual)
	}

}

// TestFetchEntity tests Fetch Entity
func TestFetchEntity(t *testing.T) {
	ts := table.NewTableService(connectionString)

	partitionKey := "PK"
	rowKey := "RK"

	expected := true
	actual := true

	response, err := ts.FetchEntity(tableName, partitionKey, rowKey)
	if err != nil || len(response) == 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table fetching data of %s to be %t but instead got %t!", tableName, expected, actual)
	}

}

// TestFetchEntities tests Fetch Entities
func TestFetchEntities(t *testing.T) {
	ts := table.NewTableService(connectionString)

	partitionKey := "PK"

	expected := true
	actual := true

	response, err := ts.FetchEntities(tableName, partitionKey)
	if err != nil || len(response) == 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table data fetching of %s to be %t but instead got %t!", tableName, expected, actual)
	}

}

// TestDeleteEntity tests Delete Entity
func TestDeleteEntity(t *testing.T) {
	ts := table.NewTableService(connectionString)

	partitionKey := "PK"
	rowKey := "RK"

	expected := true
	actual := true

	data, err := ts.Delete(tableName, partitionKey, rowKey)
	if err != nil || len(data) > 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table delete entity of %s "+
			"with Partition Key:%s and RowKey:%s to be %t but instead got %t!", tableName, partitionKey, rowKey, expected, actual)
	}

}

// TestDeleteEntities tests Delete Entities
func TestDeleteEntities(t *testing.T) {
	ts := table.NewTableService(connectionString)

	partitionKey := "PK"

	datas := custom.Datas{}

	expected := true
	actual := true

	err := ts.DeleteEntities(tableName, partitionKey, &datas)
	if err != nil {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table delete entity of %s "+
			"with Partition Key:%s to be %t but instead got %t!", tableName, partitionKey, expected, actual)
	}

}

// TestDeleteTable tests Delete Table
func TestDeleteTable(t *testing.T) {
	ts := table.NewTableService(connectionString)

	expected := true
	actual := true

	data, err := ts.DeleteTable(tableName)
	if err != nil || len(data) > 0 {
		actual = false
	}

	if actual != expected {
		t.Errorf("Expected table deletion of %s to be %t but instead got %t!", tableName, expected, actual)
	}

}
