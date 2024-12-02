package main

import (
	"context"
	"github.com/rgzr/sshtun"
	"log"
	"time"
)

func tunel() {
	sshTun := sshtun.New(5433, "sviatkyou.cv.ua", 5432)
	sshTun.SetPassword("Thegovernmentsucks1488!")
	sshTun.SetUser("sviatkyo")
	sshTun.SetPort(21098)

	// We print each tunneled state to see the connections status
	sshTun.SetTunneledConnState(func(tun *sshtun.SSHTun, state *sshtun.TunneledConnState) {
		log.Printf("%+v", state)
	})

	// We set a callback to know when the tunnel is ready
	sshTun.SetConnState(func(tun *sshtun.SSHTun, state sshtun.ConnState) {
		switch state {
		case sshtun.StateStarting:
			log.Printf("STATE is Starting")
		case sshtun.StateStarted:
			log.Printf("STATE is Started")
		case sshtun.StateStopped:
			log.Printf("STATE is Stopped")
		}
	})

	// We start the tunnel (and restart it every time it is stopped)
	go func() {
		if err := sshTun.Start(context.Background()); err != nil {
			log.Printf("SSH tunnel error: %v", err)
			time.Sleep(time.Second) // don't flood if there's a start error :)
		}
	}()
}
