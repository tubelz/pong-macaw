package main

import (
	"fmt"

	"github.com/tubelz/macaw"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/math"
	"github.com/tubelz/macaw/system"
	"github.com/tubelz/pong-macaw/psystem"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	fmt.Println("Pong")
	var err error
	err = macaw.Initialize()
	if err != nil {
		fmt.Println("Macaw could not initialize")
	}
	defer macaw.Quit()

	mfont := entity.MFont{File: "assets/manaspc.ttf", Size: uint8(22)}
	font := mfont.Open()
	defer mfont.Close()

	input := &input.Manager{}
	entityManager := &entity.Manager{}
	systems := initializeSystems(input, entityManager)
	initializeEntities(entityManager, systems, font)

	gameLoop := initializeGameLoop(systems, input)

	gameLoop.Run()
}

func initializeSystems(im *input.Manager, em *entity.Manager) []system.Systemer {
	render := &system.RenderSystem{Name: "render system", Window: macaw.Window, EntityManager: em}
	player := &psystem.PlayerSystem{Name: "player system", InputManager: im, EntityManager: em}
	physics := &system.PhysicsSystem{Name: "physics system", EntityManager: em}
	collision := &system.CollisionSystem{Name: "collision system", EntityManager: em}
	score := &psystem.ScoreSystem{Name: "score system", CollisionSystem: collision, EntityManager: em}
	change := &psystem.ChangeSystem{Name: "change system", CollisionSystem: collision, EntityManager: em}
	ai := &psystem.AiSystem{Name: "ai system", EntityManager: em}
	// initialize some systems that require such actions
	render.Init()

	systems := []system.Systemer{
		render,
		player,
		physics,
		collision,
		score,
		change,
		ai,
	}
	return systems
}

// func initializeEntities(systems []system.Systemer, font *ttf.Font) ([]entity.Entitier){
func initializeEntities(em *entity.Manager, systems []system.Systemer, font *ttf.Font) {
	player := em.Create("player")
	computer := em.Create("computer")
	ball := em.Create("ball")
	playerScore := em.Create("pscore")
	computerScore := em.Create("cscore")
	camera := em.Create("camera")

	//load sprite
	render := systems[0].(*system.RenderSystem)

	acc := &math.FPoint{0, 0}
	vel := &math.FPoint{0, 1}

	// player
	player.AddComponent(&entity.PositionComponent{&sdl.Point{20, 20}})
	player.AddComponent(&entity.CollisionComponent{CollisionAreas: []sdl.Rect{sdl.Rect{0, 0, 10, 80}}})
	player.AddComponent(&entity.RenderComponent{RenderType: entity.RTGeometry})
	player.AddComponent(&entity.RectangleComponent{
		Size:   &sdl.Point{10, 80},
		Color:  &sdl.Color{0x66, 0x66, 0x66, 0xFF},
		Filled: true,
	})

	// computer
	computer.AddComponent(&entity.PositionComponent{&sdl.Point{770, 20}})
	computer.AddComponent(&entity.PhysicsComponent{Vel: vel, Acc: acc, FuturePos: &math.FPoint{770, 20}})
	computer.AddComponent(&entity.CollisionComponent{CollisionAreas: []sdl.Rect{sdl.Rect{0, 0, 10, 80}}})
	computer.AddComponent(&entity.RenderComponent{RenderType: entity.RTGeometry})
	computer.AddComponent(&entity.RectangleComponent{
		Size:   &sdl.Point{10, 80},
		Color:  &sdl.Color{0xFF, 0x66, 0x66, 0xFF},
		Filled: true,
	})

	// ball
	ball.AddComponent(&entity.PositionComponent{&sdl.Point{300, 20}})
	ball.AddComponent(&entity.PhysicsComponent{Vel: &math.FPoint{8, 1}, Acc: acc, FuturePos: &math.FPoint{300, 20}})
	ball.AddComponent(&entity.CollisionComponent{CollisionAreas: []sdl.Rect{sdl.Rect{0, 0, 10, 10}}})
	ball.AddComponent(&entity.RenderComponent{RenderType: entity.RTGeometry})
	ball.AddComponent(&entity.RectangleComponent{
		Size:   &sdl.Point{10, 10},
		Color:  &sdl.Color{0x00, 0x00, 0x00, 0xFF},
		Filled: false,
	})

	// player score
	playerScore.AddComponent(&entity.PositionComponent{&sdl.Point{200, 20}})
	playerScore.AddComponent(&entity.FontComponent{Text: "0", Modified: true, Font: font})
	playerScore.AddComponent(&entity.RenderComponent{RenderType: entity.RTFont})

	// computer score
	computerScore.AddComponent(&entity.PositionComponent{&sdl.Point{500, 20}})
	computerScore.AddComponent(&entity.FontComponent{Text: "0", Modified: true, Font: font})
	computerScore.AddComponent(&entity.RenderComponent{RenderType: entity.RTFont})

	// camera
	camera.AddComponent(&entity.PositionComponent{&sdl.Point{0, 0}})
	camera.AddComponent(&entity.CameraComponent{
		ViewportSize: sdl.Point{800, 600},
		WorldSize:    sdl.Point{800, 600},
	})
	render.SetCamera(camera)

}

func initializeGameLoop(systems []system.Systemer, im *input.Manager) *macaw.GameLoop {
	gameLoop := &macaw.GameLoop{InputManager: im}
	// game loop
	sceneGame := &macaw.Scene{Name: "game"}
	sceneGame.AddRenderSystem(systems[0].(*system.RenderSystem))
	for _, system := range systems[1:] {
		sceneGame.AddGameUpdateSystem(system)
	}

	gameLoop.AddScene(sceneGame)

	return gameLoop
}
