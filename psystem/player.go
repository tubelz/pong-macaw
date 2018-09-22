package psystem

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/veandco/go-sdl2/sdl"
)

// PlayerSystem is the struct that contains the controllable stick
type PlayerSystem struct {
	EntityManager *entity.Manager
	Name          string
	InputManager  *input.Manager
}

// Init initialize this system
func (p *PlayerSystem) Init() {}

// Update handle the input event
func (p *PlayerSystem) Update() {
	empty := sdl.KeyboardEvent{}
	if button := p.InputManager.Button(); button != empty {
		if button.Keysym.Sym == sdl.K_UP {
			e := p.EntityManager.Get(0)
			p := e.GetComponent(&entity.PositionComponent{})
			pos := p.(*entity.PositionComponent)
			if pos.Pos.Y > 0 {
				pos.Pos.Y -= 10
			}
		}
		if button.Keysym.Sym == sdl.K_DOWN {
			e := p.EntityManager.Get(0)
			p := e.GetComponent(&entity.PositionComponent{})
			pos := p.(*entity.PositionComponent)
			if pos.Pos.Y < 520 {
				pos.Pos.Y += 10
			}
		}
		if button.Keysym.Sym == sdl.K_a && button.State == 0 {
			e := p.EntityManager.Get(2)
			p := e.GetComponent(&entity.PhysicsComponent{})
			phs := p.(*entity.PhysicsComponent)
			phs.Vel.X *= -1
			phs.Vel.Y *= -1
		}
	}
}
