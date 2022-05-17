package sqlite

import (
	"fmt"
	"forum/internal/models"
)

func (s *Store) InsertMessage(message models.MessageData) error {
	createRow, err := s.db.Prepare(`
		INSERT INTO chat 
		(chat_id, sender, receiver, message, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("InsertMessage, Prepare: %w", err)
	}
	defer createRow.Close()

	_, err = createRow.Exec(
		message.ID,
		message.SenderNickname,
		message.ReceiverNickname,
		message.Message,
		message.Timestamp,
	)
	if err != nil {
		return fmt.Errorf("InsertMessage, Exec: %w", err)
	}

	return nil
}

func (s *Store) GetChatByID(chatID int, senderNick string) ([]models.MessageData, error) {
	var messages []models.MessageData

	rows, err := s.db.Query(`
		SELECT * FROM chat WHERE chat_id = ? AND sender = ?
	`, chatID, senderNick)
	if err != nil {
		return messages, fmt.Errorf("GetChatByID, Query: %w", err)
	}
	defer rows.Close()

	var message models.MessageData
	for rows.Next() {
		if err := rows.Scan(&message.ID, &message.SenderNickname, &message.ReceiverNickname, &message.Message, &message.Timestamp); err != nil {
			return messages, fmt.Errorf("GetChatByID, Scan: %w", err)
		}
		messages = append(messages, message)
	}

	return messages, nil
}
