package main

import (
	"fmt"

	blueprint_client "github.com/TheBitDrifter/blueprint/client"
	"github.com/TheBitDrifter/warehouse"
)

// TODO:
// Many:
//
//   - Wrap sto.NewEntity() in blueprint.NewEntities() and wrap storage for local blueprint version
func createControllerEntity(sto warehouse.Storage) error {
	// controllers := sto.NewEntities()
	return nil
}

func createCameraEntities(count, screenResX, screenResY int, sto warehouse.Storage) error {
	if count < 1 || count > 8 {
		return fmt.Errorf("camera count must be between 1 and 8")
	}

	// Calculate grid dimensions
	rows, cols := calculateGridDimensions(count)

	// Calculate camera dimensions
	camWidth := screenResX / cols
	camHeight := screenResY / rows

	entities, err := sto.NewEntities(count, blueprint_client.Components.Camera)
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		row := i / cols
		col := i % cols

		en := entities[i]
		cam := CameraComponent.GetFromEntity(en)
		cam.Positions.Screen.X = float64(col * camWidth)
		cam.Positions.Screen.Y = float64(row * camHeight)
		cam.Height = camHeight
		cam.Width = camWidth
	}

	return nil
}

func calculateGridDimensions(count int) (rows, cols int) {
	switch count {
	case 1:
		return 1, 1
	case 2:
		return 1, 2
	case 3, 4:
		return 2, 2
	case 5, 6:
		return 2, 3
	case 7, 8:
		return 2, 4
	default:
		return 1, 1
	}
}
