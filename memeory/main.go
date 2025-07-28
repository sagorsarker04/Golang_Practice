package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	runtime.MemProfileRate = 2048
	profileFile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("Failed to create memory profile file: ", err)
	}
	defer profileFile.Close()
	generateDataStructures()
	generateDataStructures2()
	if err := pprof.WriteHeapProfile(profileFile); err != nil {
		log.Fatal("Unable to write memory profile: ", err)
	}
	log.Println("Memory profile saved to mem.prof")
}

func generateDataStructures() []map[int]string {
	structures := make([]map[int]string, 800000)
	for i := range structures {
		structures[i] = map[int]string{
			0: "A", 1: "B", 2: "C", 3: "D", 4: "E",
		}
	}
	return structures
}
func generateDataStructures2() []map[int]string {
	structures := make([]map[int]string, 80000)
	for i := range structures {
		structures[i] = map[int]string{
			0: "A", 1: "B", 2: "C", 3: "D", 4: "E",
		}
	}
	return structures
}

// Run: go run main.go
// Then: go tool pprof -http=:8080 mem.prof