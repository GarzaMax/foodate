// Code generated by DIGEN; DO NOT EDIT.
// This file was generated by Dependency Injection Container Generator  (built at ).
// See docs at https://github.com/strider2038/digen

package lookup

import (
	"cmd/app/config"
	gathering_place_api "cmd/app/entities/gatheringPlace/api"
	gathering_place_repository "cmd/app/entities/gatheringPlace/repository"
	gathering_place_usecase "cmd/app/entities/gatheringPlace/usecases"
	meeting_api "cmd/app/entities/meeting/api"
	meeting_repository "cmd/app/entities/meeting/repository"
	meeting_usecase "cmd/app/entities/meeting/usecases"
	user_api "cmd/app/entities/user/api"
	user_repository "cmd/app/entities/user/repository"
	user_usecase "cmd/app/entities/user/usecases"
	"context"
	"database/sql"
	chi "github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Container interface {
	// SetError sets the first error into container. The error is used in the public container to return an initialization error.
	SetError(err error)

	Config(ctx context.Context) config.Params
	Logger(ctx context.Context) *log.Logger
	DB(ctx context.Context) *sql.DB
	Server(ctx context.Context) *http.Server
	Router(ctx context.Context) *chi.Mux

	API() APIContainer
	UseCases() UseCaseContainer
	Repositories() RepositoryContainer
}

type APIContainer interface {
	CreateUserHandler(ctx context.Context) *user_api.CreateUserHandler
	UpdateUserHandler(ctx context.Context) *user_api.UpdateUserHandler
	DeleteUserHandler(ctx context.Context) *user_api.DeleteUserHandler
	FindUserHandler(ctx context.Context) *user_api.FindUserByIdHandler
	FindUsersHandler(ctx context.Context) *user_api.FindUsersByCriteriaHandler
	FindMeetingHandler(ctx context.Context) *meeting_api.FindMeetingByIdHandler
	CreateMeetingHandler(ctx context.Context) *meeting_api.CreateMeetingHandler
	UpdateMeetingHandler(ctx context.Context) *meeting_api.UpdateMeetingHandler
	DeleteMeetingHandler(ctx context.Context) *meeting_api.DeleteMeetingHandler
	FindGatheringPlaceHandler(ctx context.Context) *gathering_place_api.FindGatheringPlaceByIdHandler
	CreateGatheringPlaceHandler(ctx context.Context) *gathering_place_api.CreateGatheringPlaceHandler
	UpdateGatheringPlaceHandler(ctx context.Context) *gathering_place_api.UpdateGatheringPlaceHandler
	DeleteGatheringPlaceHandler(ctx context.Context) *gathering_place_api.DeleteGatheringPlaceHandler
}

type UseCaseContainer interface {
	FindUser(ctx context.Context) *user_usecase.FindUserByIdUseCase
	FindUsers(ctx context.Context) *user_usecase.FindUsersByCriteriaUseCase
	CreateUser(ctx context.Context) *user_usecase.CreateUserUseCase
	UpdateUser(ctx context.Context) *user_usecase.UpdateUserUseCase
	DeleteUser(ctx context.Context) *user_usecase.DeleteUserUseCase
	FindMeeting(ctx context.Context) *meeting_usecase.FindMeetingByIdUseCase
	CreateMeeting(ctx context.Context) *meeting_usecase.CreateMeetingUseCase
	UpdateMeeting(ctx context.Context) *meeting_usecase.UpdateMeetingUseCase
	DeleteMeeting(ctx context.Context) *meeting_usecase.DeleteMeetingUseCase
	FindGatheringPlace(ctx context.Context) *gathering_place_usecase.FindGatheringPlaceByIdUseCase
	CreateGatheringPlace(ctx context.Context) *gathering_place_usecase.CreateGatheringPlaceUseCase
	UpdateGatheringPlace(ctx context.Context) *gathering_place_usecase.UpdateGatheringPlaceUseCase
	DeleteGatheringPlace(ctx context.Context) *gathering_place_usecase.DeleteGatheringPlaceUseCase
}

type RepositoryContainer interface {
	MeetingRepository(ctx context.Context) meeting_repository.MeetingsRepository
	UserRepository(ctx context.Context) user_repository.UsersRepository
	GatheringPlaceRepository(ctx context.Context) gathering_place_repository.PlacesRepository
}
