package unittest_services

import (
	"testing"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	mockitem_repositories "github.com/Prompiriya084/go-authen/Test/UnitTest/MockItem/Repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	userid1 := uuid.New()
	userid2 := uuid.New()
	userauthid1 := uuid.New()
	userauthid2 := uuid.New()
	mockData := []entities.User{
		{ID: userid1, Name: "Name1", Surname: "Surname1", UserAuthID: userauthid1},
		{ID: userid2, Name: "Name2", Surname: "Surname2", UserAuthID: userauthid2},
	}
	testcase := []struct {
		description       string
		sendingFilter     *entities.User
		mockReturnedData  []entities.User
		mockReturnedError error
		expection         int //Orders length
	}{
		{
			description:       "[OK]Sending nothing",
			sendingFilter:     nil,
			mockReturnedData:  mockData,
			mockReturnedError: nil,
			expection:         2,
		},
		{
			description: "[OK]Sending correct params",
			sendingFilter: &entities.User{
				Name: "Name1",
			},
			mockReturnedData: []entities.User{
				mockData[1],
			},
			mockReturnedError: nil,
			expection:         1,
		},
		{
			description: "[Error]User not found.",
			sendingFilter: &entities.User{
				ID: uuid.New(),
			},
			mockReturnedData:  nil,
			mockReturnedError: nil,
			expection:         1,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.description, func(t *testing.T) {
			mockRepo := &mockitem_repositories.MockUserRepository{
				MockRepositoryImpl: mockitem_repositories.MockRepositoryImpl[entities.User]{
					GetAllFn: func(filters *entities.User, preload []string) ([]entities.User, error) {
						return tc.mockReturnedData, tc.mockReturnedError
					},
				},
			}
			userService := services.NewUserService(mockRepo)
			response, err := userService.GetUserAll(tc.sendingFilter)

			assert.NoError(t, err)
			assert.Equal(t, tc.expection, len(response))
		})
	}
}
func TestGetById(t *testing.T) {
	userid := uuid.New()
	useridString := userid.String()
	userauthId := uuid.New()
	mockData := &entities.User{
		ID:         userid,
		Name:       "Name1",
		Surname:    "Surname1",
		UserAuthID: userauthId,
	}

	testcase := []struct {
		description       string
		sendingFilter     string
		mockReturnedData  *entities.User
		mockReturnedError error
		expection         *entities.User
	}{
		{
			description:       "[OK]Sending nothing",
			sendingFilter:     "",
			mockReturnedData:  nil,
			mockReturnedError: nil,
			expection:         nil,
		},
		{
			description:       "[OK]Sending id params",
			sendingFilter:     useridString,
			mockReturnedData:  mockData,
			mockReturnedError: nil,
			expection:         mockData,
		},

		{
			description:       "[Error]User not found.",
			sendingFilter:     uuid.New().String(),
			mockReturnedData:  nil,
			mockReturnedError: nil,
			expection:         nil,
		},
	}
	for _, tc := range testcase {
		t.Run(tc.description, func(t *testing.T) {
			mockRepo := &mockitem_repositories.MockUserRepository{
				MockRepositoryImpl: mockitem_repositories.MockRepositoryImpl[entities.User]{
					GetFn: func(filters *entities.User, preload []string) (*entities.User, error) {
						return tc.mockReturnedData, tc.mockReturnedError
					},
				},
			}
			userService := services.NewUserService(mockRepo)
			response, err := userService.GetUserById(tc.sendingFilter)

			assert.NoError(t, err)
			assert.Equal(t, tc.expection, response)
		})
	}
}
func TestCreate(t *testing.T) {

}
