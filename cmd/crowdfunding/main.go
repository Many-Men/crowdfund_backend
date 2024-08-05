package main

import (
	"github.com/Many-Men/crowdfund_backend/config"
	"github.com/Many-Men/crowdfund_backend/internal/app/db"
	"github.com/Many-Men/crowdfund_backend/internal/app/server"
)

func main() {
	cfg := config.Load()
	dbs := db.ConnectToDB(cfg)

	server.RunHTTPServer(cfg, dbs)
}
