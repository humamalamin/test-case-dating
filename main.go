package main

import (
	"fmt"
	"os"
	"time"

	"github.com/humamalamin/test-case-dating/pkg/manager"
	"github.com/humamalamin/test-case-dating/pkg/server"

	authRoutes "github.com/humamalamin/test-case-dating/api/handlers/auth"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	mgr, err := manager.NewInit()
	if err != nil {
		return err
	}

	tzLocation, err := time.LoadLocation(mgr.GetConfig().AppTz)
	if err != nil {
		return err
	}
	time.Local = tzLocation

	server := server.NewServer(mgr.GetConfig())

	server.Router.Use(mgr.GetMiddleware().InitLog)

	server.RegisterRouter(server.Router)
	authRoutes.NewRoutes(server.Router, mgr)

	return server.ListenAndServe()
}
