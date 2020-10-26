package model

// Message is message board content in db
type Message struct {
	MsgID       int    `db:"messageId"`
	AuthorName  string `db:"authorName"`
	AuthorEmail string `db:"authorEmail"`
	Msg         string `db:"message"`
	MsgTime     string `db:"messageTime"`
}
