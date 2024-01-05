package repositories

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/humamalamin/test-case-dating/api/domains/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewAuthRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	dbGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}

	type args struct {
		database *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want interfaces.AuthRepository
	}{
		{
			name: "success",
			args: args{
				database: dbGorm,
			},
			want: &authRepo{
				DB: dbGorm,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepository(tt.args.database); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}

}
