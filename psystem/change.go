package psystem

import (
	"log"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/system"
)

// ChangeSystem is the struct responsible to add the observer and handler to the collision event
type ChangeSystem struct {
	Entities []entity.Entitier
	Name string
	system.Subject
}

// Assign assign entities with this system
func (c *ChangeSystem) Assign(entities []entity.Entitier) {
	c.Entities = entities
}

// Update is here just to comply with the system interface
func (c *ChangeSystem) Update() {
}

// Init adds the collision handler
func (c *ChangeSystem) Init(col *system.CollisionSystem, phs *system.PhysicsSystem) {
	log.Printf("Init %s", c.Name)
	col.AddHandler("collision event", system.InvertVel)
	phs.AddHandler("border event", invertYAxis)
}

// invertYAxis invert the vel on the Y axis of the collided object.
func invertYAxis(event system.Event) {
	log.Printf("come on")
	border := event.(*system.BorderEvent)
	
	component, _ := border.Ent.GetComponent("position")
	position := component.(*entity.PositionComponent)
	
	component, _ = border.Ent.GetComponent("physics")
	physics := component.(*entity.PhysicsComponent)
  
	switch border.Side {
		case "top":
			position.Pos.Y = 1
			physics.FuturePos.Y = 1
			physics.Vel.Y *= -1
		case "bottom":
			position.Pos.Y = 599
			physics.FuturePos.Y = 599
			physics.Vel.Y *= -1
	}
}

