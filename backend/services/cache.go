package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/opencti-lite/backend/database"
)

var ctx = context.Background()

func Set(key string, value interface{}, ttl time.Duration) {
	if database.RDB == nil {
		return
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		log.Printf("Cache marshalling error: %v", err)
		return
	}

	err = database.RDB.Set(ctx, key, serialized, ttl).Err()
	if err != nil {
		log.Printf("Cache write error: %v", err)
	}
}

func Get(key string) (string, error) {
	if database.RDB == nil {
		return "", sqlErr()
	}
	return database.RDB.Get(ctx, key).Result()
}

func Delete(key string) {
	if database.RDB == nil {
		return
	}
	database.RDB.Del(ctx, key)
}

func InvalidatePattern(pattern string) {
	if database.RDB == nil {
		return
	}

	var cursor uint64
	for {
		keys, nextCursor, err := database.RDB.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			log.Printf("Cache SCAN error: %v", err)
			return
		}

		if len(keys) > 0 {
			err = database.RDB.Del(ctx, keys...).Err()
			if err != nil {
				log.Printf("Cache DEL error: %v", err)
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
}

func sqlErr() error {
	return context.DeadlineExceeded // Just an error to signify redis is nil/unreachable
}
