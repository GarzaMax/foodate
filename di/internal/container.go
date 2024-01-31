// Code generated by DIGEN; DO NOT EDIT.
// This file was generated by Dependency Injection Container Generator  (built at ).
// See docs at https://github.com/strider2038/digen

package internal

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
	"cmd/di/internal/factories"
	"cmd/di/internal/lookup"
	"context"
	"database/sql"
	chi "github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Container struct {
	err error

	config config.Params
	logger *log.Logger
	db     *sql.DB
	server *http.Server
	router *chi.Mux

	api          *APIContainer
	useCases     *UseCaseContainer
	repositories *RepositoryContainer
}

func NewContainer() *Container {
	c := &Container{}
	c.api = &APIContainer{Container: c}
	c.useCases = &UseCaseContainer{Container: c}
	c.repositories = &RepositoryContainer{Container: c}

	return c
}

// Error returns the first initialization error, which can be set via SetError in a service definition.
func (c *Container) Error() error {
	return c.err
}

// SetError sets the first error into container. The error is used in the public container to return an initialization error.
func (c *Container) SetError(err error) {
	if err != nil && c.err == nil {
		c.err = err
	}
}

type APIContainer struct {
	*Container

	createUserHandler           *user_api.CreateUserHandler
	updateUserHandler           *user_api.UpdateUserHandler
	deleteUserHandler           *user_api.DeleteUserHandler
	findUserHandler             *user_api.FindUserByIdHandler
	findUsersHandler            *user_api.FindUsersByCriteriaHandler
	findMeetingHandler          *meeting_api.FindMeetingByIdHandler
	createMeetingHandler        *meeting_api.CreateMeetingHandler
	updateMeetingHandler        *meeting_api.UpdateMeetingHandler
	deleteMeetingHandler        *meeting_api.DeleteMeetingHandler
	findGatheringPlaceHandler   *gathering_place_api.FindGatheringPlaceByIdHandler
	findGatheringPlacesHandler  *gathering_place_api.FindGatheringPlacesByCriteriaHandler
	createGatheringPlaceHandler *gathering_place_api.CreateGatheringPlaceHandler
	updateGatheringPlaceHandler *gathering_place_api.UpdateGatheringPlaceHandler
	deleteGatheringPlaceHandler *gathering_place_api.DeleteGatheringPlaceHandler
}

type UseCaseContainer struct {
	*Container

	findUser             *user_usecase.FindUserByIdUseCase
	findUsers            *user_usecase.FindUsersByCriteriaUseCase
	createUser           *user_usecase.CreateUserUseCase
	updateUser           *user_usecase.UpdateUserUseCase
	deleteUser           *user_usecase.DeleteUserUseCase
	findMeeting          *meeting_usecase.FindMeetingByIdUseCase
	createMeeting        *meeting_usecase.CreateMeetingUseCase
	updateMeeting        *meeting_usecase.UpdateMeetingUseCase
	deleteMeeting        *meeting_usecase.DeleteMeetingUseCase
	findGatheringPlace   *gathering_place_usecase.FindGatheringPlaceByIdUseCase
	findGatheringPlaces  *gathering_place_usecase.FindGatheringPlacesByCriteriaUseCase
	createGatheringPlace *gathering_place_usecase.CreateGatheringPlaceUseCase
	updateGatheringPlace *gathering_place_usecase.UpdateGatheringPlaceUseCase
	deleteGatheringPlace *gathering_place_usecase.DeleteGatheringPlaceUseCase
}

type RepositoryContainer struct {
	*Container

	meetingRepository        meeting_repository.MeetingsRepository
	userRepository           user_repository.UsersRepository
	gatheringPlaceRepository gathering_place_repository.PlacesRepository
}

func (c *Container) Config(ctx context.Context) config.Params {
	return c.config
}

func (c *Container) Logger(ctx context.Context) *log.Logger {
	if c.logger == nil && c.err == nil {
		c.logger = factories.CreateLogger(ctx, c)
	}
	return c.logger
}

