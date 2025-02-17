package notifservice_test

import (
	"errors"
	"project_pos_app/helper"
	"project_pos_app/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateNotification(t *testing.T) {
	t.Run("Successfully create a notification", func(t *testing.T) {
		now := time.Now()
		dynamicDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		input := model.Notification{
			ID:        1,
			Title:     "Testing",
			Message:   "Test notification",
			Status:    "new",
			CreatedAt: dynamicDate,
			UpdatedAt: dynamicDate,
		}

		mockDB, service := helper.InitService()
		mockDB.On("Create", input).Once().Return(nil)
		err := service.Notif.CreateNotification(input)
		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	})

	t.Run("Failed create a notification invalid title", func(t *testing.T) {
		input := model.Notification{}

		mockDB, service := helper.InitService()
		mockDB.On("Create", input).Return(nil)
		err := service.Notif.CreateNotification(input)
		assert.Error(t, err)
	})
}

func TestGetAllNotifications(t *testing.T) {
	mockDB, service := helper.InitService()
	now := time.Now()

	t.Run("Successfully get all notification", func(t *testing.T) {
		notifications := []model.Notification{
			{
				ID:        1,
				Title:     "Testing",
				Message:   "Test notification",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				Title:     "Testing2",
				Message:   "Test notification2",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		// Ekspektasi untuk GetAll
		mockDB.On("GetAll").Once().Return(notifications, nil)

		// Panggil service
		result, err := service.Notif.GetAllNotifications("new")

		// Verifikasi hasil
		assert.NoError(t, err)
		assert.Equal(t, notifications, result)
		mockDB.AssertExpectations(t)
	})

	t.Run("failed get all notification data is not found", func(t *testing.T) {
		notifications := []model.Notification{}

		// Ekspektasi untuk GetAll
		mockDB.On("GetAll").Return(notifications, nil)

		// Panggil service
		result, err := service.Notif.GetAllNotifications("new")

		// Verifikasi hasil
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})
}

func TestFindByID(t *testing.T) {
	mockDB, service := helper.InitService()
	now := time.Now()

	t.Run("Successfully get a notification", func(t *testing.T) {
		expectedNotif := &model.Notification{
			ID:        1,
			Title:     "Testing",
			Message:   "Test notification",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockDB.On("FindByID", 1).Return(expectedNotif, nil)

		result, err := service.Notif.GetNotificationByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedNotif, result)

		mockDB.AssertExpectations(t)
	})

	t.Run("Failed get a notification id invalid", func(t *testing.T) {
		// expectedNotif := model.Notification{ID: 9999}

		mockDB.On("FindByID", 9999).Return(nil, errors.New("invalid id"))

		result, err := service.Notif.GetNotificationByID(9999)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdateNotification(t *testing.T) {
	mockDB, service := helper.InitService()
	now := time.Now()

	t.Run("Successfully update a notification", func(t *testing.T) {
		notif := &model.Notification{
			ID:        1,
			Title:     "Testing",
			Message:   "Test notification",
			Status:    "new",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockDB.On("FindByID", notif.ID).Return(notif, nil)
		mockDB.On("Update", notif, 1).Return(notif, nil)

		err := service.Notif.UpdateNotification(1)

		assert.NoError(t, err)
		assert.Equal(t, "readed", notif.Status)
		mockDB.AssertExpectations(t)
	})

	t.Run("Failed update a notification invalid id", func(t *testing.T) {
		notif := &model.Notification{ID: 999999}

		mockDB.On("FindByID", notif.ID).Return(notif, nil)
		mockDB.On("Update", mock.Anything, notif.ID).Return(nil).Run(func(args mock.Arguments) {
			updatedNotif := args.Get(0).(*model.Notification)
			notif = updatedNotif
		})

		err := service.Notif.UpdateNotification(notif.ID)

		assert.Error(t, err)
		assert.Equal(t, notif.Title, "")
		mockDB.AssertNotCalled(t, "Update", mock.Anything, notif.ID)
	})
}

func TestDeleteNotification(t *testing.T) {
	mockDB, service := helper.InitService()
	now := time.Now()

	t.Run("Successfully delete a notification", func(t *testing.T) {
		notif := &model.Notification{
			ID:        1,
			Title:     "Testing",
			Message:   "Test notification",
			Status:    "new",
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockDB.On("FindByID", 1).Return(notif, nil)
		mockDB.On("Delete", 1).Return(&gorm.DB{Error: nil})

		err := service.Notif.DeleteNotification(1)

		assert.NoError(t, err)
		mockDB.AssertCalled(t, "Delete", 1)
	})

	t.Run("Failed delete a notification", func(t *testing.T) {
		notif := &model.Notification{}

		mockDB.On("FindByID", notif.ID).Return(notif, nil)
		mockDB.On("Delete", notif.ID).Return(&gorm.DB{Error: nil})

		err := service.Notif.DeleteNotification(notif.ID)

		assert.Error(t, err)
		mockDB.AssertNotCalled(t, "Delete", notif.ID)
	})
}

func TestMarkAllAsRead(t *testing.T) {
	mockDB, service := helper.InitService()
	now := time.Now()

	t.Run("Successfully update status all notification", func(t *testing.T) {
		notifications := []model.Notification{
			{
				ID:        1,
				Title:     "Testing",
				Message:   "Test notification",
				Status:    "new",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				Title:     "Testing2",
				Message:   "Test notification2",
				Status:    "new",
				CreatedAt: now,
				UpdatedAt: now,
			},
		}

		mockDB.On("MarkAllAsRead").Once().Return(notifications, nil)

		err := service.Notif.MarkAllNotificationsAsRead()

		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	})

	t.Run("Failed update status all notification", func(t *testing.T) {
		notifications := []model.Notification{}
		mockDB.On("MarkAllAsRead").Return(notifications, nil)
		_ = service.Notif.MarkAllNotificationsAsRead()
		assert.Equal(t, 0, len(notifications))
	})
}
