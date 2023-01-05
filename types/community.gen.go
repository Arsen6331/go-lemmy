package types

type GetCommunity struct {
	ID   Optional[int]    `json:"id" url:"id,omitempty"`
	Name Optional[string] `json:"name" url:"name,omitempty"`
	Auth Optional[string] `json:"auth" url:"auth,omitempty"`
}
type GetCommunityResponse struct {
	CommunityView CommunityView            `json:"community_view" url:"community_view,omitempty"`
	Site          Optional[Site]           `json:"site" url:"site,omitempty"`
	Moderators    []CommunityModeratorView `json:"moderators" url:"moderators,omitempty"`
	Online        uint                     `json:"online" url:"online,omitempty"`
	LemmyResponse
}
type CreateCommunity struct {
	Name                    string           `json:"name" url:"name,omitempty"`
	Title                   string           `json:"title" url:"title,omitempty"`
	Description             Optional[string] `json:"description" url:"description,omitempty"`
	Icon                    Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner" url:"banner,omitempty"`
	NSFW                    Optional[bool]   `json:"nsfw" url:"nsfw,omitempty"`
	PostingRestrictedToMods Optional[bool]   `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
	Auth                    string           `json:"auth" url:"auth,omitempty"`
}
type CommunityResponse struct {
	CommunityView CommunityView `json:"community_view" url:"community_view,omitempty"`
	LemmyResponse
}
type ListCommunities struct {
	Type  Optional[ListingType] `json:"type_" url:"type_,omitempty"`
	Sort  Optional[SortType]    `json:"sort" url:"sort,omitempty"`
	Page  Optional[int64]       `json:"page" url:"page,omitempty"`
	Limit Optional[int64]       `json:"limit" url:"limit,omitempty"`
	Auth  Optional[string]      `json:"auth" url:"auth,omitempty"`
}
type ListCommunitiesResponse struct {
	Communities []CommunityView `json:"communities" url:"communities,omitempty"`
	LemmyResponse
}
type BanFromCommunity struct {
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	PersonID    int              `json:"person_id" url:"person_id,omitempty"`
	Ban         bool             `json:"ban" url:"ban,omitempty"`
	RemoveData  Optional[bool]   `json:"remove_data" url:"remove_data,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Expires     Optional[int64]  `json:"expires" url:"expires,omitempty"`
	Auth        string           `json:"auth" url:"auth,omitempty"`
}
type BanFromCommunityResponse struct {
	PersonView PersonViewSafe `json:"person_view" url:"person_view,omitempty"`
	Banned     bool           `json:"banned" url:"banned,omitempty"`
	LemmyResponse
}
type AddModToCommunity struct {
	CommunityID int    `json:"community_id" url:"community_id,omitempty"`
	PersonID    int    `json:"person_id" url:"person_id,omitempty"`
	Added       bool   `json:"added" url:"added,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type AddModToCommunityResponse struct {
	Moderators []CommunityModeratorView `json:"moderators" url:"moderators,omitempty"`
	LemmyResponse
}
type EditCommunity struct {
	CommunityID             int              `json:"community_id" url:"community_id,omitempty"`
	Title                   Optional[string] `json:"title" url:"title,omitempty"`
	Description             Optional[string] `json:"description" url:"description,omitempty"`
	Icon                    Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                  Optional[string] `json:"banner" url:"banner,omitempty"`
	NSFW                    Optional[bool]   `json:"nsfw" url:"nsfw,omitempty"`
	PostingRestrictedToMods Optional[bool]   `json:"posting_restricted_to_mods" url:"posting_restricted_to_mods,omitempty"`
	Auth                    string           `json:"auth" url:"auth,omitempty"`
}
type HideCommunity struct {
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Hidden      bool             `json:"hidden" url:"hidden,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Auth        string           `json:"auth" url:"auth,omitempty"`
}
type DeleteCommunity struct {
	CommunityID int    `json:"community_id" url:"community_id,omitempty"`
	Deleted     bool   `json:"deleted" url:"deleted,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type RemoveCommunity struct {
	CommunityID int              `json:"community_id" url:"community_id,omitempty"`
	Removed     bool             `json:"removed" url:"removed,omitempty"`
	Reason      Optional[string] `json:"reason" url:"reason,omitempty"`
	Expires     Optional[int64]  `json:"expires" url:"expires,omitempty"`
	Auth        string           `json:"auth" url:"auth,omitempty"`
}
type FollowCommunity struct {
	CommunityID int    `json:"community_id" url:"community_id,omitempty"`
	Follow      bool   `json:"follow" url:"follow,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type BlockCommunity struct {
	CommunityID int    `json:"community_id" url:"community_id,omitempty"`
	Block       bool   `json:"block" url:"block,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type BlockCommunityResponse struct {
	CommunityView CommunityView `json:"community_view" url:"community_view,omitempty"`
	Blocked       bool          `json:"blocked" url:"blocked,omitempty"`
	LemmyResponse
}
type TransferCommunity struct {
	CommunityID int    `json:"community_id" url:"community_id,omitempty"`
	PersonID    int    `json:"person_id" url:"person_id,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}