func (c *Container) DB(ctx context.Context) *sql.DB {
	if c.db == nil && c.err == nil {
		c.db = factories.CreateDB(ctx, c)
	}
	return c.db
}

func (c *Container) Server(ctx context.Context) *http.Server {
	if c.server == nil && c.err == nil {
		c.server = factories.CreateServer(ctx, c)
	}
	return c.server
}

func (c *Container) Router(ctx context.Context) *chi.Mux {
	if c.router == nil && c.err == nil {
		c.router = factories.CreateRouter(ctx, c)
	}
	return c.router
}

func (c *Container) API() lookup.APIContainer {
	return c.api
}

func (c *APIContainer) CreateUserHandler(ctx context.Context) *user_api.CreateUserHandler {
	if c.createUserHandler == nil && c.err == nil {
		c.createUserHandler = factories.CreateAPICreateUserHandler(ctx, c)
	}
	return c.createUserHandler
}

func (c *APIContainer) UpdateUserHandler(ctx context.Context) *user_api.UpdateUserHandler {
	if c.updateUserHandler == nil && c.err == nil {
		c.updateUserHandler = factories.CreateAPIUpdateUserHandler(ctx, c)
	}
	return c.updateUserHandler
}

func (c *APIContainer) DeleteUserHandler(ctx context.Context) *user_api.DeleteUserHandler {
	if c.deleteUserHandler == nil && c.err == nil {
		c.deleteUserHandler = factories.CreateAPIDeleteUserHandler(ctx, c)
	}
	return c.deleteUserHandler
}

func (c *APIContainer) FindUserHandler(ctx context.Context) *user_api.FindUserByIdHandler {
	if c.findUserHandler == nil && c.err == nil {
		c.findUserHandler = factories.CreateAPIFindUserHandler(ctx, c)
	}
	return c.findUserHandler
}

func (c *APIContainer) FindUsersHandler(ctx context.Context) *user_api.FindUsersByCriteriaHandler {
	if c.findUsersHandler == nil && c.err == nil {
		c.findUsersHandler = factories.CreateAPIFindUsersHandler(ctx, c)
	}
	return c.findUsersHandler
}

func (c *APIContainer) FindMeetingHandler(ctx context.Context) *meeting_api.FindMeetingByIdHandler {
	if c.findMeetingHandler == nil && c.err == nil {
		c.findMeetingHandler = factories.CreateAPIFindMeetingHandler(ctx, c)
	}
	return c.findMeetingHandler
}

func (c *APIContainer) CreateMeetingHandler(ctx context.Context) *meeting_api.CreateMeetingHandler {
	if c.createMeetingHandler == nil && c.err == nil {
		c.createMeetingHandler = factories.CreateAPICreateMeetingHandler(ctx, c)
	}
	return c.createMeetingHandler
}

func (c *APIContainer) UpdateMeetingHandler(ctx context.Context) *meeting_api.UpdateMeetingHandler {
	if c.updateMeetingHandler == nil && c.err == nil {
		c.updateMeetingHandler = factories.CreateAPIUpdateMeetingHandler(ctx, c)
	}
	return c.updateMeetingHandler
}

func (c *APIContainer) DeleteMeetingHandler(ctx context.Context) *meeting_api.DeleteMeetingHandler {
	if c.deleteMeetingHandler == nil && c.err == nil {
		c.deleteMeetingHandler = factories.CreateAPIDeleteMeetingHandler(ctx, c)
	}
	return c.deleteMeetingHandler
}

func (c *APIContainer) FindGatheringPlaceHandler(ctx context.Context) *gathering_place_api.FindGatheringPlaceByIdHandler {
	if c.findGatheringPlaceHandler == nil && c.err == nil {
		c.findGatheringPlaceHandler = factories.CreateAPIFindGatheringPlaceHandler(ctx, c)
	}
	return c.findGatheringPlaceHandler
}

