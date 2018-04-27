package debug

import (
	"os"
	"os/signal"
	"syscall"
	"runtime"
	"log"
	"time"
)

var (
	dumpFile string = "stack.log"
	ch chan os.Signal
)

func SetupStackTrap(args ...string) {
	if len(args) > 1 {
		dumpFile = args[1]
	}
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	go func() {
		for range ch {
			dumpStack()
		}
	}()
}

func StartStack() {
	ch <- syscall.SIGHUP
}


func dumpStack() {
	buf := make([]byte, 1024000)
	buf = buf[:runtime.Stack(buf, true)]

	f, err := os.OpenFile(dumpFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString("\n\n")
	f.WriteString(now + ", stdout " + "\n")
	f.Write(buf)
}