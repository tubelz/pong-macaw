package psystem

import (
	"log"
	"github.com/tubelz/macaw/entity"
)

// AiSystem is the struct responsible to add the observer and handler to the collision event
type AiSystem struct {
	Entities []entity.Entitier
	Name string
}

// Assign assign entities with this system
func (a *AiSystem) Assign(entities []entity.Entitier) {
	a.Entities = entities
}

// Update will make the computer Y vel follow the same Y vel of the ball
func (a *AiSystem) Update() {
	computer := a.Entities[1].(*entity.Entity)
	ball := a.Entities[2].(*entity.Entity)

	component, _ := ball.GetComponent("physics")
	ballPhysics := component.(*entity.PhysicsComponent)

	component, _ = computer.GetComponent("physics")
	computerPhysics := component.(*entity.PhysicsComponent)

	if (ballPhysics.Vel.Y > 0 && computerPhysics.Vel.Y < 0) ||
	 	(ballPhysics.Vel.Y < 0 && computerPhysics.Vel.Y > 0) {
		computerPhysics.Vel.Y *= -1
	}
}
