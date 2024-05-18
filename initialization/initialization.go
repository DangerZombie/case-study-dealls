package initialization

import (
	"fmt"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/database"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	transport "github.com/DangerZombie/case-study-dealls/http"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func DbInit() (db *gorm.DB, err error) {
	// Initial DB Connection
	driver := viper.GetString("database.driver")
	dbName := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	port := viper.GetInt("database.port")

	db, err = database.NewDBConnection(driver, dbName, host, username, password, port)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{}, &entity.UserInteraction{}, &entity.Match{})

	return db, nil
}

func UserCronInit(db *gorm.DB, userService user_service.UserService) (err error) {
	c := cron.New()

	c.AddFunc("* * * * *", func() {
		resetSwipeCountInput := request.ResetSwipeCountRequest{
			Status:     static.UserFree,
			SwipeCount: 10,
		}

		result, err := userService.ResetSwipeCount(resetSwipeCountInput)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result.Message)
	})

	c.Start()

	return nil
}

func ServerInit(db *gorm.DB) {
	// base repository
	baseRepository := repository.NewBaseRepository(db)
	userRepository := user_repository.NewUserRepository(baseRepository)

	// auth helper
	authHelper := auth.NewAuthHelper(
		baseRepository,
		userRepository,
	)

	// service
	userSvc := user_service.NewUserService(
		authHelper,
		baseRepository,
		userRepository,
	)

	r := echo.New()

	// group endpoint
	apiGroupUser := r.Group("/api/v1/user")

	// transport
	transportHandler := transport.NewHttp(
		authHelper,
	)

	transportHandler.UserHandler(apiGroupUser, userSvc)

	// run cron
	UserCronInit(db, userSvc)

	r.Start(":9000")
}
