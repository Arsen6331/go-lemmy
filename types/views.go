package types

type PersonViewSafe struct {
	Person PersonSafe       `json:"person,omitempty" url:"person,omitempty"`
	Counts PersonAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}

type PersonMentionView struct {
	PersonMention              PersonMention     `json:"person_mention,omitempty" url:"person_mention,omitempty"`
	Comment                    Comment           `json:"comment,omitempty" url:"comment,omitempty"`
	Creator                    PersonSafe        `json:"creator,omitempty" url:"creator,omitempty"`
	Post                       Post              `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe              CommunitySafe     `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Recipient                  PersonSafe        `json:"recipient,omitempty" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Subscribed                 SubscribedType    `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved,omitempty" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int]     `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}

type LocalUserSettingsView struct {
	LocalUserSettings LocalUserSettings `json:"local_user_settings,omitempty" url:"local_user_settings,omitempty"`
	Person            PersonSafe        `json:"person,omitempty" url:"person,omitempty"`
	Counts            PersonAggregates  `json:"counts,omitempty" url:"counts,omitempty"`
}

type SiteView struct {
	Site               Site                `json:"site,omitempty" url:"site,omitempty"`
	LocalSite          LocalSite           `json:"local_site,omitempty" url:"local_site,omitempty"`
	LocalSiteRateLimit LocalSiteRateLimit  `json:"local_site_rate_limit,omitempty" url:"local_site_rate_limit,omitempty"`
	Taglines           Optional[[]Tagline] `json:"taglines,omitempty" url:"taglines,omitempty"`
	Counts             SiteAggregates      `json:"counts,omitempty" url:"counts,omitempty"`
}

type PrivateMessageView struct {
	PrivateMessage PrivateMessage `json:"private_message,omitempty" url:"private_message,omitempty"`
	Creator        PersonSafe     `json:"creator,omitempty" url:"creator,omitempty"`
	Recipient      PersonSafe     `json:"recipient,omitempty" url:"recipient,omitempty"`
}

type PostView struct {
	Post                       Post           `json:"post,omitempty" url:"post,omitempty"`
	Creator                    PersonSafe     `json:"creator,omitempty" url:"creator,omitempty"`
	Community                  CommunitySafe  `json:"community,omitempty" url:"community,omitempty"`
	CreatorBannedFromCommunity bool           `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Counts                     PostAggregates `json:"counts,omitempty" url:"counts,omitempty"`
	Subscribed                 bool           `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool           `json:"saved,omitempty" url:"saved,omitempty"`
	Read                       bool           `json:"read,omitempty" url:"read,omitempty"`
	CreatorBlocked             bool           `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int]  `json:"my_vote,omitempty" url:"my_vote,omitempty"`
	UnreadComments             int            `json:"unread_comments,omitempty" url:"unread_comments,omitempty"`
}

type PostReportView struct {
	PostReport                 PostReport           `json:"post_report,omitempty" url:"post_report,omitempty"`
	Post                       Post                 `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe        `json:"community,omitempty" url:"community,omitempty"`
	Creator                    PersonSafe           `json:"creator,omitempty" url:"creator,omitempty"`
	PostCreator                PersonSafe           `json:"post_creator,omitempty" url:"post_creator,omitempty"`
	CreatorBannedFromCommunity bool                 `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int]        `json:"my_vote,omitempty" url:"my_vote,omitempty"`
	Counts                     PostAggregates       `json:"counts,omitempty" url:"counts,omitempty"`
	Resolver                   Optional[PersonSafe] `json:"resolver,omitempty" url:"resolver,omitempty"`
}

type CommentView struct {
	Comment                    Comment           `json:"comment,omitempty" url:"comment,omitempty"`
	Creator                    PersonSafe        `json:"creator,omitempty" url:"creator,omitempty"`
	Post                       Post              `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe     `json:"community,omitempty" url:"community,omitempty"`
	Counts                     CommentAggregates `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Subscribed                 SubscribedType    `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved,omitempty" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int]     `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}

type CommentReportView struct {
	CommentReport              CommentReport        `json:"comment_report,omitempty" url:"comment_report,omitempty"`
	Comment                    Comment              `json:"comment,omitempty" url:"comment,omitempty"`
	Post                       Post                 `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe              CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Creator                    PersonSafe           `json:"creator,omitempty" url:"creator,omitempty"`
	CommentCreator             PersonSafe           `json:"comment_creator,omitempty" url:"comment_creator,omitempty"`
	Counts                     CommentAggregates    `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool                 `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	MyVote                     Optional[int]        `json:"my_vote,omitempty" url:"my_vote,omitempty"`
	Resolver                   Optional[PersonSafe] `json:"resolver,omitempty" url:"resolver,omitempty"`
}

type ModAddCommunityView struct {
	ModAddCommunity ModAddCommunity      `json:"mod_add_community,omitempty" url:"mod_add_community,omitempty"`
	Moderator       Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	CommunitySafe   CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	ModdedPerson    PersonSafe           `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}

type ModTransferCommunityView struct {
	ModTransferCommunity ModTransferCommunity `json:"mod_transfer_community,omitempty" url:"mod_transfer_community,omitempty"`
	Moderator            Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	CommunitySafe        CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	ModdedPerson         PersonSafe           `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}

type ModAddView struct {
	ModAdd       ModAdd               `json:"mod_add,omitempty" url:"mod_add,omitempty"`
	Moderator    Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	ModdedPerson PersonSafe           `json:"modded_person,omitempty" url:"modded_person,omitempty"`
}

type ModBanFromCommunityView struct {
	ModBanFromCommunity ModBanFromCommunity  `json:"mod_ban_from_community,omitempty" url:"mod_ban_from_community,omitempty"`
	Moderator           Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	CommunitySafe       CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	BannedPerson        PersonSafe           `json:"banned_person,omitempty" url:"banned_person,omitempty"`
}

type ModBanView struct {
	ModBan       ModBan               `json:"mod_ban,omitempty" url:"mod_ban,omitempty"`
	Moderator    Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	BannedPerson PersonSafe           `json:"banned_person,omitempty" url:"banned_person,omitempty"`
}

type ModLockPostView struct {
	ModLockPost   ModLockPost          `json:"mod_lock_post,omitempty" url:"mod_lock_post,omitempty"`
	Moderator     Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post          Post                 `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type ModRemoveCommentView struct {
	ModRemoveComment ModRemoveComment     `json:"mod_remove_comment,omitempty" url:"mod_remove_comment,omitempty"`
	Moderator        Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	Comment          Comment              `json:"comment,omitempty" url:"comment,omitempty"`
	Commenter        PersonSafe           `json:"commenter,omitempty" url:"commenter,omitempty"`
	Post             Post                 `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe    CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type ModRemoveCommunityView struct {
	ModRemoveCommunity ModRemoveCommunity   `json:"mod_remove_community,omitempty" url:"mod_remove_community,omitempty"`
	Moderator          Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	CommunitySafe      CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type ModRemovePostView struct {
	ModRemovePost ModRemovePost        `json:"mod_remove_post,omitempty" url:"mod_remove_post,omitempty"`
	Moderator     Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post          Post                 `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type ModStickyPostView struct {
	ModStickyPost ModStickyPost        `json:"mod_sticky_post,omitempty" url:"mod_sticky_post,omitempty"`
	Moderator     Optional[PersonSafe] `json:"moderator,omitempty" url:"moderator,omitempty"`
	Post          Post                 `json:"post,omitempty" url:"post,omitempty"`
	CommunitySafe CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type AdminPurgeCommunityView struct {
	AdminPurgeCommunity AdminPurgeCommunity  `json:"admin_purge_community,omitempty" url:"admin_purge_community,omitempty"`
	Admin               Optional[PersonSafe] `json:"admin,omitempty" url:"admin,omitempty"`
}

type AdminPurgePersonView struct {
	AdminPurgePerson AdminPurgePerson     `json:"admin_purge_person,omitempty" url:"admin_purge_person,omitempty"`
	Admin            Optional[PersonSafe] `json:"admin,omitempty" url:"admin,omitempty"`
}

type AdminPurgePostView struct {
	AdminPurgePost AdminPurgePost       `json:"admin_purge_post,omitempty" url:"admin_purge_post,omitempty"`
	Admin          Optional[PersonSafe] `json:"admin,omitempty" url:"admin,omitempty"`
	CommunitySafe  CommunitySafe        `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type AdminPurgeCommentView struct {
	AdminPurgeComment AdminPurgeComment    `json:"admin_purge_comment,omitempty" url:"admin_purge_comment,omitempty"`
	Admin             Optional[PersonSafe] `json:"admin,omitempty" url:"admin,omitempty"`
	Post              Post                 `json:"post,omitempty" url:"post,omitempty"`
}

type CommunityFollowerView struct {
	CommunitySafe CommunitySafe `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Follower      PersonSafe    `json:"follower,omitempty" url:"follower,omitempty"`
}

type CommunityBlockView struct {
	PersonSafe    PersonSafe    `json:"person_safe,omitempty" url:"person_safe,omitempty"`
	CommunitySafe CommunitySafe `json:"community_safe,omitempty" url:"community_safe,omitempty"`
}

type CommunityModeratorView struct {
	CommunitySafe CommunitySafe `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Moderator     PersonSafe    `json:"moderator,omitempty" url:"moderator,omitempty"`
}

type CommunityPersonBanView struct {
	CommunitySafe CommunitySafe `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Person        PersonSafe    `json:"person,omitempty" url:"person,omitempty"`
}

type PersonBlockView struct {
	PersonSafe PersonSafe `json:"person_safe,omitempty" url:"person_safe,omitempty"`
	Target     PersonSafe `json:"target,omitempty" url:"target,omitempty"`
}

type CommunityView struct {
	CommunitySafe CommunitySafe       `json:"community_safe,omitempty" url:"community_safe,omitempty"`
	Subscribed    SubscribedType      `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Blocked       bool                `json:"blocked,omitempty" url:"blocked,omitempty"`
	Counts        CommunityAggregates `json:"counts,omitempty" url:"counts,omitempty"`
}

type RegistrationApplicationView struct {
	RegistrationApplication RegistrationApplication `json:"registration_application,omitempty" url:"registration_application,omitempty"`
	CreatorLocalUser        LocalUserSettings       `json:"creator_local_user,omitempty" url:"creator_local_user,omitempty"`
	Creator                 PersonSafe              `json:"creator,omitempty" url:"creator,omitempty"`
	Admin                   Optional[PersonSafe]    `json:"admin,omitempty" url:"admin,omitempty"`
}

type PrivateMessageReportView struct {
	PrivateMessageReport  PrivateMessageReport `json:"private_message_report,omitempty" url:"private_message_report,omitempty"`
	PrivateMessage        PrivateMessage       `json:"private_message,omitempty" url:"private_message,omitempty"`
	PrivateMessageCreator PersonSafe           `json:"private_message_creator,omitempty" url:"private_message_creator,omitempty"`
	Creator               PersonSafe           `json:"creator,omitempty" url:"creator,omitempty"`
	Resolver              Optional[PersonSafe] `json:"resolver,omitempty" url:"resolver,omitempty"`
}

type Tagline struct {
	ID          int              `json:"id,omitempty" url:"id,omitempty"`
	LocalSiteID int              `json:"local_site_id,omitempty" url:"local_site_id,omitempty"`
	Content     string           `json:"content,omitempty" url:"content,omitempty"`
	Published   string           `json:"published,omitempty" url:"published,omitempty"`
	Updated     Optional[string] `json:"updated,omitempty" url:"updated,omitempty"`
}

type CommentReplyView struct {
	CommentReply               CommentReply      `json:"comment_reply,omitempty" url:"comment_reply,omitempty"`
	Comment                    Comment           `json:"comment,omitempty" url:"comment,omitempty"`
	Creator                    PersonSafe        `json:"creator,omitempty" url:"creator,omitempty"`
	Post                       Post              `json:"post,omitempty" url:"post,omitempty"`
	Community                  CommunitySafe     `json:"community,omitempty" url:"community,omitempty"`
	Recipient                  PersonSafe        `json:"recipient,omitempty" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts,omitempty" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community,omitempty" url:"creator_banned_from_community,omitempty"`
	Subscribed                 SubscribedType    `json:"subscribed,omitempty" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved,omitempty" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked,omitempty" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int]     `json:"my_vote,omitempty" url:"my_vote,omitempty"`
}