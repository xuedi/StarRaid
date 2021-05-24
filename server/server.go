package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	tps          int
	tickSleep    time.Duration
	name         string
	secondTarget time.Time
	config       Config
	objects      Objects
	users        Users
	net          Network
}

// Init register interrupt signals and timers
func (s *Server) Init() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Println("End server with CTRL+C")
	go func() {
		<-sigChan
		fmt.Println("\n... like tears in rain. Time to die")
		os.Exit(1)
	}()

	// create config file
	s.config = Config{}
	s.config.load("assets/config.ini")

	// first target way in the past
	s.secondTarget = time.Unix(0, 0)
	s.tickSleep = 10

	// load once database objects
	s.objects.Load(s.config.dsn)
	s.users.Load(s.config.dsn)

	// get the network started
	s.net = Network{}
	s.net.RegisterObjects(s.objects)
	s.net.load(s.config.host, s.config.port)
}

// Tick Should be in a separate go routine and updating via channel
func (s *Server) Tick() {
	s.objects.Tick()
	s.TickSecond()
	s.tps++
	time.Sleep(s.tickSleep * time.Millisecond) // give the cpu some rest
}

// TickSecond Just some debugging and updating the load/rest deltas
func (s *Server) TickSecond() {
	if time.Now().Before(s.secondTarget) {
		return
	}

	if s.tps >= s.config.targetTps {
		if s.tickSleep < 10 {
			s.tickSleep++
		}
	} else {
		if s.tickSleep > 1 {
			s.tickSleep--
		}
	}

	fmt.Println("Ticks:", s.tps, "Sleep:", s.tickSleep, "Objects:", s.objects.Count(), "Users:", s.users.Count())
	s.secondTarget = time.Now().Add(time.Second)
	s.tps = 0
}
