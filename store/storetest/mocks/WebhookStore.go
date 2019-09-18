// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/model"
	mock "github.com/stretchr/testify/mock"
)

// WebhookStore is an autogenerated mock type for the WebhookStore type
type WebhookStore struct {
	mock.Mock
}

// AnalyticsIncomingCount provides a mock function with given fields: teamId
func (_m *WebhookStore) AnalyticsIncomingCount(teamId string) (int64, error) {
	ret := _m.Called(teamId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(teamId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(teamId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// AnalyticsOutgoingCount provides a mock function with given fields: teamId
func (_m *WebhookStore) AnalyticsOutgoingCount(teamId string) (int64, error) {
	ret := _m.Called(teamId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(teamId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(teamId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ClearCaches provides a mock function with given fields:
func (_m *WebhookStore) ClearCaches() {
	_m.Called()
}

// DeleteIncoming provides a mock function with given fields: webhookId, time
func (_m *WebhookStore) DeleteIncoming(webhookId string, time int64) error {
	ret := _m.Called(webhookId, time)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int64) *model.AppError); ok {
		r0 = rf(webhookId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteOutgoing provides a mock function with given fields: webhookId, time
func (_m *WebhookStore) DeleteOutgoing(webhookId string, time int64) error {
	ret := _m.Called(webhookId, time)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int64) *model.AppError); ok {
		r0 = rf(webhookId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetIncoming provides a mock function with given fields: id, allowFromCache
func (_m *WebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, error) {
	ret := _m.Called(id, allowFromCache)

	var r0 *model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(string, bool) *model.IncomingWebhook); ok {
		r0 = rf(id, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, bool) *model.AppError); ok {
		r1 = rf(id, allowFromCache)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetIncomingByChannel provides a mock function with given fields: channelId
func (_m *WebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, error) {
	ret := _m.Called(channelId)

	var r0 []*model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(string) []*model.IncomingWebhook); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(channelId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetIncomingByTeam provides a mock function with given fields: teamId, offset, limit
func (_m *WebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	ret := _m.Called(teamId, offset, limit)

	var r0 []*model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.IncomingWebhook); ok {
		r0 = rf(teamId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(teamId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetIncomingByTeamByUser provides a mock function with given fields: teamId, userId, offset, limit
func (_m *WebhookStore) GetIncomingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	ret := _m.Called(teamId, userId, offset, limit)

	var r0 []*model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.IncomingWebhook); ok {
		r0 = rf(teamId, userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, int, int) *model.AppError); ok {
		r1 = rf(teamId, userId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetIncomingList provides a mock function with given fields: offset, limit
func (_m *WebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, error) {
	ret := _m.Called(offset, limit)

	var r0 []*model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(int, int) []*model.IncomingWebhook); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, int) *model.AppError); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetIncomingListByUser provides a mock function with given fields: userId, offset, limit
func (_m *WebhookStore) GetIncomingListByUser(userId string, offset int, limit int) ([]*model.IncomingWebhook, error) {
	ret := _m.Called(userId, offset, limit)

	var r0 []*model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.IncomingWebhook); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(userId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoing provides a mock function with given fields: id
func (_m *WebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, error) {
	ret := _m.Called(id)

	var r0 *model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string) *model.OutgoingWebhook); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingByChannel provides a mock function with given fields: channelId, offset, limit
func (_m *WebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(channelId, offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(channelId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(channelId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingByChannelByUser provides a mock function with given fields: channelId, userId, offset, limit
func (_m *WebhookStore) GetOutgoingByChannelByUser(channelId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(channelId, userId, offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(channelId, userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, int, int) *model.AppError); ok {
		r1 = rf(channelId, userId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingByTeam provides a mock function with given fields: teamId, offset, limit
func (_m *WebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(teamId, offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(teamId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(teamId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingByTeamByUser provides a mock function with given fields: teamId, userId, offset, limit
func (_m *WebhookStore) GetOutgoingByTeamByUser(teamId string, userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(teamId, userId, offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(teamId, userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, int, int) *model.AppError); ok {
		r1 = rf(teamId, userId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingList provides a mock function with given fields: offset, limit
func (_m *WebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, int) *model.AppError); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOutgoingListByUser provides a mock function with given fields: userId, offset, limit
func (_m *WebhookStore) GetOutgoingListByUser(userId string, offset int, limit int) ([]*model.OutgoingWebhook, error) {
	ret := _m.Called(userId, offset, limit)

	var r0 []*model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.OutgoingWebhook); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(userId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// InvalidateWebhookCache provides a mock function with given fields: webhook
func (_m *WebhookStore) InvalidateWebhookCache(webhook string) {
	_m.Called(webhook)
}

// PermanentDeleteIncomingByChannel provides a mock function with given fields: channelId
func (_m *WebhookStore) PermanentDeleteIncomingByChannel(channelId string) error {
	ret := _m.Called(channelId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PermanentDeleteIncomingByUser provides a mock function with given fields: userId
func (_m *WebhookStore) PermanentDeleteIncomingByUser(userId string) error {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PermanentDeleteOutgoingByChannel provides a mock function with given fields: channelId
func (_m *WebhookStore) PermanentDeleteOutgoingByChannel(channelId string) error {
	ret := _m.Called(channelId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PermanentDeleteOutgoingByUser provides a mock function with given fields: userId
func (_m *WebhookStore) PermanentDeleteOutgoingByUser(userId string) error {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// SaveIncoming provides a mock function with given fields: webhook
func (_m *WebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	ret := _m.Called(webhook)

	var r0 *model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(*model.IncomingWebhook) *model.IncomingWebhook); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.IncomingWebhook) *model.AppError); ok {
		r1 = rf(webhook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SaveOutgoing provides a mock function with given fields: webhook
func (_m *WebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	ret := _m.Called(webhook)

	var r0 *model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(*model.OutgoingWebhook) *model.OutgoingWebhook); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.OutgoingWebhook) *model.AppError); ok {
		r1 = rf(webhook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateIncoming provides a mock function with given fields: webhook
func (_m *WebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	ret := _m.Called(webhook)

	var r0 *model.IncomingWebhook
	if rf, ok := ret.Get(0).(func(*model.IncomingWebhook) *model.IncomingWebhook); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.IncomingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.IncomingWebhook) *model.AppError); ok {
		r1 = rf(webhook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateOutgoing provides a mock function with given fields: hook
func (_m *WebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	ret := _m.Called(hook)

	var r0 *model.OutgoingWebhook
	if rf, ok := ret.Get(0).(func(*model.OutgoingWebhook) *model.OutgoingWebhook); ok {
		r0 = rf(hook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingWebhook)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.OutgoingWebhook) *model.AppError); ok {
		r1 = rf(hook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}
