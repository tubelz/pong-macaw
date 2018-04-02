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
	EntityManager *entity.Manager
	Name string
	system.Subject
	CollisionSystem *system.CollisionSystem
}

// Update is here just to comply with the system interface
func (s *ScoreSystem) Update() {
}

// Init adds the collision handler
func (s *ScoreSystem) Init() {
	log.Printf("Init %s", s.Name)
	s.CollisionSystem.AddHandler("border event", s.checkScore)
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

	if border.Side == "right" {
		log.Printf("entity: %d", obj.GetID())
		log.Println("you scored")
		s.EntityManager.Delete(obj.GetID())
		myScore := s.Entities[3].(*entity.Entity)
		f := myScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score + 1)
		f.Modified = true
	}
	if border.Side == "left" {
		log.Printf("entity: %d", obj.GetID())
		log.Println(position.Pos.X)
		log.Println("he scored")

		hisScore := s.Entities[4].(*entity.Entity)
		f := hisScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score + 1)
		f.Modified = true
	}
}
