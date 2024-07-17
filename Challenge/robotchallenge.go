package challenge

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Warehouse interface {
	Robots() []Robot
}

type CrateWarehouse interface {
	Warehouse
	AddCrate(x uint, y uint) error
	DelCrate(x uint, y uint) error
}

type RobotState struct {
	X        uint
	Y        uint
	HasCrate bool
}
type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)
	CancelTask(taskID string) error
	CurrentState() RobotState
}

type warehouse struct {
	robots []Robot
	//can use a object as key as well instead of multi-dimensional key
	crates map[uint]map[uint]bool
	//mutex is used for locking so only one at a time is enforced
	mu sync.Mutex
}

func WarehouseBuilder() *warehouse {
	warehouseInstance := new(warehouse)
	warehouseInstance.robots = []Robot{
		&robot{
			state: RobotState{
				X:        0,
				Y:        0,
				HasCrate: false,
			},
			tasks: make(map[string]string),
		}}
	crates := make(map[uint]map[uint]bool)
	warehouseInstance.crates = crates
	return warehouseInstance
}
func HasCrate(w *warehouse, x uint, y uint) bool {
	println(w.crates[x][y])
	return w.crates[x][y] == true
}

// Implementation of Warehouse Interface
func (w *warehouse) Robots() []Robot {
	return w.robots
}

// Implementation of CrateWarehouse Interface
func (w *warehouse) AddCrate(x uint, y uint) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.crates[x][y] {
		return errors.New("crate already exists at this location")
	}

	w.crates[x] = map[uint]bool{
		y: true}
	return nil
}

func (w *warehouse) DelCrate(x uint, y uint) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.crates[x][y] {
		return errors.New("no crate at this location")
	}

	delete(w.crates[x], y)

	return nil
}

type robot struct {
	state RobotState
	tasks map[string]string
	mu    sync.Mutex
}

func (robotvar *robot) EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error) {
	robotvar.mu.Lock()
	defer robotvar.mu.Unlock()

	taskID = fmt.Sprintf("%d", len(robotvar.tasks)+1)
	robotvar.tasks[taskID] = commands
	//fmt.Println(commands)
	idPost := strings.Split(commands, ``)
	for _, val := range idPost {
		switch val {
		case "N":
			robotvar.state.Y++
		case "E":
			robotvar.state.X++
		case "S":
			robotvar.state.Y--
		case "W":
			robotvar.state.X--
		case "G":
			robotvar.state.HasCrate = true
		case "D":
			robotvar.state.HasCrate = false
		}
		time.Sleep(1 * time.Second)
	}

	return taskID, nil, nil

}

func (r *robot) CancelTask(taskID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}

	//close(taskChan)
	delete(r.tasks, taskID)

	return nil
}

func (r *robot) CurrentState() RobotState {
	return r.state
}

// func main() {
// 	// Create a new warehouse
// 	warehouseInstance := new(warehouse)
// 	warehouseInstance.robots = []Robot{
// 		&robot{
// 			state: RobotState{
// 				X:        0,
// 				Y:        0,
// 				HasCrate: false,
// 			},
// 			tasks: make(map[string]string),
// 		}}
// 	warehouseInstance.crates = make(map[uint]map[uint]bool)
// 	robots := warehouseInstance.Robots()

// 	// Enqueue a task for the first robot
// 	taskID, _, err := robots[0].EnqueueTask("NNEEG")
// 	if err != nil {
// 		fmt.Println("Error enqueuing task:", err)
// 		return
// 	}
// 	fmt.Println("Task ID:", taskID)

// 	fmt.Println(robots[0].CurrentState().X)
// 	fmt.Println(robots[0].CurrentState().Y)
// 	// Cancel the task
// 	err1 := robots[0].CancelTask(taskID)
// 	if err1 != nil {
// 		fmt.Println("Error cancelling task:", err)
// 		return
// 	}
// }
