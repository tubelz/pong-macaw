package psystem

import (
	"github.com/tubelz/macaw/entity"
)

// AiSystem is the struct responsible to add the observer and handler to the collision event
type AiSystem struct {
	EntityManager *entity.Manager
	Name          string
}

// Init initialize this system
func (a *AiSystem) Init() {}

// Update will make the computer Y vel follow the same Y vel of the ball
func (a *AiSystem) Update() {
	computer := a.EntityManager.Get(1)
	ball := a.EntityManager.Get(2)

	component, _ := ball.GetComponent("physics")
	ballPhysics := component.(*entity.PhysicsComponent)

	component, _ = computer.GetComponent("physics")
	computerPhysics := component.(*entity.PhysicsComponent)

	if (ballPhysics.Vel.Y > 0 && computerPhysics.Vel.Y < 0) ||
		(ballPhysics.Vel.Y < 0 && computerPhysics.Vel.Y > 0) {
		computerPhysics.Vel.Y *= -1
	}
}
