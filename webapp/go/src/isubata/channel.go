package main

func fetchChannelCountDict() (map[int64]int64, error) {
	counts := []ChannelCount{}
	err := db.Select(&counts, "SELECT channel_id, COUNT(*) as cnt FROM message GROUP BY channel_id")
	if err != nil {
		return nil, err
	}
	res := map[int64]int64{}
	for _, v := range counts {
		res[v.ChannelID] = v.Count
	}
	return res, nil
}
