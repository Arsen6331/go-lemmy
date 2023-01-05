package types

type PersonMention struct {
	ID          int       `json:"id" url:"id,omitempty"`
	RecipientID int       `json:"recipient_id" url:"recipient_id,omitempty"`
	CommentID   int       `json:"comment_id" url:"comment_id,omitempty"`
	Read        bool      `json:"read" url:"read,omitempty"`
	Published   LemmyTime `json:"published" url:"published,omitempty"`
}
type PersonMentionForm struct {
	RecipientID int            `json:"recipient_id" url:"recipient_id,omitempty"`
	CommentID   int            `json:"comment_id" url:"comment_id,omitempty"`
	Read        Optional[bool] `json:"read" url:"read,omitempty"`
}