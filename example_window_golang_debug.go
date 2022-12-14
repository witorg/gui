package gui

import "log"

import (
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
)

func GolangDebugWindow() {
	var w, t *Node

	Config.Title = "Go Language Debug Window"
	Config.Width = 400
	Config.Height = 400
	Config.Exit = StandardClose
	w = NewWindow()

	t = w.NewTab("Debug Tab")
	log.Println("debugWindow() START")


	///////////////////////////////  Column DEBUG GOLANG   //////////////////////
	g := t.NewGroup("GO Language")

	g.NewButton("runtime.Stack()", func () {
		log.Println("\tSTART")
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		log.Printf("%s", buf)
		log.Println("\tEND")
	})
	g.NewButton("dumpModuleInfo()", func () {
		log.Println("\tSTART")
		dumpModuleInfo()
		log.Println("\tEND")
	})
	g.NewButton("debug.PrintStack()", func () {
		log.Println("\tSTART")
		debug.PrintStack()
		log.Println("\tEND")
	})
	g.NewButton("pprof.Lookup(goroutine)", func () {
		log.Println("\tSTART")
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		log.Println("\tEND")
	})
	g.NewButton("pprof.Lookup(heap)", func () {
		log.Println("\tSTART")
		pprof.Lookup("heap").WriteTo(os.Stdout, 1)
		log.Println("\tEND")
	})
	g.NewButton("pprof.Lookup(block)", func () {
		log.Println("\tSTART")
		pprof.Lookup("block").WriteTo(os.Stdout, 1)
		log.Println("\tEND")
	})
	g.NewButton("pprof.Lookup threadcreate", func () {
		log.Println("\tSTART")
		pprof.Lookup("threadcreate").WriteTo(os.Stdout, 1)
		log.Println("\tEND")
	})
	g.NewButton("runtime.ReadMemStats", func () {
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		log.Printf("alloc: %v bytes\n", s.Alloc)
		log.Printf("total-alloc: %v bytes\n", s.TotalAlloc)
		log.Printf("sys: %v bytes\n", s.Sys)
		log.Printf("lookups: %v\n", s.Lookups)
		log.Printf("mallocs: %v\n", s.Mallocs)
		log.Printf("frees: %v\n", s.Frees)
		log.Printf("heap-alloc: %v bytes\n", s.HeapAlloc)
		log.Printf("heap-sys: %v bytes\n", s.HeapSys)
		log.Printf("heap-idle: %v bytes\n", s.HeapIdle)
		log.Printf("heap-in-use: %v bytes\n", s.HeapInuse)
		log.Printf("heap-released: %v bytes\n", s.HeapReleased)
		log.Printf("heap-objects: %v\n", s.HeapObjects)
		log.Printf("stack-in-use: %v bytes\n", s.StackInuse)
		log.Printf("stack-sys: %v bytes\n", s.StackSys)
		log.Printf("next-gc: when heap-alloc >= %v bytes\n", s.NextGC)
		log.Printf("last-gc: %v ns\n", s.LastGC)
		log.Printf("gc-pause: %v ns\n", s.PauseTotalNs)
		log.Printf("num-gc: %v\n", s.NumGC)
		log.Printf("enable-gc: %v\n", s.EnableGC)
		log.Printf("debug-gc: %v\n", s.DebugGC)
	})
}

func dumpModuleInfo() {
	tmp, _ := debug.ReadBuildInfo()
	if tmp == nil {
		log.Println("This wasn't compiled with go module support")
		return
	}
	log.Println("mod.Path         = ", tmp.Path)
	log.Println("mod.Main.Path    = ", tmp.Main.Path)
	log.Println("mod.Main.Version = ", tmp.Main.Version)
	log.Println("mod.Main.Sum     = ", tmp.Main.Sum)
	for _, value := range tmp.Deps {
		log.Println("\tmod.Path    = ", value.Path)
		log.Println("\tmod.Version = ", value.Version)
	}
}

