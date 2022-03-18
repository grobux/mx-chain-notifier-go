package factory

import (
	"github.com/ElrondNetwork/notifier-go/api/groups"
	"github.com/ElrondNetwork/notifier-go/common"
	"github.com/ElrondNetwork/notifier-go/config"
	"github.com/ElrondNetwork/notifier-go/disabled"
	"github.com/ElrondNetwork/notifier-go/redis"
)

func CreateLockService(apiType common.APIType, config *config.GeneralConfig) (redis.LockService, error) {
	var lockService groups.LockService

	if !config.ConnectorApi.CheckDuplicates || apiType == common.WSAPIType {
		return disabled.NewDisabledRedlockWrapper(), nil
	}

	redisClient, err := redis.CreateFailoverClient(config.PubSub)
	if err != nil {
		return nil, err
	}

	lockService, err = redis.NewRedlockWrapper(redisClient)
	if err != nil {
		return nil, err
	}

	return lockService, nil
}
