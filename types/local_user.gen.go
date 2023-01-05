package types

type LocalUser struct {
	ID                       int              `json:"id" url:"id,omitempty"`
	PersonID                 int              `json:"person_id" url:"person_id,omitempty"`
	PasswordEncrypted        string           `json:"password_encrypted" url:"password_encrypted,omitempty"`
	Email                    Optional[string] `json:"email" url:"email,omitempty"`
	ShowNSFW                 bool             `json:"show_nsfw" url:"show_nsfw,omitempty"`
	Theme                    string           `json:"theme" url:"theme,omitempty"`
	DefaultSortType          int16            `json:"default_sort_type" url:"default_sort_type,omitempty"`
	DefaultListingType       int16            `json:"default_listing_type" url:"default_listing_type,omitempty"`
	Lang                     string           `json:"lang" url:"lang,omitempty"`
	ShowAvatars              bool             `json:"show_avatars" url:"show_avatars,omitempty"`
	SendNotificationsToEmail bool             `json:"send_notifications_to_email" url:"send_notifications_to_email,omitempty"`
	ValidatorTime            LemmyTime        `json:"validator_time" url:"validator_time,omitempty"`
	ShowBotAccounts          bool             `json:"show_bot_accounts" url:"show_bot_accounts,omitempty"`
	ShowScores               bool             `json:"show_scores" url:"show_scores,omitempty"`
	ShowReadPosts            bool             `json:"show_read_posts" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs        bool             `json:"show_new_post_notifs" url:"show_new_post_notifs,omitempty"`
	EmailVerified            bool             `json:"email_verified" url:"email_verified,omitempty"`
	AcceptedApplication      bool             `json:"accepted_application" url:"accepted_application,omitempty"`
}
type LocalUserForm struct {
	PersonID                 Optional[int]              `json:"person_id" url:"person_id,omitempty"`
	PasswordEncrypted        Optional[string]           `json:"password_encrypted" url:"password_encrypted,omitempty"`
	Email                    Optional[Optional[string]] `json:"email" url:"email,omitempty"`
	ShowNSFW                 Optional[bool]             `json:"show_nsfw" url:"show_nsfw,omitempty"`
	Theme                    Optional[string]           `json:"theme" url:"theme,omitempty"`
	DefaultSortType          Optional[int16]            `json:"default_sort_type" url:"default_sort_type,omitempty"`
	DefaultListingType       Optional[int16]            `json:"default_listing_type" url:"default_listing_type,omitempty"`
	Lang                     Optional[string]           `json:"lang" url:"lang,omitempty"`
	ShowAvatars              Optional[bool]             `json:"show_avatars" url:"show_avatars,omitempty"`
	SendNotificationsToEmail Optional[bool]             `json:"send_notifications_to_email" url:"send_notifications_to_email,omitempty"`
	ShowBotAccounts          Optional[bool]             `json:"show_bot_accounts" url:"show_bot_accounts,omitempty"`
	ShowScores               Optional[bool]             `json:"show_scores" url:"show_scores,omitempty"`
	ShowReadPosts            Optional[bool]             `json:"show_read_posts" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs        Optional[bool]             `json:"show_new_post_notifs" url:"show_new_post_notifs,omitempty"`
	EmailVerified            Optional[bool]             `json:"email_verified" url:"email_verified,omitempty"`
	AcceptedApplication      Optional[bool]             `json:"accepted_application" url:"accepted_application,omitempty"`
}
type LocalUserSettings struct {
	ID                       int              `json:"id" url:"id,omitempty"`
	PersonID                 int              `json:"person_id" url:"person_id,omitempty"`
	Email                    Optional[string] `json:"email" url:"email,omitempty"`
	ShowNSFW                 bool             `json:"show_nsfw" url:"show_nsfw,omitempty"`
	Theme                    string           `json:"theme" url:"theme,omitempty"`
	DefaultSortType          int16            `json:"default_sort_type" url:"default_sort_type,omitempty"`
	DefaultListingType       int16            `json:"default_listing_type" url:"default_listing_type,omitempty"`
	Lang                     string           `json:"lang" url:"lang,omitempty"`
	ShowAvatars              bool             `json:"show_avatars" url:"show_avatars,omitempty"`
	SendNotificationsToEmail bool             `json:"send_notifications_to_email" url:"send_notifications_to_email,omitempty"`
	ValidatorTime            LemmyTime        `json:"validator_time" url:"validator_time,omitempty"`
	ShowBotAccounts          bool             `json:"show_bot_accounts" url:"show_bot_accounts,omitempty"`
	ShowScores               bool             `json:"show_scores" url:"show_scores,omitempty"`
	ShowReadPosts            bool             `json:"show_read_posts" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs        bool             `json:"show_new_post_notifs" url:"show_new_post_notifs,omitempty"`
	EmailVerified            bool             `json:"email_verified" url:"email_verified,omitempty"`
	AcceptedApplication      bool             `json:"accepted_application" url:"accepted_application,omitempty"`
}
