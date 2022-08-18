package vo

type MsgVoV1 struct {
	// open_id | user_id | union_id | email | chat_id
	ReceiveIdTyp string `json:"receive_id_type,omitempty"`
	// text | post | image | file | audio | media | sticker | interactive | share_chat | share_user
	MsgType string `json:"msg_type"`

	Content interface{} `json:"content,omitempty"`
}

type MsgContentV1Text struct {
	Text string `json:"text"`
}

type MsgPostContentAtV1 struct {
	Tag string `json:"tag"`
	// open_id | union_id | user_id
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
