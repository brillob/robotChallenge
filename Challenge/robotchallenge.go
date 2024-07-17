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

func WarehouseBuilder(robotType string) *warehouse {
	warehouseInstance := new(warehouse)
	BuildRobot(warehouseInstance, robotType)
	crates := make(map[uint]map[uint]bool)
	warehouseInstance.crates = crates
	return warehouseInstance
}

func BuildRobot(warehouseInstance *warehouse, robotype string) {
	robotState := RobotState{
		X:        0,
		Y:        0,
		HasCrate: false,
	}
	robotTask := make(map[string]string)
	if robotype == "D" {
		warehouseInstance.robots = []Robot{
			&diagonalrobot{
				state: robotState,
				tasks: robotTask,
			}}
	} else {
		warehouseInstance.robots = []Robot{
			&robot{
				state: robotState,
				tasks: robotTask,
			}}
	}

}

//	func WarehouseBuilderDiagonalRobot() *warehouse {
//		warehouseInstance := new(warehouse)
//		warehouseInstance.robots = []Robot{
//			&diagonalrobot{
//				state: RobotState{
//					X:        0,
//					Y:        0,
//					HasCrate: false,
//				},
//				tasks: make(map[string]string),
//			}}
//		crates := make(map[uint]map[uint]bool)
//		warehouseInstance.crates = crates
//		return warehouseInstance
//	}
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
	indivualCommand := strings.Split(commands, ``)
	for _, val := range indivualCommand {
		switch val {
		case "N":
			robotvar.state.Y++
			sleep(1)
		case "E":
			robotvar.state.X++
			sleep(1)
		case "S":
			robotvar.state.Y--
			sleep(1)
		case "W":
			robotvar.state.X--
			sleep(1)
		case "G":
			robotvar.state.HasCrate = true
			sleep(1)
		case "D":
			robotvar.state.HasCrate = false
			sleep(1)
		}

	}

	return taskID, nil, nil

}
func (robotvar *diagonalrobot) EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error) {
	robotvar.mu.Lock()
	defer robotvar.mu.Unlock()

	taskID = fmt.Sprintf("%d", len(robotvar.tasks)+1)
	robotvar.tasks[taskID] = commands
	//fmt.Println(commands)
	individualcommands := strings.Split(commands, ``)
	length := len(individualcommands)
	for i := 0; i < length; i++ {

		switch individualcommands[i] {
		case "N":
			robotvar.state.Y++
			if moveHorizontal(&i, individualcommands, &robotvar.state) == true {
				length--
			}
			sleep(1)
		case "E":
			robotvar.state.X++
			if moveVertical(&i, individualcommands, &robotvar.state) == true {
				length--
			}
			sleep(1)
		case "S":
			robotvar.state.Y--
			if moveHorizontal(&i, individualcommands, &robotvar.state) == true {
				length--
			}
			sleep(1)
		case "W":
			robotvar.state.X--
			if moveVertical(&i, individualcommands, &robotvar.state) == true {
				length--
			}
			sleep(1)
		case "G":
			robotvar.state.HasCrate = true
			sleep(1)
		case "D":
			robotvar.state.HasCrate = false
			sleep(1)
		}

	}
	return taskID, nil, nil

}
func moveHorizontal(i *int, individualcommands []string, state *RobotState) bool {
	m := *i
	isIncrement := false
	robotState := *state
	if m < len(individualcommands)-1 {
		if individualcommands[m+1] == "E" {
			robotState.X++
			isIncrement = true
		}
		if individualcommands[m+1] == "W" {
			robotState.X--
			isIncrement = true

		}
	}

	return isIncrement
}
func moveVertical(i *int, individualcommands []string, state *RobotState) bool {
	m := *i
	isIncrement := false
	robotState := *state
	if m < len(individualcommands)-1 {
		if individualcommands[m+1] == "N" {
			robotState.Y++
			isIncrement = true
		}
		if individualcommands[m+1] == "S" {
			robotState.Y--
			isIncrement = true
		}
	}

	return isIncrement
}

// can also do a baseclass and inheritance
type diagonalrobot struct {
	state RobotState
	tasks map[string]string
	mu    sync.Mutex
}

func sleep(numberOfSeconds time.Duration) {
	time.Sleep(numberOfSeconds * time.Second)
}
func (r *robot) CancelTask(taskID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, taskID)

	return nil
}

func (r *robot) CurrentState() RobotState {
	return r.state
}
func (r *diagonalrobot) CancelTask(taskID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, taskID)

	return nil
}

func (r *diagonalrobot) CurrentState() RobotState {
	return r.state
}
