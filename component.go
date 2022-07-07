package main

type ComponentType string
type ComponentTyper interface {
	Type() ComponentType
}
