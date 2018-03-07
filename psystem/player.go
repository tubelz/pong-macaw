package psystem

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/veandco/go-sdl2/sdl"
)

// PlayerSystem is the struct that contains the controllable stick
type PlayerSystem struct {
	Entities []entity.Entitier
	Name string
	InputManager *input.Manager
}

// Assign assign entities with this system
func (p *PlayerSystem) Assign(entities []entity.Entitier) {
	p.Entities = entities
}

// Update handle the input event
func (p *PlayerSystem) Update() {
	if button := p.InputManager.Button(); button != nil {
		if button.Keysym.Sym == sdl.K_UP {
			e := p.Entities[0].(*entity.Entity)
			p, _ := e.GetComponent("position")
			pos := p.(*entity.PositionComponent)
			if pos.Pos.Y > 0 {
				pos.Pos.Y -= 10
			}
		}
		if button.Keysym.Sym == sdl.K_DOWN {
			e := p.Entities[0].(*entity.Entity)
			p, _ := e.GetComponent("position")
			pos := p.(*entity.PositionComponent)
			if pos.Pos.Y < 520 {
				pos.Pos.Y += 10
			}
		}
		if button.Keysym.Sym == sdl.K_a && button.State == 0 {
			e := p.Entities[2].(*entity.Entity)
			p, _ := e.GetComponent("physics")
			phs := p.(*entity.PhysicsComponent)
			phs.Vel.X *= -1
			phs.Vel.Y *= -1
		}
	}
}
