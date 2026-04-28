//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func Build() (Application, func(), error) {
	panic(wire.Build(ProviderSet))
}
