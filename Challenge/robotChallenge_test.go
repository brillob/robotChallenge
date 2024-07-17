package challenge_test

import (
	challenge "challenge/robotwarehouse/Challenge"
	"testing"
	"time"
)

func TestRobotMovement(t *testing.T) {
	warehouseInstance := challenge.WarehouseBuilder()
	robots := warehouseInstance.Robots()
	robot := robots[0]
	// Test moving north
	robot.EnqueueTask("N")
	state := robot.CurrentState()
	if state.X != 0 || state.Y != 1 {
		t.Errorf("Expected robot to be at (0, 1), but got (%d, %d)", state.X, state.Y)
	}

	// Test moving east
	robot.EnqueueTask("E")
	state = robot.CurrentState()
	if state.X != 1 || state.Y != 1 {
		t.Errorf("Expected robot to be at (1, 1), but got (%d, %d)", state.X, state.Y)
	}

	// Test moving south
	robot.EnqueueTask("S")
	state = robot.CurrentState()
	if state.X != 1 || state.Y != 0 {
		t.Errorf("Expected robot to be at (1, 0), but got (%d, %d)", state.X, state.Y)
	}

	// Test moving west
	robot.EnqueueTask("W")
	state = robot.CurrentState()
	if state.X != 0 || state.Y != 0 {
		t.Errorf("Expected robot to be at (0, 0), but got (%d, %d)", state.X, state.Y)
	}
}

func TestRobotTaskQueue(t *testing.T) {
	warehouseInstance := challenge.WarehouseBuilder()
	robots := warehouseInstance.Robots()
	robot := robots[0]

	// Enqueue multiple tasks
	robot.EnqueueTask("N E")
	robot.EnqueueTask("S W")

	// Wait for tasks to complete
	time.Sleep(3 * time.Second)

	state := robot.CurrentState()
	if state.X != 0 || state.Y != 0 {
		t.Errorf("Expected robot to be at (0, 0), but got (%d, %d)", state.X, state.Y)
	}
}

func TestRobotAddCrate(t *testing.T) {
	warehouseInstance := challenge.WarehouseBuilder()
	warehouseInstance.AddCrate(1, 2)
	if !challenge.HasCrate(warehouseInstance, 1, 2) {
		t.Errorf("Expected a crate at that location")
	}
}

func TestRobotDeleteCrate(t *testing.T) {
	warehouseInstance := challenge.WarehouseBuilder()
	//warehouseInstance.AddCrate(1, 2)
	warehouseInstance.DelCrate(1, 2)
	if challenge.HasCrate(warehouseInstance, 1, 2) {
		t.Errorf("Expected empty crate at that location")
	}
}
