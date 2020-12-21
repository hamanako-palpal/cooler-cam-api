//+build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/handle"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/infra"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/services"
)

// Initialize 初期化
func Initialize(db *sql.DB) (*handle.HandlerGen, error) {

	wire.Build(
		infra.InitVisionCli,
		infra.InitLabelCli,
		services.NewCamService,
		handle.InitSmplHandler,
		handle.InitCamHandler,
		handle.InitHandlerGen,
	)

	return &handle.HandlerGen{}, nil
}
