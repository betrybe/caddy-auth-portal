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
	"fmt"
	"github.com/greenpau/caddy-auth-portal/pkg/errors"
)

// Backend cache
type Backend string

const (
	expirationTime = 5 * 60

	// MemoryBackend is a memory cache backend
	MemoryBackend Backend = "memory"
	// MemcachedBackend is a memcache cache backend
	MemcachedBackend Backend = "memcached"
	// RedisBackend is a redis cache backend
	RedisBackend Backend = "redis"
)

var (
	validCacheTypes = map[Backend]bool{
		MemoryBackend:    true,
		MemcachedBackend: true,
		RedisBackend:     true,
	}
)

// Cache stores the state of the OAuth2 flow.
type Cache interface {
	Init() error
	Add(key string, data interface{}) error
	Get(key string, output interface{}) error
	Del(key string) error
	Exists(key string) (bool, error)
}

// Config stores a simple config to initialize the backend once again
type Config struct {
	Backend Backend
	Config  []string
}

// Validate validates the config
func Validate(name Backend) error {
	_, ok := validCacheTypes[name]
	if !ok {
		return errors.ErrCacheBackendNotFound.WithArgs(name)
	}
	return nil
}

// RequiresParameters returns true if the backend requires config parameters
func RequiresParameters(name Backend) bool {
	switch name {
	case MemoryBackend:
		return false
	case MemcachedBackend:
		return true
	case RedisBackend:
		return true
	default:
		panic(fmt.Sprintf("invalid cache type supplied %s", name))
	}
}

// NewFromArgs creates a new cache backend from the supplied arguments
func NewFromArgs(config *Config) Cache {
	switch config.Backend {
	case MemoryBackend:
		return NewMemoryCache()
	case MemcachedBackend:
		return NewMemcachedCache(config.Config...)
	case RedisBackend:
		return NewRedisCache(config.Config...)
	default:
		panic(fmt.Sprintf("invalid cache provided %s", config.Backend))
	}
}
