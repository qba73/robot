// package robotname contains solution for Robot Name exercise on Exercism.
package robotname

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
)

// limit defines the number of names possible.
const limit = 26 * 26 * 10 * 10 * 10

// memory stores all then names current in use by robots.
//var memory = make(map[string]bool)

type memory struct {
	mu    sync.Mutex
	items map[string]bool
}

func (m *memory) len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.items)
}

func (m *memory) get(item string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.items[item]
	return ok
}

func (m *memory) set(item string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[item] = true
}

var mem = &memory{items: make(map[string]bool)}

// Robot defines a robot, it has only one field which is its name.
type Robot struct {
	name string
}

// Name method returns the name of the robot, if it doesn't have a name then it would
// generate a random unique name.
func (r *Robot) Name() (string, error) {

	if r.name != "" {
		return r.name, nil
	}

	//if limit == len(memory) {
	if limit == mem.len() {
		return "", errors.New("maximum limit of unique names that could be created reached")
	}
	r.name = r.nameGen()

	return r.name, nil
}

// Reset method reset's the robot to its default settings. i.e. it replaces the robot's current
// name with a random unique name.
func (r *Robot) Reset() error {
	// cause of the restrictions of this exercise, this would cause tests to fail. even tho its
	// sensible.
	// delete(memory, r.name)
	r.name = ""
	_, err := r.Name()
	if err != nil {
		return err
	}
	return nil
}

// nameGen generates a random unique name in a specific format and returns it.
func (r *Robot) nameGen() string {

	var output strings.Builder
	for i := 0; i < 2; i++ {
		output.WriteByte(byte('A') + byte(rand.Intn(26)))
	}
	for i := 0; i < 3; i++ {
		output.WriteByte(byte('0') + byte(rand.Intn(10)))
	}
	name := output.String()
	//if memory[name] {
	if mem.get(name) {
		return r.nameGen()
	}
	//memory[name] = true
	mem.set(name)
	return name
}
