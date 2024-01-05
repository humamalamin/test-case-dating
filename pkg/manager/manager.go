package manager

import (
	paginationHelper "github.com/humamalamin/test-case-dating/helpers/pagination"
	jwtAuth "github.com/humamalamin/test-case-dating/pkg/auth/jwt"
	"github.com/humamalamin/test-case-dating/pkg/auth/middleware"
	"github.com/humamalamin/test-case-dating/pkg/config"
	"github.com/humamalamin/test-case-dating/pkg/database"
	"github.com/humamalamin/test-case-dating/pkg/server"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Manager interface {
	GetConfig() *config.Config
	GetDatabase() *gorm.DB
	GetServer() *server.Server
	GetJwt() jwtAuth.Jwt
	GetMiddleware() middleware.Middleware
	GetPaginate() paginationHelper.Pagination
}

type manager struct {
	config         *config.Config
	database       *gorm.DB
	server         *server.Server
	jwtAuth        jwtAuth.Jwt
	middlewareAuth middleware.Middleware
	pagination     paginationHelper.Pagination
}

// GetPaginate implements Manager.
func (m *manager) GetPaginate() paginationHelper.Pagination {
	return m.pagination
}

// GetJwt implements Manager.
func (m *manager) GetJwt() jwtAuth.Jwt {
	return m.jwtAuth
}

// GetMiddleware implements Manager.
func (m *manager) GetMiddleware() middleware.Middleware {
	return m.middlewareAuth
}

// GetConfig implements Manager.
func (m *manager) GetConfig() *config.Config {
	return m.config
}

// GetDatabase implements Manager.
func (m *manager) GetDatabase() *gorm.DB {
	return m.database
}

// GetServer implements Manager.
func (m *manager) GetServer() *server.Server {
	return m.server
}

func NewInit() (Manager, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-1] Failed to Initialize Configuration")
		return nil, err
	}

	srv := server.NewServer(cfg)
	db, err := database.NewGorm(cfg).Connect()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-2] Failed to Initialize Database" + cfg.DatabaseName)
		return nil, err
	}

	jwt := jwtAuth.NewJwt(cfg)
	middleware := middleware.NewMiddleware(cfg)

	paginationHelper := paginationHelper.NewPagination()
	return &manager{
		config:         cfg,
		database:       db,
		server:         srv,
		jwtAuth:        jwt,
		middlewareAuth: middleware,
		pagination:     paginationHelper,
	}, nil
}
