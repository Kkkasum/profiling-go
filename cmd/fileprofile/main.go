package main

import (
	"context"
	"log"
	"os"
	"runtime/pprof"

	"profiling/factorial"
)

func main() {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}

	if err = pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}

	defer func() {
		pprof.StopCPUProfile()
	}()

	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("label", "calculate"), func(ctx context.Context) {
		factorial.Calculate()
	})

	// профайлинг памяти производится в конце
	memFile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	if err = pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal(err)
	}
}