func (c *APIContainer) FindGatheringPlacesHandler(ctx context.Context) *gathering_place_api.FindGatheringPlacesByCriteriaHandler {
	if c.findGatheringPlacesHandler == nil && c.err == nil {
		c.findGatheringPlacesHandler = factories.CreateAPIFindGatheringPlacesHandler(ctx, c)
	}
	return c.findGatheringPlacesHandler
}

func (c *APIContainer) CreateGatheringPlaceHandler(ctx context.Context) *gathering_place_api.CreateGatheringPlaceHandler {
	if c.createGatheringPlaceHandler == nil && c.err == nil {
		c.createGatheringPlaceHandler = factories.CreateAPICreateGatheringPlaceHandler(ctx, c)
	}
	return c.createGatheringPlaceHandler
}

func (c *APIContainer) UpdateGatheringPlaceHandler(ctx context.Context) *gathering_place_api.UpdateGatheringPlaceHandler {
	if c.updateGatheringPlaceHandler == nil && c.err == nil {
		c.updateGatheringPlaceHandler = factories.CreateAPIUpdateGatheringPlaceHandler(ctx, c)
	}
	return c.updateGatheringPlaceHandler
}

func (c *APIContainer) DeleteGatheringPlaceHandler(ctx context.Context) *gathering_place_api.DeleteGatheringPlaceHandler {
	if c.deleteGatheringPlaceHandler == nil && c.err == nil {
		c.deleteGatheringPlaceHandler = factories.CreateAPIDeleteGatheringPlaceHandler(ctx, c)
	}
	return c.deleteGatheringPlaceHandler
}

func (c *Container) UseCases() lookup.UseCaseContainer {
	return c.useCases
}

func (c *UseCaseContainer) FindUser(ctx context.Context) *user_usecase.FindUserByIdUseCase {
	if c.findUser == nil && c.err == nil {
		c.findUser = factories.CreateUseCasesFindUser(ctx, c)
	}
	return c.findUser
}

func (c *UseCaseContainer) FindUsers(ctx context.Context) *user_usecase.FindUsersByCriteriaUseCase {
	if c.findUsers == nil && c.err == nil {
		c.findUsers = factories.CreateUseCasesFindUsers(ctx, c)
	}
	return c.findUsers
}

func (c *UseCaseContainer) CreateUser(ctx context.Context) *user_usecase.CreateUserUseCase {
	if c.createUser == nil && c.err == nil {
		c.createUser = factories.CreateUseCasesCreateUser(ctx, c)
	}
	return c.createUser
}

func (c *UseCaseContainer) UpdateUser(ctx context.Context) *user_usecase.UpdateUserUseCase {
	if c.updateUser == nil && c.err == nil {
		c.updateUser = factories.CreateUseCasesUpdateUser(ctx, c)
	}
	return c.updateUser
}

func (c *UseCaseContainer) DeleteUser(ctx context.Context) *user_usecase.DeleteUserUseCase {
	if c.deleteUser == nil && c.err == nil {
		c.deleteUser = factories.CreateUseCasesDeleteUser(ctx, c)
	}
	return c.deleteUser
}

func (c *UseCaseContainer) FindMeeting(ctx context.Context) *meeting_usecase.FindMeetingByIdUseCase {
	if c.findMeeting == nil && c.err == nil {
		c.findMeeting = factories.CreateUseCasesFindMeeting(ctx, c)
	}
	return c.findMeeting
}

func (c *UseCaseContainer) CreateMeeting(ctx context.Context) *meeting_usecase.CreateMeetingUseCase {
	if c.createMeeting == nil && c.err == nil {
		c.createMeeting = factories.CreateUseCasesCreateMeeting(ctx, c)
	}
	return c.createMeeting
}

func (c *UseCaseContainer) UpdateMeeting(ctx context.Context) *meeting_usecase.UpdateMeetingUseCase {
	if c.updateMeeting == nil && c.err == nil {
		c.updateMeeting = factories.CreateUseCasesUpdateMeeting(ctx, c)
	}
	return c.updateMeeting
}

