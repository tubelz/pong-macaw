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
func (c *ChangeSystem) Init(col *system.CollisionSystem) {
	log.Printf("Init %s", c.Name)
	col.AddHandler("collision event", system.InvertVel)
	col.AddHandler("border event", invertYAxis)
}

// invertYAxis invert the vel on the Y axis of the collided object.
func invertYAxis(event system.Event) {
	border := event.(*system.BorderEvent)

	if border.Ent.GetID() == 0 {
		return
	}
	component, _ := border.Ent.GetComponent("position")
	position := component.(*entity.PositionComponent)

	component, _ = border.Ent.GetComponent("physics")
	physics := component.(*entity.PhysicsComponent)

	component, _ = border.Ent.GetComponent("collision")
	collision := component.(*entity.CollisionComponent)

	switch border.Side {
		case "top":
			position.Pos.Y = 1
			physics.FuturePos.Y = 1
			physics.Vel.Y *= -1
		case "bottom":
			size := collision.Size.Y
			position.Pos.Y = 599 - size
			physics.FuturePos.Y = float32(599 - size)
			physics.Vel.Y *= -1
		case "left":
			position.Pos.X = 1
			physics.FuturePos.X = 1
			physics.Vel.X *= -1
		case "right":
			size := collision.Size.X
			position.Pos.X = 799 - size
			physics.FuturePos.X = float32(799 - size)
			physics.Vel.X *= -1
	}
}
