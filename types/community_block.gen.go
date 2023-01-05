package types

type CommunityBlock struct {
	ID          int       `json:"id" url:"id,omitempty"`
	PersonID    int       `json:"person_id" url:"person_id,omitempty"`
	CommunityID int       `json:"community_id" url:"community_id,omitempty"`
	Published   LemmyTime `json:"published" url:"published,omitempty"`
}
type CommunityBlockForm struct {
	PersonID    int `json:"person_id" url:"person_id,omitempty"`
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
}