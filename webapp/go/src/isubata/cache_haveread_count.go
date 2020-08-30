package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

func getHavereadCountKey(userID, channelID int64) string {
	return fmt.Sprintf("HAVEREAD_COUNT-userID-%d-channelID-%d", userID, channelID)
}

func setHavereadCount(client *redisClient, userID, channelID, count int64) {
	key := getHavereadCountKey(userID, channelID)
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(count))
	err := client.SingleSet(key, value)
	if err != nil {
		log.Printf("Failed to Cache Haveread Count: %s", err)
	}
}

func getHaveReadCount(client *redisClient, userID, channelID int64) (int64, error) {
	key := getHavereadCountKey(userID, channelID)
	bytes, err := client.SingleGet(key)
	if err != nil {
		log.Printf("Failed to Get Cache Of Haveread Count: %s", err)
		return 0, nil
	}
	return int64(binary.LittleEndian.Uint64(bytes)), nil
}
