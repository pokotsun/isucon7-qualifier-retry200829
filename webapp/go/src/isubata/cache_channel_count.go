package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func getChannelCountKey(channelID int64) string {
	return fmt.Sprintf("CHANNEL_COUNT-channelID-%d", channelID)
}

func multiSetChannelCount(channelCounts []ChannelCount) {
	cmap := make(map[string][]byte)

	for _, chCount := range channelCounts {
		key := getChannelCountKey(chCount.ChannelID)
		v, err := json.Marshal(chCount.Count)
		if err != nil {
			sugar.Errorf("jsonMarshal Err: %s", err)
		}
		cmap[key] = v
	}
	err := cacheClient.MultiSet(cmap)
	if err != nil {
		sugar.Errorf("cache MSet ChannelCount: %s", err)
	}
}

func setChannelCount(channelID, count int64) {
	key := getChannelCountKey(channelID)
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(count))
	err := cacheClient.SingleSet(key, value)
	if err != nil {
		sugar.Errorf("Failed to Set Channel Count: %s", err)
	}
}

func incrementChannelCount(channelID int64) {
	key := getChannelCountKey(channelID)
	_, err := cacheClient.Increment(key, 1)
	if err != nil {
		sugar.Errorf("Channel Count Increment Err: %s", err)
	}
}

func getChannelCount(channelID int64) (int64, error) {
	key := getChannelCountKey(channelID)
	bytes, err := cacheClient.SingleGet(key)
	if err != nil {
		sugar.Errorf("Failed to Get Cache Channel Count: %s", err)
		return 0, nil
	}
	return int64(binary.LittleEndian.Uint64(bytes)), nil
}
