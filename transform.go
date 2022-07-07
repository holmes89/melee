package main

const TransformType ComponentType = "TRANSFORM"

// TransformComponent describes the current position of an entity
type TransformComponent struct {
	PosX, PosY float64
}

func (t *TransformComponent) Type() ComponentType { return TransformType }
