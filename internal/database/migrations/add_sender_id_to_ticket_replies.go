package migrations

import (
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/models"
	"gorm.io/gorm"
)

func AddSenderIdToTicketReplies(db *gorm.DB) error {
	if db.Migrator().HasColumn(&models.TicketReply{},"sender_id") {
		return nil
	}
	if err := db.Migrator().AddColumn(&models.TicketReply{},"sender_id");err != nil {
		return err
	}
	return db.Exec(`
		ALTER TABLE ticket_replies 
		ADD CONSTRAINT fk_ticket_replies_sender 
		FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
	`).Error
}