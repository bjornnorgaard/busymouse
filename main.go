package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/bjornnorgaard/busymouse/config"
	"github.com/go-vgo/robotgo"
)

func main() {
	err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	if config.GetDebug() {
		slog.Info("debug is enabled")
		moveMouse()
		return
	}

	ticker := time.NewTicker(config.GetInterval())
	for range ticker.C {
		moveMouse()
	}
}

func moveMouse() {
	slog.Info("moving mouse")

	robotgo.MouseSleep = 300

	robotgo.Move(100, 100)
	fmt.Println(robotgo.Location())
	robotgo.Move(100, -200) // multi screen supported
	robotgo.MoveSmooth(120, -150)
	fmt.Println(robotgo.Location())

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)
	// robotgo.ScrollRelative(10, -100)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)

	robotgo.MoveSmooth(100, 200, 1.0, 10.0)

}
