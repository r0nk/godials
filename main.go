//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var run = true

func write_state(v1 float64) {
	str := fmt.Sprintf("%0.4v\n", v1)
	if err := os.WriteFile("/tmp/godials.txt", []byte(str), 0666); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	rand.Seed(time.Now().UTC().UnixNano())
	pc := widgets.NewPieChart()
	pc.Title = "godials"
	pc.SetRect(5, 5, 30, 20)
	pc.Data = []float64{.15, .85}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("")
	}

	ui.Render(pc)

	uiEvents := ui.PollEvents()
	//	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<MouseWheelUp>":
				pc.AngleOffset += 0.2
				pc.Title = fmt.Sprintf("%0.4v", pc.AngleOffset)
				ui.Render(pc)
			case "<MouseWheelDown>":
				pc.Title = fmt.Sprintf("%0.4v", pc.AngleOffset)
				pc.AngleOffset -= 0.2
				ui.Render(pc)
			}
		}
		write_state(pc.AngleOffset)
	}
}
