package main

import (
	"log"

	"github.com/TheBitDrifter/coldbrew"
)

func main() {
	client := coldbrew.NewClient()
	client.SetWindowSize(1920, 1080)
	client.SetResolution(640, 360)
	client.SetResizable(true)
	client.SetTitle("Bappa Example!")

	client.NewScene("scene one", 3000, 3000, sceneOnePlan, []coldbrew.RenderSystem{})
	client.NewScene("scene two", 3000, 3000, sceneTwoPlan, []coldbrew.RenderSystem{})
	client.NewScene("scene three", 3000, 3000, sceneThreePlan, []coldbrew.RenderSystem{})

	if err := client.Start(); err != nil {
		log.Fatal(err)
	}
}
