//  Source: lemmy/crates/api_common/src/websocket.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type UserJoin struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type UserJoinResponse struct {
	Joined bool `json:"joined" url:"joined,omitempty"`
	LemmyResponse
}
type CommunityJoin struct {
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
}
type CommunityJoinResponse struct {
	Joined bool `json:"joined" url:"joined,omitempty"`
	LemmyResponse
}
type ModJoin struct {
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
}
type ModJoinResponse struct {
	Joined bool `json:"joined" url:"joined,omitempty"`
	LemmyResponse
}
type PostJoin struct {
	PostID int `json:"post_id" url:"post_id,omitempty"`
}
type PostJoinResponse struct {
	Joined bool `json:"joined" url:"joined,omitempty"`
	LemmyResponse
}
