package psystem

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/system"
	"log"
)

// ChangeSystem is the struct responsible to add the observer and handler to the collision event
type ChangeSystem struct {
	EntityManager *entity.Manager
	Name          string
	CollisionSystem *system.CollisionSystem
}

// Update is here just to comply with the system interface
func (c *ChangeSystem) Update() {
}

// Init adds the collision handler
func (c *ChangeSystem) Init() {
	log.Printf("Init %s", c.Name)
	c.CollisionSystem.AddHandler("collision event", increaseVel)
	c.CollisionSystem.AddHandler("collision event", system.InvertVel)
	c.CollisionSystem.AddHandler("border event", invertAxis)
}

func increaseVel(event system.Event) {
	collision := event.(*system.CollisionEvent)

	if collision.Ent.GetID() != 2 {
		return
	}
	component, _ := collision.Ent.GetComponent("physics")
	physics := component.(*entity.PhysicsComponent)

	log.Printf("%v", physics.Vel)
	if physics.Vel.X > 0 && physics.Vel.X < 13 {
		physics.Vel.X++
		physics.Vel.Y += .2
	} else if physics.Vel.X < 0 && physics.Vel.X > -13 {
		physics.Vel.X--
		physics.Vel.Y -= .2
	}
}

// invertAxis invert the vel according to the border it hit
func invertAxis(event system.Event) {
	border := event.(*system.BorderEvent)

	if border.Ent.GetID() == 0 {
		return
	}
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
		size := int32(10)
		position.Pos.Y = 599 - size
		physics.FuturePos.Y = float32(599 - size)
		physics.Vel.Y *= -1
	case "left":
		position.Pos.X = 1
		physics.FuturePos.X = 1
		physics.Vel.X *= -1
	case "right":
		size := int32(10)
		position.Pos.X = 799 - size
		physics.FuturePos.X = float32(799 - size)
		physics.Vel.X *= -1
	}
}