func (c *UseCaseContainer) DeleteMeeting(ctx context.Context) *meeting_usecase.DeleteMeetingUseCase {
	if c.deleteMeeting == nil && c.err == nil {
		c.deleteMeeting = factories.CreateUseCasesDeleteMeeting(ctx, c)
	}
	return c.deleteMeeting
}

func (c *UseCaseContainer) FindGatheringPlace(ctx context.Context) *gathering_place_usecase.FindGatheringPlaceByIdUseCase {
	if c.findGatheringPlace == nil && c.err == nil {
		c.findGatheringPlace = factories.CreateUseCasesFindGatheringPlace(ctx, c)
	}
	return c.findGatheringPlace
}

func (c *UseCaseContainer) FindGatheringPlaces(ctx context.Context) *gathering_place_usecase.FindGatheringPlacesByCriteriaUseCase {
	if c.findGatheringPlaces == nil && c.err == nil {
		c.findGatheringPlaces = factories.CreateUseCasesFindGatheringPlaces(ctx, c)
	}
	return c.findGatheringPlaces
}

func (c *UseCaseContainer) CreateGatheringPlace(ctx context.Context) *gathering_place_usecase.CreateGatheringPlaceUseCase {
	if c.createGatheringPlace == nil && c.err == nil {
		c.createGatheringPlace = factories.CreateUseCasesCreateGatheringPlace(ctx, c)
	}
	return c.createGatheringPlace
}

func (c *UseCaseContainer) UpdateGatheringPlace(ctx context.Context) *gathering_place_usecase.UpdateGatheringPlaceUseCase {
	if c.updateGatheringPlace == nil && c.err == nil {
		c.updateGatheringPlace = factories.CreateUseCasesUpdateGatheringPlace(ctx, c)
	}
	return c.updateGatheringPlace
}

func (c *UseCaseContainer) DeleteGatheringPlace(ctx context.Context) *gathering_place_usecase.DeleteGatheringPlaceUseCase {
	if c.deleteGatheringPlace == nil && c.err == nil {
		c.deleteGatheringPlace = factories.CreateUseCasesDeleteGatheringPlace(ctx, c)
	}
	return c.deleteGatheringPlace
}

func (c *Container) Repositories() lookup.RepositoryContainer {
	return c.repositories
}

func (c *RepositoryContainer) MeetingRepository(ctx context.Context) meeting_repository.MeetingsRepository {
	if c.meetingRepository == nil && c.err == nil {
		c.meetingRepository = factories.CreateRepositoriesMeetingRepository(ctx, c)
	}
	return c.meetingRepository
}

func (c *RepositoryContainer) UserRepository(ctx context.Context) user_repository.UsersRepository {
	if c.userRepository == nil && c.err == nil {
		c.userRepository = factories.CreateRepositoriesUserRepository(ctx, c)
	}
	return c.userRepository
}

func (c *RepositoryContainer) GatheringPlaceRepository(ctx context.Context) gathering_place_repository.PlacesRepository {
	if c.gatheringPlaceRepository == nil && c.err == nil {
		c.gatheringPlaceRepository = factories.CreateRepositoriesGatheringPlaceRepository(ctx, c)
	}
	return c.gatheringPlaceRepository
}

func (c *Container) SetConfig(s config.Params) {
	c.config = s
}

func (c *RepositoryContainer) SetMeetingRepository(s meeting_repository.MeetingsRepository) {
	c.meetingRepository = s
}

func (c *RepositoryContainer) SetUserRepository(s user_repository.UsersRepository) {
	c.userRepository = s
}

func (c *RepositoryContainer) SetGatheringPlaceRepository(s gathering_place_repository.PlacesRepository) {
	c.gatheringPlaceRepository = s
}

func (c *Container) Close() {
	if c.db != nil {
		c.db.Close()
	}

	if c.server != nil {
		c.server.Close()
	}
}
