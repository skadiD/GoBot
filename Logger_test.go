package main

import (
	"fmt"
	"gobot/cmd/globals"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	b.ResetTimer()
	Log := globals.Logger.Create()
	for n := 0; n < b.N; n++ {
		go fmt.Printf(Log.Warn("t"))
		//fmt.Print("test154")
		//Log.Warn("test154")
		//os.Stdout.WriteString("test154")
		//os.Stdout.Write([]byte{96, 96, 96, 96})
	}
}
