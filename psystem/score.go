package psystem

import (
	"fmt"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/system"
	"log"
	"strconv"
)

// ScoreSystem is the struct responsible to keep the score of the game
type ScoreSystem struct {
	Entities      []entity.Entitier
	EntityManager *entity.Manager
	Name          string
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

	components := obj.GetComponents()
	_, ok = components["position"]
	if !ok {
		return
	}
	if border.Side == "right" {
		log.Printf("entity: %d", obj.GetID())
		log.Println("you scored")

		myScore := s.EntityManager.Get(3)
		f := myScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score+1)
		f.Modified = true
	}
	if border.Side == "left" {
		log.Printf("entity: %d", obj.GetID())
		log.Println("he scored")

		hisScore := s.EntityManager.Get(4)
		f := hisScore.GetComponents()["font"].(*entity.FontComponent)
		score, _ := strconv.Atoi(f.Text)
		f.Text = fmt.Sprintf("%d", score+1)
		f.Modified = true
	}
}
