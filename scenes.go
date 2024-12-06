package main

import (
	blueprint_client "github.com/TheBitDrifter/blueprint/client"
	"github.com/TheBitDrifter/warehouse"
)

func sceneOnePlan(sto warehouse.Storage) error {
	// Accepts resolution...should not be hardcoded like this.
	err := createCameraEntities(2, 640, 360, sto)
	if err != nil {
		return err
	}
	backgrounds, err := blueprint_client.BackgroundBuilder.AddParallaxBackgrounds(
		sto,
		"backgrounds/city/0",
		"backgrounds/city/1",
		"backgrounds/city/2",
		"backgrounds/city/3",
	)
	if err != nil {
		return err
	}
	err = backgrounds.Set(
		blueprint_client.ParallaxBackground{
			SpeedX: 0,
			SpeedY: 0,
		},
		blueprint_client.ParallaxBackground{
			SpeedX: 0.2,
			SpeedY: 0.2,
		},
		blueprint_client.ParallaxBackground{
			SpeedX: 0.4,
			SpeedY: 0.4,
		},
		blueprint_client.ParallaxBackground{
			SpeedX: 0.7,
			SpeedY: 0.7,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func sceneTwoPlan(sto warehouse.Storage) error {
	err := createCameraEntities(1, 640, 360, sto)
	if err != nil {
		return err
	}

	return nil
}

func sceneThreePlan(sto warehouse.Storage) error {
	return nil
}
