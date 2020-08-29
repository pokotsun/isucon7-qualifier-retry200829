package main

import "github.com/jmoiron/sqlx"

func fetchUserDictByMessages(msgs []Message) (map[int64]*User, error) {
	if len(msgs) == 0 {
		return nil, nil
	}
	res := map[int64]*User{}
	var userIDs []int64
	for _, v := range msgs {
		userIDs = append(userIDs, v.UserID)
	}
	query := "SELECT * FROM user WHERE id IN (?)"
	inQuery, inArgs, err := sqlx.In(query, userIDs)
	if err != nil {
		return nil, err
	}
	var users []User
	err = db.Select(&users, inQuery, inArgs...)
	if err != nil {
		return nil, err
	}
	for i := range users {
		res[users[i].ID] = &users[i]
	}

	return res, nil
}
