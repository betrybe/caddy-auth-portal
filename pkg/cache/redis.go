// Copyright 2020 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cache

import (
	"context"
	"encoding/json"

	"github.com/betrybe/caddy-auth-portal/pkg/errors"
	"github.com/go-redis/redis/v8"
)

// Redis implements StateManager interface using redis as backend.
type Redis struct {
	Servers []string
	client  *redis.Client
	ctx     context.Context
}

// NewRedisCache creates a new key manager using the specified servers.
func NewRedisCache(server []string) *Redis {
	return &Redis{
		Servers: server,
		client:  redis.NewClient(&redis.Options{Addr: server[0]}),
		ctx:     context.Background(),
	}
}

// Add a key and its value to the backend.
func (sm *Redis) Add(key string, value interface{}) error {
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return errors.ErrCache.WithArgs("add", err)
	}
	err = sm.client.Set(sm.ctx, key, encodedValue, 0).Err()
	if err != nil {
		return errors.ErrCache.WithArgs("add", err)
	}
	return nil
}

// Get gets a value from the cache already cast to your type.
func (sm *Redis) Get(key string, output interface{}) error {
	js, err := sm.client.Get(sm.ctx, key).Result()
	if err != nil {
		return errors.ErrCache.WithArgs("get", err)
	}
	err = json.Unmarshal([]byte(js), output)
	if err != nil {
		return errors.ErrCache.WithArgs("get", err)
	}
	return nil
}

// Del deletes a key from the backend.
func (sm *Redis) Del(key string) error {
	err := sm.client.Del(sm.ctx, key).Err()
	if err != nil {
		return errors.ErrCache.WithArgs("del", err)
	}
	return nil
}

// Exists checks if the key exists in the backend.
func (sm *Redis) Exists(key string) (bool, error) {
	result, err := sm.client.Exists(sm.ctx, key).Result()
	if err != nil {
		return false, errors.ErrCache.WithArgs("exists", err)
	}
	return result > 0, nil
}

// Init makes sure the redis connection is valid.
func (sm *Redis) Init() error {
	err := sm.client.Ping(sm.ctx).Err()
	if err != nil {
		return errors.ErrCache.WithArgs("init", err)
	}
	return nil
}

func (sm *Redis) String() string {
	return string(RedisBackend)
}
