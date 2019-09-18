// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import time "time"

// ElasticsearchInterface is an autogenerated mock type for the ElasticsearchInterface type
type ElasticsearchInterface struct {
	mock.Mock
}

// DataRetentionDeleteIndexes provides a mock function with given fields: cutoff
func (_m *ElasticsearchInterface) DataRetentionDeleteIndexes(cutoff time.Time) error {
	ret := _m.Called(cutoff)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(time.Time) *model.AppError); ok {
		r0 = rf(cutoff)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteChannel provides a mock function with given fields: channel
func (_m *ElasticsearchInterface) DeleteChannel(channel *model.Channel) error {
	ret := _m.Called(channel)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Channel) *model.AppError); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeletePost provides a mock function with given fields: post
func (_m *ElasticsearchInterface) DeletePost(post *model.Post) error {
	ret := _m.Called(post)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post) *model.AppError); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// DeleteUser provides a mock function with given fields: user
func (_m *ElasticsearchInterface) DeleteUser(user *model.User) error {
	ret := _m.Called(user)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User) *model.AppError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexChannel provides a mock function with given fields: channel
func (_m *ElasticsearchInterface) IndexChannel(channel *model.Channel) error {
	ret := _m.Called(channel)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Channel) *model.AppError); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexPost provides a mock function with given fields: post, teamId
func (_m *ElasticsearchInterface) IndexPost(post *model.Post, teamId string) error {
	ret := _m.Called(post, teamId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Post, string) *model.AppError); ok {
		r0 = rf(post, teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// IndexUser provides a mock function with given fields: user, teamsIds, channelsIds
func (_m *ElasticsearchInterface) IndexUser(user *model.User, teamsIds []string, channelsIds []string) error {
	ret := _m.Called(user, teamsIds, channelsIds)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User, []string, []string) *model.AppError); ok {
		r0 = rf(user, teamsIds, channelsIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PurgeIndexes provides a mock function with given fields:
func (_m *ElasticsearchInterface) PurgeIndexes() error {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// SearchChannels provides a mock function with given fields: teamId, term
func (_m *ElasticsearchInterface) SearchChannels(teamId string, term string) ([]string, error) {
	ret := _m.Called(teamId, term)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(teamId, term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(teamId, term)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchPosts provides a mock function with given fields: channels, searchParams, page, perPage
func (_m *ElasticsearchInterface) SearchPosts(channels *model.ChannelList, searchParams []*model.SearchParams, page int, perPage int) ([]string, model.PostSearchMatches, error) {
	ret := _m.Called(channels, searchParams, page, perPage)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*model.ChannelList, []*model.SearchParams, int, int) []string); ok {
		r0 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 model.PostSearchMatches
	if rf, ok := ret.Get(1).(func(*model.ChannelList, []*model.SearchParams, int, int) model.PostSearchMatches); ok {
		r1 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(model.PostSearchMatches)
		}
	}

	var r2 *model.AppError
	if rf, ok := ret.Get(2).(func(*model.ChannelList, []*model.SearchParams, int, int) *model.AppError); ok {
		r2 = rf(channels, searchParams, page, perPage)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInChannel provides a mock function with given fields: teamId, channelId, restrictedToChannels, term, options
func (_m *ElasticsearchInterface) SearchUsersInChannel(teamId string, channelId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, []string, error) {
	ret := _m.Called(teamId, channelId, restrictedToChannels, term, options)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string, string, []string, string, *model.UserSearchOptions) []string); ok {
		r1 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 *model.AppError
	if rf, ok := ret.Get(2).(func(string, string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r2 = rf(teamId, channelId, restrictedToChannels, term, options)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// SearchUsersInTeam provides a mock function with given fields: teamId, restrictedToChannels, term, options
func (_m *ElasticsearchInterface) SearchUsersInTeam(teamId string, restrictedToChannels []string, term string, options *model.UserSearchOptions) ([]string, error) {
	ret := _m.Called(teamId, restrictedToChannels, term, options)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, []string, string, *model.UserSearchOptions) []string); ok {
		r0 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(teamId, restrictedToChannels, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *ElasticsearchInterface) Start() error {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *ElasticsearchInterface) Stop() error {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// TestConfig provides a mock function with given fields: cfg
func (_m *ElasticsearchInterface) TestConfig(cfg *model.Config) error {
	ret := _m.Called(cfg)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Config) *model.AppError); ok {
		r0 = rf(cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}
