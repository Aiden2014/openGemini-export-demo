package main

import (
	"flag"
	"fmt"
	"github.com/openGemini/openGemini/engine/index/tsi"
	"log"
	"os"
	"time"
)

func main() {
	flag.Parse()
	// TODO: Use toml or commands instead of hard coding
	databases := ""  // ie. "db0" "db0,NOAA_water_database"
	exportPath := "" // ie. "/tmp/output"
	outPutPath := exportPath + "/" + "export.txt"
	actualDataPath := "" // ie. "/tmp/openGemini/data/data"
	actualWalPath := ""  // ie. "/tmp/openGemini/data/wal"

	filter := &LineFilter{
		startTime: 0,
		endTime:   time.Now().UnixNano(),
	}

	e := &Exporter{
		stderrLogger:                    log.New(os.Stderr, "uint_test", log.LstdFlags),
		stdoutLogger:                    log.New(os.Stdout, "uint_test", log.LstdFlags),
		manifest:                        make(map[string]struct{}),
		rpNameToMeasurementTsspFilesMap: make(map[string]map[string][]string),
		rpNameToIdToIndexMap:            make(map[string]map[uint64]*tsi.MergeSetIndex),
		rpNameToWalFilesMap:             make(map[string][]string),
		Stderr:                          os.Stderr,
		Stdout:                          os.Stdout,
		actualDataPath:                  actualDataPath,
		actualWalPath:                   actualWalPath,
		databases:                       databases,
		outPutPath:                      outPutPath,
		filter:                          filter,
	}
	err := e.Export()
	if err != nil {
		fmt.Print(err)
	}
}
