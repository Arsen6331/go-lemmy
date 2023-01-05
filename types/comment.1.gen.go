package types

type Comment struct {
	ID        int           `json:"id" url:"id,omitempty"`
	CreatorID int           `json:"creator_id" url:"creator_id,omitempty"`
	PostID    int           `json:"post_id" url:"post_id,omitempty"`
	ParentID  Optional[int] `json:"parent_id" url:"parent_id,omitempty"`
	Content   string        `json:"content" url:"content,omitempty"`
	Removed   bool          `json:"removed" url:"removed,omitempty"`
	Read      bool          `json:"read" url:"read,omitempty"`
	Published LemmyTime     `json:"published" url:"published,omitempty"`
	Updated   LemmyTime     `json:"updated" url:"updated,omitempty"`
	Deleted   bool          `json:"deleted" url:"deleted,omitempty"`
	ApID      string        `json:"ap_id" url:"ap_id,omitempty"`
	Local     bool          `json:"local" url:"local,omitempty"`
}
type CommentAlias1 struct {
	ID        int           `json:"id" url:"id,omitempty"`
	CreatorID int           `json:"creator_id" url:"creator_id,omitempty"`
	PostID    int           `json:"post_id" url:"post_id,omitempty"`
	ParentID  Optional[int] `json:"parent_id" url:"parent_id,omitempty"`
	Content   string        `json:"content" url:"content,omitempty"`
	Removed   bool          `json:"removed" url:"removed,omitempty"`
	Read      bool          `json:"read" url:"read,omitempty"`
	Published LemmyTime     `json:"published" url:"published,omitempty"`
	Updated   LemmyTime     `json:"updated" url:"updated,omitempty"`
	Deleted   bool          `json:"deleted" url:"deleted,omitempty"`
	ApID      string        `json:"ap_id" url:"ap_id,omitempty"`
	Local     bool          `json:"local" url:"local,omitempty"`
}
type CommentForm struct {
	CreatorID int              `json:"creator_id" url:"creator_id,omitempty"`
	PostID    int              `json:"post_id" url:"post_id,omitempty"`
	Content   string           `json:"content" url:"content,omitempty"`
	ParentID  Optional[int]    `json:"parent_id" url:"parent_id,omitempty"`
	Removed   Optional[bool]   `json:"removed" url:"removed,omitempty"`
	Read      Optional[bool]   `json:"read" url:"read,omitempty"`
	Published LemmyTime        `json:"published" url:"published,omitempty"`
	Updated   LemmyTime        `json:"updated" url:"updated,omitempty"`
	Deleted   Optional[bool]   `json:"deleted" url:"deleted,omitempty"`
	ApID      Optional[string] `json:"ap_id" url:"ap_id,omitempty"`
	Local     Optional[bool]   `json:"local" url:"local,omitempty"`
}
type CommentLike struct {
	ID        int32     `json:"id" url:"id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	CommentID int       `json:"comment_id" url:"comment_id,omitempty"`
	PostID    int       `json:"post_id" url:"post_id,omitempty"`
	Score     int16     `json:"score" url:"score,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type CommentLikeForm struct {
	PersonID  int   `json:"person_id" url:"person_id,omitempty"`
	CommentID int   `json:"comment_id" url:"comment_id,omitempty"`
	PostID    int   `json:"post_id" url:"post_id,omitempty"`
	Score     int16 `json:"score" url:"score,omitempty"`
}
type CommentSaved struct {
	ID        int32     `json:"id" url:"id,omitempty"`
	CommentID int       `json:"comment_id" url:"comment_id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type CommentSavedForm struct {
	CommentID int `json:"comment_id" url:"comment_id,omitempty"`
	PersonID  int `json:"person_id" url:"person_id,omitempty"`
}