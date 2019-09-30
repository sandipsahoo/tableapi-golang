[![godoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/sandipsahoo/tableapi-golang/table) [![Go Report Card](https://goreportcard.com/badge/github.com/sandipsahoo/tableapi-golang)](https://goreportcard.com/report/github.com/sandipsahoo/tableapi-golang) [![License: Apache2.0](https://img.shields.io/badge/License-Apache2.0-green.svg)](https://opensource.org/licenses/Apache-2.0)[![GitHub release](https://img.shields.io/badge/release-v1.0--alpha-yellowgreen)](https://github.com/sandipsahoo/tableapi-golang/releases/)

# TableAPIGolang
This is GO SDK for CosmosDB using Table API

### Prerequisites

```
golang 1.10 or above
```

### Installing

```
go get github.com/sasahoo/tableapi-golang/table

```

## Usage

### Create a Table Service

```
connectionString := "" // Enter your azure subscription cosmosdb connection string

ts := table.NewTableService(connectionString)
```
### Create a Table

```
tableName := "" //Enter your cosmosdb table name

ts.CreateTable(tableName)

```
### Insert Table Entity
Create your own data model or modify Custom data model given in the repo

```
data := &custom.Data{PartitionKey: "PK", RowKey: "RK"}

ts.InsertEntity(tableName, data)

```

### Fetch Table Entity
Use Custom UnMarshal to fetch the custom model entities
```
fdata, err = ts.FetchEntity(tableName, "PK", "RK")
	if err == nil {
		log.Println(string(fdata))
	} else {
		log.Println(err)
	}

```
### Fetch Table Entities
Use Custom UnMarshal to fetch the custom model entities

```
fdata, _ := ts.FetchEntities(tableName, "PK")
	datas := custom.Datas{}
	entitiesJson := string(fdata);
	err := json.Unmarshal([]byte(entitiesJson), &datas)
	if err != nil {
		log.Printf("JSON Marshall Error:*%s", err)
	}
	for _, value := range datas.Value {
		log.Printf("PartitionKey:%s, RowKey:%s", value.PartitionKey, value.RowKey)
	}

```
### Delete Table Entity

```
ts.Delete(tableName, "PK", "RK")

```
### Delete Table Entities

```
datas = custom.Datas{}
err = ts.DeleteEntities(tableName, "PK", &datas)

```

### Delete Table

```
ts.DeleteTable(tableName)

```
## Contributing

Please read [CONTRIBUTING.md](https://github.com/sandipsahoo/tableapi-golang/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/sandipsahoo/tableapi-golang/tags). 

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](https://github.com/sandipsahoo/tableapi-golang/blob/master/LICENSE) file for details

## Acknowledgments

* Azure Storage Services REST API
* Azure CosmosDB


