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
	err = macaw.Initialize(false, true)
	if err != nil {
		fmt.Println("Macaw could not initialize")
	}

	mfont := entity.MFont{File: "assets/manaspc.ttf", Size: uint8(22)}
	font := mfont.Open()
	defer mfont.Close()

	input := &input.Manager{}
	systems := initializeSystems(input)
	entities := initializeEntities(systems, font)

	// lazy way to assign entities to the system. usually you'd want to assign the
	// right entities only, but this works for our problem here.
	for _, system := range systems {
		system.Assign(entities)
	}

	gameLoop := initializeGameLoop(systems, input)

	gameLoop.Run()
}

func initializeSystems(im *input.Manager) []system.Systemer {
	render := &system.RenderSystem{Name: "render system"}
	player := &psystem.PlayerSystem{Name: "player system", InputManager: im}
	physics := &system.PhysicsSystem{Name: "physics system"}
	collision := &system.CollisionSystem{Name: "collision system"}
	score := &psystem.ScoreSystem{Name: "score system"}
	change := &psystem.ChangeSystem{Name: "change system"}
	ai := &psystem.AiSystem{Name: "ai system"}
	// initialize some systems that require such act
	render.Init(macaw.Window)
	change.Init(collision, physics)
	score.Init(physics)

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
func initializeEntities(systems []system.Systemer, font *ttf.Font) []entity.Entitier {
	player := &entity.Entity{}
	computer := &entity.Entity{}
	ball := &entity.Entity{}
	playerScore := &entity.Entity{}
	computerScore := &entity.Entity{}
	entities := []entity.Entitier{player, computer, ball, playerScore, computerScore}

	for _, e := range entities {
		e.Init()
	}

	//load sprite
	render := systems[0].(*system.RenderSystem)
	// collision := systems[1].(*system.CollisionSystem)

	acc := &math.FPoint{0, 0}
	vel := &math.FPoint{0, 1}

	// player
	player.AddComponent("position", &entity.PositionComponent{&sdl.Point{20, 20}})
	player.AddComponent("collision", &entity.CollisionComponent{Size: &sdl.Point{10, 80}})
	player.AddComponent("geometry", &entity.RectangleComponent{
		Size:   &sdl.Point{10, 80},
		Color:  &sdl.Color{0x66, 0x66, 0x66, 0xFF},
		Filled: true,
	})

	// computer
	computer.AddComponent("position", &entity.PositionComponent{&sdl.Point{770, 20}})
	computer.AddComponent("physics", &entity.PhysicsComponent{Vel: vel, Acc: acc, FuturePos: &math.FPoint{770, 20}})
	computer.AddComponent("collision", &entity.CollisionComponent{Size: &sdl.Point{10, 80}})
	computer.AddComponent("geometry", &entity.RectangleComponent{
		Size:   &sdl.Point{10, 80},
		Color:  &sdl.Color{0xFF, 0x66, 0x66, 0xFF},
		Filled: true,
	})

	// ball
	ball.AddComponent("position", &entity.PositionComponent{&sdl.Point{300, 20}})
	ball.AddComponent("physics", &entity.PhysicsComponent{Vel: &math.FPoint{8, 1}, Acc: acc, FuturePos: &math.FPoint{300, 20}})
	ball.AddComponent("collision", &entity.CollisionComponent{Size: &sdl.Point{10, 10}})
	ball.AddComponent("geometry", &entity.RectangleComponent{
		Size:   &sdl.Point{10, 10},
		Color:  &sdl.Color{0x00, 0x00, 0x00, 0xFF},
		Filled: false,
	})

	// player score
	playerScore.AddComponent("position", &entity.PositionComponent{&sdl.Point{200, 20}})
	playerScore.AddComponent("font", &entity.FontComponent{Text: "0", Modified: true, Font: font})
	playerScore.AddComponent("render", &entity.RenderComponent{Renderer: render.Renderer})

	// computer score
	computerScore.AddComponent("position", &entity.PositionComponent{&sdl.Point{500, 20}})
	computerScore.AddComponent("font", &entity.FontComponent{Text: "0", Modified: true, Font: font})
	computerScore.AddComponent("render", &entity.RenderComponent{Renderer: render.Renderer})

	return entities
}

func initializeGameLoop(systems []system.Systemer, im *input.Manager) *macaw.GameLoop {
	gameLoop := &macaw.GameLoop{InputManager: im}
	gameLoop.AddRenderSystem(systems[0].(*system.RenderSystem))
	for _, system := range systems[1:] {
		gameLoop.AddGameUpdateSystem(system)
	}
	return gameLoop
}
