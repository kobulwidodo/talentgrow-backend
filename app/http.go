package main

import (
	"fmt"
	"log"
	"os"
	"talentgrow-backend/domain"
	_eventHttpHandler "talentgrow-backend/event/delivery/http"
	_eventRepository "talentgrow-backend/event/repository/postgresql"
	_eventUsecase "talentgrow-backend/event/usecase"
	_eventParticipantHttpHandler "talentgrow-backend/event_participant/delivery/http"
	_eventParticipantRepository "talentgrow-backend/event_participant/repository/postgresql"
	_eventParticipantUsecase "talentgrow-backend/event_participant/usecase"
	_internshipHttpHandler "talentgrow-backend/internship/delivery/http"
	_internshipRepository "talentgrow-backend/internship/repository/postgresql"
	_internshipUsecase "talentgrow-backend/internship/usecase"
	_internshipApplicantHttpHandler "talentgrow-backend/internship_applicant/delivery/http"
	_internshipApplicantRepository "talentgrow-backend/internship_applicant/repository/postgresql"
	_internshipApplicantUsecase "talentgrow-backend/internship_applicant/usecase"
	"talentgrow-backend/middleware"
	_userHttpHandler "talentgrow-backend/user/delivery/http"
	_userRepository "talentgrow-backend/user/repository/postgresql"
	_userUsecase "talentgrow-backend/user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env from local")
	}
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	r := gin.Default()
	r.Static("/cv", "./cv")
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	api := r.Group("/api")

	jwtMiddleware := middleware.NewAuthMiddleware()
	mustAdminMiddleware := middleware.MustAdmin()

	userRepository := _userRepository.NewUserRepository(db)
	internshipRepository := _internshipRepository.NewInternshipRepository(db)
	internshipApplicantRepository := _internshipApplicantRepository.NewInternshipApplicantPostgresRepository(db)
	eventRepository := _eventRepository.NewEventRepository(db)
	eventParticipantRepository := _eventParticipantRepository.NewEventParticipantRepository(db)

	userUseCase := _userUsecase.NewUserUseCase(userRepository)
	internshipUsecase := _internshipUsecase.NewInternshipUseCase(internshipRepository)
	internshipApplicantUsecase := _internshipApplicantUsecase.NewInternshipApplicantUseCase(internshipApplicantRepository, internshipRepository)
	eventUsecase := _eventUsecase.NewEventRepository(eventRepository)
	eventParticipantUsecase := _eventParticipantUsecase.NewEventParticipantUsecase(eventParticipantRepository)

	_userHttpHandler.NewUserHandler(api, userUseCase, jwtMiddleware)
	_internshipHttpHandler.NewInternshipHandler(api, internshipUsecase, jwtMiddleware, mustAdminMiddleware)
	_internshipApplicantHttpHandler.NewInternshipApplicantHandler(api, internshipApplicantUsecase, jwtMiddleware)
	_eventHttpHandler.NewEventHandler(api, eventUsecase, jwtMiddleware, mustAdminMiddleware)
	_eventParticipantHttpHandler.NewEventParticipantHandler(api, eventParticipantUsecase, jwtMiddleware)

	r.Run()
}

func initDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
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

	if err := db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Internship{}, &domain.Partnership{}, &domain.EventParticipant{}, &domain.InternshipApplicant{}); err != nil {
		return nil, err
	}

	return db, nil
}
