package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getMessage(c echo.Context) error {
	userID := sessUserID(c)
	if userID == 0 {
		return c.NoContent(http.StatusForbidden)
	}

	chanID, err := strconv.ParseInt(c.QueryParam("channel_id"), 10, 64)
	if err != nil {
		return err
	}
	lastID, err := strconv.ParseInt(c.QueryParam("last_message_id"), 10, 64)
	if err != nil {
		return err
	}

	messages, dict, err := queryMessagesWithUserDict(chanID, lastID)
	if err != nil {
		fmt.Println("queryMessagesWithUserDict: err", err)
		return err
	}

	response := make([]map[string]interface{}, 0)
	for i := len(messages) - 1; i >= 0; i-- {
		m := messages[i]
		r := make(map[string]interface{})
		r["id"] = m.ID
		r["user"], _ = dict[m.UserID]
		r["date"] = m.CreatedAt.Format("2006/01/02 15:04:05")
		r["content"] = m.Content
		response = append(response, r)
	}

	if len(messages) > 0 {
		_, err := db.Exec("INSERT INTO haveread (user_id, channel_id, message_id, updated_at, created_at)"+
			" VALUES (?, ?, ?, NOW(), NOW())"+
			" ON DUPLICATE KEY UPDATE message_id = ?, updated_at = NOW()",
			userID, chanID, messages[0].ID, messages[0].ID)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, response)
}

func queryMessagesWithUser(chanID, lastID int64) ([]MessageUser, error) {
	msgs := []MessageUser{}
	err := db.Select(&msgs, "SELECT m.*, u.name, u.display_name, u.avatar_icon FROM message m INNER JOIN user u ON m.user_id = u.id WHERE m.id > ? AND m.channel_id = ? ORDER BY m.id DESC LIMIT 100",
		lastID, chanID)
	return msgs, err
}

func queryMessagesWithUserDict(chanID, lastID int64) ([]Message, map[int64]*User, error) {
	var msgs []Message
	err := db.Select(&msgs, "SELECT * FROM message WHERE id > ? AND channel_id = ? ORDER BY id DESC LIMIT 100",
		lastID, chanID)
	if err != nil {
		return nil, nil, err
	}
	dict, err := fetchUserDictByMessages(msgs)
	return msgs, dict, err
}
