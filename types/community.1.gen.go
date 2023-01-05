package types

type Community struct {
	ID                      int              `json:"id" url:"id,omitempty"`
	Name                    string           `json:"name" url:"name,omitempty"`
	Title                   string           `json:"title" url:"title,omitempty"`
	Description             Optional[string] `json:"description" url:"description,omitempty"`
	Removed                 bool             `json:"removed" url:"removed,omitempty"`
	Published               LemmyTime        `json:"published" url:"published,omitempty"`
	Updated                 LemmyTime        `json:"updated" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id" url:"actor_id,omitempty"`
	Local                   bool             `json:"local" url:"local,omitempty"`
	PrivateKey              Optional[string] `json:"private_key" url:"private_key,omitempty"`
	PublicKey               string           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt         LemmyTime        `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Icon                    Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner" url:"banner,omitempty"`
	FollowersURL            string           `json:"followers_url" url:"followers_url,omitempty"`
	InboxURL                string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL          Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	Hidden                  bool             `json:"hidden" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
}
type CommunitySafe struct {
	ID                      int              `json:"id" url:"id,omitempty"`
	Name                    string           `json:"name" url:"name,omitempty"`
	Title                   string           `json:"title" url:"title,omitempty"`
	Description             Optional[string] `json:"description" url:"description,omitempty"`
	Removed                 bool             `json:"removed" url:"removed,omitempty"`
	Published               LemmyTime        `json:"published" url:"published,omitempty"`
	Updated                 LemmyTime        `json:"updated" url:"updated,omitempty"`
	Deleted                 bool             `json:"deleted" url:"deleted,omitempty"`
	NSFW                    bool             `json:"nsfw" url:"nsfw,omitempty"`
	ActorID                 string           `json:"actor_id" url:"actor_id,omitempty"`
	Local                   bool             `json:"local" url:"local,omitempty"`
	Icon                    Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner" url:"banner,omitempty"`
	Hidden                  bool             `json:"hidden" url:"hidden,omitempty"`
	PostingRestrictedToMods bool             `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
}
type CommunityForm struct {
	Name                    string                     `json:"name" url:"name,omitempty"`
	Title                   string                     `json:"title" url:"title,omitempty"`
	Description             Optional[string]           `json:"description" url:"description,omitempty"`
	Removed                 Optional[bool]             `json:"removed" url:"removed,omitempty"`
	Published               LemmyTime                  `json:"published" url:"published,omitempty"`
	Updated                 LemmyTime                  `json:"updated" url:"updated,omitempty"`
	Deleted                 Optional[bool]             `json:"deleted" url:"deleted,omitempty"`
	NSFW                    Optional[bool]             `json:"nsfw" url:"nsfw,omitempty"`
	ActorID                 Optional[string]           `json:"actor_id" url:"actor_id,omitempty"`
	Local                   Optional[bool]             `json:"local" url:"local,omitempty"`
	PrivateKey              Optional[Optional[string]] `json:"private_key" url:"private_key,omitempty"`
	PublicKey               Optional[string]           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt         LemmyTime                  `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Icon                    Optional[Optional[string]] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[Optional[string]] `json:"banner" url:"banner,omitempty"`
	FollowersURL            Optional[string]           `json:"followers_url" url:"followers_url,omitempty"`
	InboxURL                Optional[string]           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL          Optional[Optional[string]] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	Hidden                  Optional[bool]             `json:"hidden" url:"hidden,omitempty"`
	PostingRestrictedToMods Optional[bool]             `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
}
type CommunityModerator struct {
	ID          int32     `json:"id" url:"id,omitempty"`
	CommunityID int       `json:"community_id" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id" url:"person_id,omitempty"`
	Published   LemmyTime `json:"published" url:"published,omitempty"`
}
type CommunityModeratorForm struct {
	CommunityID int `json:"community_id" url:"community_id,omitempty"`
	PersonID    int `json:"person_id" url:"person_id,omitempty"`
}
type CommunityPersonBan struct {
	ID          int32     `json:"id" url:"id,omitempty"`
	CommunityID int       `json:"community_id" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id" url:"person_id,omitempty"`
	Published   LemmyTime `json:"published" url:"published,omitempty"`
	Expires     LemmyTime `json:"expires" url:"expires,omitempty"`
}
type CommunityPersonBanForm struct {
	CommunityID int       `json:"community_id" url:"community_id,omitempty"`
	PersonID    int       `json:"person_id" url:"person_id,omitempty"`
	Expires     LemmyTime `json:"expires" url:"expires,omitempty"`
}
type CommunityFollower struct {
	ID          int32          `json:"id" url:"id,omitempty"`
	CommunityID int            `json:"community_id" url:"community_id,omitempty"`
	PersonID    int            `json:"person_id" url:"person_id,omitempty"`
	Published   LemmyTime      `json:"published" url:"published,omitempty"`
	Pending     Optional[bool] `json:"pending" url:"pending,omitempty"`
}
type CommunityFollowerForm struct {
	CommunityID int  `json:"community_id" url:"community_id,omitempty"`
	PersonID    int  `json:"person_id" url:"person_id,omitempty"`
	Pending     bool `json:"pending" url:"pending,omitempty"`
}