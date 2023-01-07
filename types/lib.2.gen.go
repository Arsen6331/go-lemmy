//  Source: lemmy/crates/websocket/src/lib.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type UserOperation string

const (
	UserOperationLogin                                 UserOperation = "Login"
	UserOperationGetCaptcha                            UserOperation = "GetCaptcha"
	UserOperationMarkCommentAsRead                     UserOperation = "MarkCommentAsRead"
	UserOperationSaveComment                           UserOperation = "SaveComment"
	UserOperationCreateCommentLike                     UserOperation = "CreateCommentLike"
	UserOperationCreateCommentReport                   UserOperation = "CreateCommentReport"
	UserOperationResolveCommentReport                  UserOperation = "ResolveCommentReport"
	UserOperationListCommentReports                    UserOperation = "ListCommentReports"
	UserOperationCreatePostLike                        UserOperation = "CreatePostLike"
	UserOperationLockPost                              UserOperation = "LockPost"
	UserOperationStickyPost                            UserOperation = "StickyPost"
	UserOperationMarkPostAsRead                        UserOperation = "MarkPostAsRead"
	UserOperationSavePost                              UserOperation = "SavePost"
	UserOperationCreatePostReport                      UserOperation = "CreatePostReport"
	UserOperationResolvePostReport                     UserOperation = "ResolvePostReport"
	UserOperationListPostReports                       UserOperation = "ListPostReports"
	UserOperationGetReportCount                        UserOperation = "GetReportCount"
	UserOperationGetUnreadCount                        UserOperation = "GetUnreadCount"
	UserOperationVerifyEmail                           UserOperation = "VerifyEmail"
	UserOperationFollowCommunity                       UserOperation = "FollowCommunity"
	UserOperationGetReplies                            UserOperation = "GetReplies"
	UserOperationGetPersonMentions                     UserOperation = "GetPersonMentions"
	UserOperationMarkPersonMentionAsRead               UserOperation = "MarkPersonMentionAsRead"
	UserOperationGetModlog                             UserOperation = "GetModlog"
	UserOperationBanFromCommunity                      UserOperation = "BanFromCommunity"
	UserOperationAddModToCommunity                     UserOperation = "AddModToCommunity"
	UserOperationAddAdmin                              UserOperation = "AddAdmin"
	UserOperationGetUnreadRegistrationApplicationCount UserOperation = "GetUnreadRegistrationApplicationCount"
	UserOperationListRegistrationApplications          UserOperation = "ListRegistrationApplications"
	UserOperationApproveRegistrationApplication        UserOperation = "ApproveRegistrationApplication"
	UserOperationBanPerson                             UserOperation = "BanPerson"
	UserOperationGetBannedPersons                      UserOperation = "GetBannedPersons"
	UserOperationSearch                                UserOperation = "Search"
	UserOperationResolveObject                         UserOperation = "ResolveObject"
	UserOperationMarkAllAsRead                         UserOperation = "MarkAllAsRead"
	UserOperationSaveUserSettings                      UserOperation = "SaveUserSettings"
	UserOperationTransferCommunity                     UserOperation = "TransferCommunity"
	UserOperationLeaveAdmin                            UserOperation = "LeaveAdmin"
	UserOperationPasswordReset                         UserOperation = "PasswordReset"
	UserOperationPasswordChange                        UserOperation = "PasswordChange"
	UserOperationMarkPrivateMessageAsRead              UserOperation = "MarkPrivateMessageAsRead"
	UserOperationUserJoin                              UserOperation = "UserJoin"
	UserOperationGetSiteConfig                         UserOperation = "GetSiteConfig"
	UserOperationSaveSiteConfig                        UserOperation = "SaveSiteConfig"
	UserOperationPostJoin                              UserOperation = "PostJoin"
	UserOperationCommunityJoin                         UserOperation = "CommunityJoin"
	UserOperationModJoin                               UserOperation = "ModJoin"
	UserOperationChangePassword                        UserOperation = "ChangePassword"
	UserOperationGetSiteMetadata                       UserOperation = "GetSiteMetadata"
	UserOperationBlockCommunity                        UserOperation = "BlockCommunity"
	UserOperationBlockPerson                           UserOperation = "BlockPerson"
)

type UserOperationCRUD string

const (
	UserOperationCRUDCreateSite           UserOperationCRUD = "CreateSite"
	UserOperationCRUDGetSite              UserOperationCRUD = "GetSite"
	UserOperationCRUDEditSite             UserOperationCRUD = "EditSite"
	UserOperationCRUDCreateCommunity      UserOperationCRUD = "CreateCommunity"
	UserOperationCRUDListCommunities      UserOperationCRUD = "ListCommunities"
	UserOperationCRUDGetCommunity         UserOperationCRUD = "GetCommunity"
	UserOperationCRUDEditCommunity        UserOperationCRUD = "EditCommunity"
	UserOperationCRUDDeleteCommunity      UserOperationCRUD = "DeleteCommunity"
	UserOperationCRUDRemoveCommunity      UserOperationCRUD = "RemoveCommunity"
	UserOperationCRUDCreatePost           UserOperationCRUD = "CreatePost"
	UserOperationCRUDGetPost              UserOperationCRUD = "GetPost"
	UserOperationCRUDGetPosts             UserOperationCRUD = "GetPosts"
	UserOperationCRUDEditPost             UserOperationCRUD = "EditPost"
	UserOperationCRUDDeletePost           UserOperationCRUD = "DeletePost"
	UserOperationCRUDRemovePost           UserOperationCRUD = "RemovePost"
	UserOperationCRUDCreateComment        UserOperationCRUD = "CreateComment"
	UserOperationCRUDGetComment           UserOperationCRUD = "GetComment"
	UserOperationCRUDGetComments          UserOperationCRUD = "GetComments"
	UserOperationCRUDEditComment          UserOperationCRUD = "EditComment"
	UserOperationCRUDDeleteComment        UserOperationCRUD = "DeleteComment"
	UserOperationCRUDRemoveComment        UserOperationCRUD = "RemoveComment"
	UserOperationCRUDRegister             UserOperationCRUD = "Register"
	UserOperationCRUDGetPersonDetails     UserOperationCRUD = "GetPersonDetails"
	UserOperationCRUDDeleteAccount        UserOperationCRUD = "DeleteAccount"
	UserOperationCRUDCreatePrivateMessage UserOperationCRUD = "CreatePrivateMessage"
	UserOperationCRUDGetPrivateMessages   UserOperationCRUD = "GetPrivateMessages"
	UserOperationCRUDEditPrivateMessage   UserOperationCRUD = "EditPrivateMessage"
	UserOperationCRUDDeletePrivateMessage UserOperationCRUD = "DeletePrivateMessage"
)
