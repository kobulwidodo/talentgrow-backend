package main

import (
	"fmt"
	"log"
	"os"
	"talentgrow-backend/domain"
	_internshipHttpHandler "talentgrow-backend/internship/delivery/http"
	_internshipRepository "talentgrow-backend/internship/repository/postgresql"
	_internshipUsecase "talentgrow-backend/internship/usecase"
	"talentgrow-backend/middleware"
	_userHttpHandler "talentgrow-backend/user/delivery/http"
	_userRepository "talentgrow-backend/user/repository"
	_userUsecase "talentgrow-backend/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env")
		panic(err)
	}
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	r := gin.Default()
	api := r.Group("/api")

	jwtMiddleware := middleware.NewAuthMiddleware()
	mustAdminMiddleware := middleware.MustAdmin()

	userRepository := _userRepository.NewUserRepository(db)
	internshipRepository := _internshipRepository.NewInternshipRepository(db)

	userUseCase := _userUsecase.NewUserUseCase(userRepository)
	internshipUsecase := _internshipUsecase.NewInternshipUseCase(internshipRepository)

	_userHttpHandler.NewUserHandler(api, userUseCase)
	_internshipHttpHandler.NewInternshipHandler(api, internshipUsecase, jwtMiddleware, mustAdminMiddleware)

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Internship{}, &domain.Partnership{}, &domain.EventParticipat{}, &domain.InternshipApplicant{}); err != nil {
		return nil, err
	}

	return db, nil
}
