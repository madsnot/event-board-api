package usecase

import (
	errorPkg "github.com/madsnot/event-board-api/pkg/error"
)

const (
	ErrInternalCode             = 1000
	ErrUserNotFoundCode         = 1001
	ErrInvalidEventCode         = 1002
	ErrInvalidStartDateCode     = 1003
	ErrDocsNotFoundCode         = 1004
	ErrInvalidToGetEventsCode   = 1005
	ErrInvalidToCreateEventCode = 1006
	ErrInvalidToCreateIndexCode = 1007
	ErrInvalidToPublishMsgCode  = 1008
)

var (
	ErrInternal             = errorPkg.NewBusinessError(ErrInternalCode, "internal business error")
	ErrUserNotFound         = errorPkg.NewBusinessError(ErrUserNotFoundCode, "user not found")
	ErrInvalidEvent         = errorPkg.NewBusinessError(ErrInvalidEventCode, "invalid event")
	ErrInvalidStartDate     = errorPkg.NewBusinessError(ErrInvalidStartDateCode, "invalid start date")
	ErrDocsNotFound         = errorPkg.NewBusinessError(ErrDocsNotFoundCode, "docs not found")
	ErrInvalidToGetEvents   = errorPkg.NewBusinessError(ErrInvalidToGetEventsCode, "invalid to get events")
	ErrInvalidToCreateEvent = errorPkg.NewBusinessError(ErrInvalidToCreateEventCode, "invalid to create event")
	ErrInvalidToCreateIndex = errorPkg.NewBusinessError(ErrInvalidToCreateIndexCode, "invalid to create index")
	ErrInvalidToPublishMsg  = errorPkg.NewBusinessError(ErrInvalidToPublishMsgCode, "invalid to publish msg")
)
