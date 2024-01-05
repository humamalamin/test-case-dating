package manager

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/humamalamin/test-case-dating/helpers/pagination"
	paginationHelper "github.com/humamalamin/test-case-dating/helpers/pagination"
	jwtAuth "github.com/humamalamin/test-case-dating/pkg/auth/jwt"
	"github.com/humamalamin/test-case-dating/pkg/auth/middleware"
	"github.com/humamalamin/test-case-dating/pkg/config"
	"github.com/humamalamin/test-case-dating/pkg/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FakeManager interface {
	Manager
}

type fakeManager struct {
	config         *config.Config
	server         *server.Server
	dbGorm         *gorm.DB
	jwtAuth        jwtAuth.Jwt
	middlewareAuth middleware.Middleware
	pagination     paginationHelper.Pagination
}

// GetJwt implements FakeManager.
func (fm *fakeManager) GetJwt() jwtAuth.Jwt {
	return fm.GetJwt()
}

// GetMiddleware implements FakeManager.
func (fm *fakeManager) GetMiddleware() middleware.Middleware {
	return fm.GetMiddleware()
}

// GetPaginate implements FakeManager.
func (fm *fakeManager) GetPaginate() pagination.Pagination {
	return fm.GetPaginate()
}

// GetConfig implements FakeManager.
func (fm *fakeManager) GetConfig() *config.Config {
	return fm.config
}

// GetDatabase implements FakeManager.
func (fm *fakeManager) GetDatabase() *gorm.DB {
	return fm.dbGorm
}

// GetServer implements FakeManager.
func (fm *fakeManager) GetServer() *server.Server {
	return fm.server
}

func NewFakeInit(ctrl *gomock.Controller) (FakeManager, error) {
	cfg := &config.Config{}

	srv := server.NewServer(cfg)
	dbMysql, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	defer dbMysql.Close()
	dialector := mysql.New(mysql.Config{
		Conn:                      dbMysql,
		SkipInitializeWithVersion: true,
	})
	dbGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	jwt := jwtAuth.NewMockJwt(ctrl)
	middleware := middleware.NewMockMiddleware(ctrl)
	pagination := paginationHelper.NewMockPagination(ctrl)
	return &fakeManager{
		config:         cfg,
		server:         srv,
		dbGorm:         dbGorm,
		jwtAuth:        jwt,
		middlewareAuth: middleware,
		pagination:     pagination,
	}, nil
}
