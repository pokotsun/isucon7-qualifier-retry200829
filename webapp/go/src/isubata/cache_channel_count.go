package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

func getChannelCountKey(channelID int64) string {
	return fmt.Sprintf("CHANNEL_COUNT-channelID-%d", channelID)
}

func setChannelCount(client *redisClient, channelID, count int64) {
	key := getChannelCountKey(channelID)
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(count))
	err := client.SingleSet(key, value)
	if err != nil {
		log.Printf("Failed to Cache Channel Count: %s", err)
	}
}

func getChannelCount(client *redisClient, channelID int64) (int64, error) {
	key := getChannelCountKey(channelID)
	bytes, err := client.SingleGet(key)
	if err != nil {
		log.Printf("Failed to Get Cache Of Channel Count: %s", err)
		return 0, nil
	}
	return int64(binary.LittleEndian.Uint64(bytes)), nil
}
