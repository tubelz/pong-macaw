package psystem

import (
	"log"
	"fmt"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/system"
	"strconv"
)

// ScoreSystem is the struct responsible to keep the score of the game
type ScoreSystem struct {
	Entities []entity.Entitier
	Name string
	system.Subject
}

// Assign assign entities with this system
func (s *ScoreSystem) Assign(entities []entity.Entitier) {
	s.Entities = entities
}

// Update is here just to comply with the system interface
func (s *ScoreSystem) Update() {	
}

// Init adds the collision handler
func (s *ScoreSystem) Init(sys *system.PhysicsSystem) {
	log.Printf("Init %s", s.Name)
	sys.AddHandler("border event", s.checkScore)
}

func (s *ScoreSystem) checkScore(event system.Event) {
	border := event.(*system.BorderEvent)
	obj := border.Ent
	var ok bool
	var component entity.Component

	components := obj.GetComponents()
	component, ok = components["position"]
	if !ok {
			return
	}
	position := component.(*entity.PositionComponent)
	
	component, ok = components["physics"]
	if !ok {
			return
	}
	log.Printf("POS: %v", position.Pos)
	physics := component.(*entity.PhysicsComponent)
	if position.Pos.X >= 800 {
		log.Printf("entity: %d", obj.GetID())
		log.Println("you scored")
		
		hisScore := s.Entities[4].(*entity.Entity)
		f := hisScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score + 1)
		f.Modified = true
		
		physics.Vel.X *= -1
		physics.FuturePos.X = 799 
		position.Pos.X = 799
	}
	if position.Pos.X < 0 {
		log.Printf("entity: %d", obj.GetID())
		log.Println(position.Pos.X)
		log.Println("he scored")
		
		myScore := s.Entities[3].(*entity.Entity)
		f := myScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score + 1)
		f.Modified = true
		
		physics.Vel.X *= -1
		physics.FuturePos.X = 1
		position.Pos.X = 1
	}
}