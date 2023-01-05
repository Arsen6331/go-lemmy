package types

type Login struct {
	UsernameOrEmail string `json:"username_or_email" url:"username_or_email,omitempty"`
	Password        string `json:"password" url:"password,omitempty"`
}
type Register struct {
	Username       string           `json:"username" url:"username,omitempty"`
	Password       string           `json:"password" url:"password,omitempty"`
	PasswordVerify string           `json:"password_verify" url:"password_verify,omitempty"`
	ShowNSFW       bool             `json:"show_nsfw" url:"show_nsfw,omitempty"`
	Email          Optional[string] `json:"email" url:"email,omitempty"`
	CaptchaUuid    Optional[string] `json:"captcha_uuid" url:"captcha_uuid,omitempty"`
	CaptchaAnswer  Optional[string] `json:"captcha_answer" url:"captcha_answer,omitempty"`
	Honeypot       Optional[string] `json:"honeypot" url:"honeypot,omitempty"`
	Answer         Optional[string] `json:"answer" url:"answer,omitempty"`
}
type GetCaptcha struct{}
type GetCaptchaResponse struct {
	Ok Optional[CaptchaResponse] `json:"ok" url:"ok,omitempty"`
	LemmyResponse
}
type CaptchaResponse struct {
	Png  string `json:"png" url:"png,omitempty"`
	Wav  string `json:"wav" url:"wav,omitempty"`
	Uuid string `json:"uuid" url:"uuid,omitempty"`
	LemmyResponse
}
type SaveUserSettings struct {
	ShowNSFW                 Optional[bool]   `json:"show_nsfw" url:"show_nsfw,omitempty"`
	ShowScores               Optional[bool]   `json:"show_scores" url:"show_scores,omitempty"`
	Theme                    Optional[string] `json:"theme" url:"theme,omitempty"`
	DefaultSortType          Optional[int16]  `json:"default_sort_type" url:"default_sort_type,omitempty"`
	DefaultListingType       Optional[int16]  `json:"default_listing_type" url:"default_listing_type,omitempty"`
	Lang                     Optional[string] `json:"lang" url:"lang,omitempty"`
	Avatar                   Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banner                   Optional[string] `json:"banner" url:"banner,omitempty"`
	DisplayName              Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Email                    Optional[string] `json:"email" url:"email,omitempty"`
	Bio                      Optional[string] `json:"bio" url:"bio,omitempty"`
	MatrixUserID             Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	ShowAvatars              Optional[bool]   `json:"show_avatars" url:"show_avatars,omitempty"`
	SendNotificationsToEmail Optional[bool]   `json:"send_notifications_to_email" url:"send_notifications_to_email,omitempty"`
	BotAccount               Optional[bool]   `json:"bot_account" url:"bot_account,omitempty"`
	ShowBotAccounts          Optional[bool]   `json:"show_bot_accounts" url:"show_bot_accounts,omitempty"`
	ShowReadPosts            Optional[bool]   `json:"show_read_posts" url:"show_read_posts,omitempty"`
	ShowNewPostNotifs        Optional[bool]   `json:"show_new_post_notifs" url:"show_new_post_notifs,omitempty"`
	Auth                     string           `json:"auth" url:"auth,omitempty"`
}
type ChangePassword struct {
	NewPassword       string `json:"new_password" url:"new_password,omitempty"`
	NewPasswordVerify string `json:"new_password_verify" url:"new_password_verify,omitempty"`
	OldPassword       string `json:"old_password" url:"old_password,omitempty"`
	Auth              string `json:"auth" url:"auth,omitempty"`
}
type LoginResponse struct {
	Jwt                 Optional[string] `json:"jwt" url:"jwt,omitempty"`
	RegistrationCreated bool             `json:"registration_created" url:"registration_created,omitempty"`
	VerifyEmailSent     bool             `json:"verify_email_sent" url:"verify_email_sent,omitempty"`
	LemmyResponse
}
type GetPersonDetails struct {
	PersonID    Optional[int]      `json:"person_id" url:"person_id,omitempty"`
	Username    Optional[string]   `json:"username" url:"username,omitempty"`
	Sort        Optional[SortType] `json:"sort" url:"sort,omitempty"`
	Page        Optional[int64]    `json:"page" url:"page,omitempty"`
	Limit       Optional[int64]    `json:"limit" url:"limit,omitempty"`
	CommunityID Optional[int]      `json:"community_id" url:"community_id,omitempty"`
	SavedOnly   Optional[bool]     `json:"saved_only" url:"saved_only,omitempty"`
	Auth        Optional[string]   `json:"auth" url:"auth,omitempty"`
}
type GetPersonDetailsResponse struct {
	PersonView PersonViewSafe           `json:"person_view" url:"person_view,omitempty"`
	Comments   []CommentView            `json:"comments" url:"comments,omitempty"`
	Posts      []PostView               `json:"posts" url:"posts,omitempty"`
	Moderates  []CommunityModeratorView `json:"moderates" url:"moderates,omitempty"`
	LemmyResponse
}
type GetRepliesResponse struct {
	Replies []CommentView `json:"replies" url:"replies,omitempty"`
	LemmyResponse
}
type GetPersonMentionsResponse struct {
	Mentions []PersonMentionView `json:"mentions" url:"mentions,omitempty"`
	LemmyResponse
}
type MarkAllAsRead struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type AddAdmin struct {
	PersonID int    `json:"person_id" url:"person_id,omitempty"`
	Added    bool   `json:"added" url:"added,omitempty"`
	Auth     string `json:"auth" url:"auth,omitempty"`
}
type AddAdminResponse struct {
	Admins []PersonViewSafe `json:"admins" url:"admins,omitempty"`
	LemmyResponse
}
type BanPerson struct {
	PersonID   int              `json:"person_id" url:"person_id,omitempty"`
	Ban        bool             `json:"ban" url:"ban,omitempty"`
	RemoveData Optional[bool]   `json:"remove_data" url:"remove_data,omitempty"`
	Reason     Optional[string] `json:"reason" url:"reason,omitempty"`
	Expires    Optional[int64]  `json:"expires" url:"expires,omitempty"`
	Auth       string           `json:"auth" url:"auth,omitempty"`
}
type GetBannedPersons struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type BannedPersonsResponse struct {
	Banned []PersonViewSafe `json:"banned" url:"banned,omitempty"`
	LemmyResponse
}
type BanPersonResponse struct {
	PersonView PersonViewSafe `json:"person_view" url:"person_view,omitempty"`
	Banned     bool           `json:"banned" url:"banned,omitempty"`
	LemmyResponse
}
type BlockPerson struct {
	PersonID int    `json:"person_id" url:"person_id,omitempty"`
	Block    bool   `json:"block" url:"block,omitempty"`
	Auth     string `json:"auth" url:"auth,omitempty"`
}
type BlockPersonResponse struct {
	PersonView PersonViewSafe `json:"person_view" url:"person_view,omitempty"`
	Blocked    bool           `json:"blocked" url:"blocked,omitempty"`
	LemmyResponse
}
type GetReplies struct {
	Sort       Optional[SortType] `json:"sort" url:"sort,omitempty"`
	Page       Optional[int64]    `json:"page" url:"page,omitempty"`
	Limit      Optional[int64]    `json:"limit" url:"limit,omitempty"`
	UnreadOnly Optional[bool]     `json:"unread_only" url:"unread_only,omitempty"`
	Auth       string             `json:"auth" url:"auth,omitempty"`
}
type GetPersonMentions struct {
	Sort       Optional[SortType] `json:"sort" url:"sort,omitempty"`
	Page       Optional[int64]    `json:"page" url:"page,omitempty"`
	Limit      Optional[int64]    `json:"limit" url:"limit,omitempty"`
	UnreadOnly Optional[bool]     `json:"unread_only" url:"unread_only,omitempty"`
	Auth       string             `json:"auth" url:"auth,omitempty"`
}
type MarkPersonMentionAsRead struct {
	PersonMentionID int    `json:"person_mention_id" url:"person_mention_id,omitempty"`
	Read            bool   `json:"read" url:"read,omitempty"`
	Auth            string `json:"auth" url:"auth,omitempty"`
}
type PersonMentionResponse struct {
	PersonMentionView PersonMentionView `json:"person_mention_view" url:"person_mention_view,omitempty"`
	LemmyResponse
}
type DeleteAccount struct {
	Password string `json:"password" url:"password,omitempty"`
	Auth     string `json:"auth" url:"auth,omitempty"`
}
type DeleteAccountResponse struct {
	LemmyResponse
}
type PasswordReset struct {
	Email string `json:"email" url:"email,omitempty"`
}
type PasswordResetResponse struct {
	LemmyResponse
}
type PasswordChangeAfterReset struct {
	Token          string `json:"token" url:"token,omitempty"`
	Password       string `json:"password" url:"password,omitempty"`
	PasswordVerify string `json:"password_verify" url:"password_verify,omitempty"`
}
type CreatePrivateMessage struct {
	Content     string `json:"content" url:"content,omitempty"`
	RecipientID int    `json:"recipient_id" url:"recipient_id,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type EditPrivateMessage struct {
	PrivateMessageID int    `json:"private_message_id" url:"private_message_id,omitempty"`
	Content          string `json:"content" url:"content,omitempty"`
	Auth             string `json:"auth" url:"auth,omitempty"`
}
type DeletePrivateMessage struct {
	PrivateMessageID int    `json:"private_message_id" url:"private_message_id,omitempty"`
	Deleted          bool   `json:"deleted" url:"deleted,omitempty"`
	Auth             string `json:"auth" url:"auth,omitempty"`
}
type MarkPrivateMessageAsRead struct {
	PrivateMessageID int    `json:"private_message_id" url:"private_message_id,omitempty"`
	Read             bool   `json:"read" url:"read,omitempty"`
	Auth             string `json:"auth" url:"auth,omitempty"`
}
type GetPrivateMessages struct {
	UnreadOnly Optional[bool]  `json:"unread_only" url:"unread_only,omitempty"`
	Page       Optional[int64] `json:"page" url:"page,omitempty"`
	Limit      Optional[int64] `json:"limit" url:"limit,omitempty"`
	Auth       string          `json:"auth" url:"auth,omitempty"`
}
type PrivateMessagesResponse struct {
	PrivateMessages []PrivateMessageView `json:"private_messages" url:"private_messages,omitempty"`
	LemmyResponse
}
type PrivateMessageResponse struct {
	PrivateMessageView PrivateMessageView `json:"private_message_view" url:"private_message_view,omitempty"`
	LemmyResponse
}
type GetReportCount struct {
	CommunityID Optional[int] `json:"community_id" url:"community_id,omitempty"`
	Auth        string        `json:"auth" url:"auth,omitempty"`
}
type GetReportCountResponse struct {
	CommunityID    Optional[int] `json:"community_id" url:"community_id,omitempty"`
	CommentReports int64         `json:"comment_reports" url:"comment_reports,omitempty"`
	PostReports    int64         `json:"post_reports" url:"post_reports,omitempty"`
	LemmyResponse
}
type GetUnreadCount struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type GetUnreadCountResponse struct {
	Replies         int64 `json:"replies" url:"replies,omitempty"`
	Mentions        int64 `json:"mentions" url:"mentions,omitempty"`
	PrivateMessages int64 `json:"private_messages" url:"private_messages,omitempty"`
	LemmyResponse
}
type VerifyEmail struct {
	Token string `json:"token" url:"token,omitempty"`
}
type VerifyEmailResponse struct {
	LemmyResponse
}