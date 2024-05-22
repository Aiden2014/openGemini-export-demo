# openGemini-export-demo

A demo to export openGemini. Reference: https://github.com/openGemini/openGemini/pull/287

## Usage

Fill in the fields listed in main.go

```go
	databases := ""  // ie. "db0" "db0,NOAA_water_database"
	exportPath := "" // ie. "/tmp/output"
	outPutPath := exportPath + "/" + "export.txt"
	actualDataPath := "" // ie. "/tmp/openGemini/data/data"
	actualWalPath := ""  // ie. "/tmp/openGemini/data/wal"
```

Run

```
go run ./main.go
```