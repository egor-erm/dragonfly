package world

import (
	"git.jetbrains.space/dragonfly/dragonfly.git/dragonfly/entity/state"
	"github.com/go-gl/mathgl/mgl32"
	"io"
)

// Entity represents an entity in the world, typically an object that may be moved around and can be
// interacted with by other entities.
// Viewers of a world may view an entity when near it.
type Entity interface {
	io.Closer
	// Position returns the current position of the entity in the world.
	Position() mgl32.Vec3
	// World returns the current world of the entity. This is always the world that the entity can actually be
	// found in.
	World() *World
	// Yaw returns the yaw of the entity. This is horizontal rotation (rotation around the vertical axis), and
	// is 0 when the entity faces forward.
	Yaw() float32
	// Pitch returns the pitch of the entity. This is vertical rotation (rotation around the horizontal axis),
	// and is 0 when the entity faces forward.
	Pitch() float32
	// State returns a list of entity states which the entity is currently subject to. Generally, these states
	// alter the way the entity looks.
	State() []state.State
}
