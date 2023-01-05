package types

type PostReport struct {
	ID               int              `json:"id" url:"id,omitempty"`
	CreatorID        int              `json:"creator_id" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason" url:"reason,omitempty"`
	Resolved         bool             `json:"resolved" url:"resolved,omitempty"`
	ResolverID       Optional[int]    `json:"resolver_id" url:"resolver_id,omitempty"`
	Published        LemmyTime        `json:"published" url:"published,omitempty"`
	Updated          LemmyTime        `json:"updated" url:"updated,omitempty"`
}
type PostReportForm struct {
	CreatorID        int              `json:"creator_id" url:"creator_id,omitempty"`
	PostID           int              `json:"post_id" url:"post_id,omitempty"`
	OriginalPostName string           `json:"original_post_name" url:"original_post_name,omitempty"`
	OriginalPostURL  Optional[string] `json:"original_post_url" url:"original_post_url,omitempty"`
	OriginalPostBody Optional[string] `json:"original_post_body" url:"original_post_body,omitempty"`
	Reason           string           `json:"reason" url:"reason,omitempty"`
}