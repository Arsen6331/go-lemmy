package types

type UserOperation string

const (
	UserOperationLogin                                 = "Login"
	UserOperationGetCaptcha                            = "GetCaptcha"
	UserOperationMarkCommentAsRead                     = "MarkCommentAsRead"
	UserOperationSaveComment                           = "SaveComment"
	UserOperationCreateCommentLike                     = "CreateCommentLike"
	UserOperationCreateCommentReport                   = "CreateCommentReport"
	UserOperationResolveCommentReport                  = "ResolveCommentReport"
	UserOperationListCommentReports                    = "ListCommentReports"
	UserOperationCreatePostLike                        = "CreatePostLike"
	UserOperationLockPost                              = "LockPost"
	UserOperationStickyPost                            = "StickyPost"
	UserOperationMarkPostAsRead                        = "MarkPostAsRead"
	UserOperationSavePost                              = "SavePost"
	UserOperationCreatePostReport                      = "CreatePostReport"
	UserOperationResolvePostReport                     = "ResolvePostReport"
	UserOperationListPostReports                       = "ListPostReports"
	UserOperationGetReportCount                        = "GetReportCount"
	UserOperationGetUnreadCount                        = "GetUnreadCount"
	UserOperationVerifyEmail                           = "VerifyEmail"
	UserOperationFollowCommunity                       = "FollowCommunity"
	UserOperationGetReplies                            = "GetReplies"
	UserOperationGetPersonMentions                     = "GetPersonMentions"
	UserOperationMarkPersonMentionAsRead               = "MarkPersonMentionAsRead"
	UserOperationGetModlog                             = "GetModlog"
	UserOperationBanFromCommunity                      = "BanFromCommunity"
	UserOperationAddModToCommunity                     = "AddModToCommunity"
	UserOperationAddAdmin                              = "AddAdmin"
	UserOperationGetUnreadRegistrationApplicationCount = "GetUnreadRegistrationApplicationCount"
	UserOperationListRegistrationApplications          = "ListRegistrationApplications"
	UserOperationApproveRegistrationApplication        = "ApproveRegistrationApplication"
	UserOperationBanPerson                             = "BanPerson"
	UserOperationGetBannedPersons                      = "GetBannedPersons"
	UserOperationSearch                                = "Search"
	UserOperationResolveObject                         = "ResolveObject"
	UserOperationMarkAllAsRead                         = "MarkAllAsRead"
	UserOperationSaveUserSettings                      = "SaveUserSettings"
	UserOperationTransferCommunity                     = "TransferCommunity"
	UserOperationLeaveAdmin                            = "LeaveAdmin"
	UserOperationPasswordReset                         = "PasswordReset"
	UserOperationPasswordChange                        = "PasswordChange"
	UserOperationMarkPrivateMessageAsRead              = "MarkPrivateMessageAsRead"
	UserOperationUserJoin                              = "UserJoin"
	UserOperationGetSiteConfig                         = "GetSiteConfig"
	UserOperationSaveSiteConfig                        = "SaveSiteConfig"
	UserOperationPostJoin                              = "PostJoin"
	UserOperationCommunityJoin                         = "CommunityJoin"
	UserOperationModJoin                               = "ModJoin"
	UserOperationChangePassword                        = "ChangePassword"
	UserOperationGetSiteMetadata                       = "GetSiteMetadata"
	UserOperationBlockCommunity                        = "BlockCommunity"
	UserOperationBlockPerson                           = "BlockPerson"
)

type UserOperationCrud string

const (
	UserOperationCrudCreateSite           = "CreateSite"
	UserOperationCrudGetSite              = "GetSite"
	UserOperationCrudEditSite             = "EditSite"
	UserOperationCrudCreateCommunity      = "CreateCommunity"
	UserOperationCrudListCommunities      = "ListCommunities"
	UserOperationCrudGetCommunity         = "GetCommunity"
	UserOperationCrudEditCommunity        = "EditCommunity"
	UserOperationCrudDeleteCommunity      = "DeleteCommunity"
	UserOperationCrudRemoveCommunity      = "RemoveCommunity"
	UserOperationCrudCreatePost           = "CreatePost"
	UserOperationCrudGetPost              = "GetPost"
	UserOperationCrudGetPosts             = "GetPosts"
	UserOperationCrudEditPost             = "EditPost"
	UserOperationCrudDeletePost           = "DeletePost"
	UserOperationCrudRemovePost           = "RemovePost"
	UserOperationCrudCreateComment        = "CreateComment"
	UserOperationCrudGetComment           = "GetComment"
	UserOperationCrudGetComments          = "GetComments"
	UserOperationCrudEditComment          = "EditComment"
	UserOperationCrudDeleteComment        = "DeleteComment"
	UserOperationCrudRemoveComment        = "RemoveComment"
	UserOperationCrudRegister             = "Register"
	UserOperationCrudGetPersonDetails     = "GetPersonDetails"
	UserOperationCrudDeleteAccount        = "DeleteAccount"
	UserOperationCrudCreatePrivateMessage = "CreatePrivateMessage"
	UserOperationCrudGetPrivateMessages   = "GetPrivateMessages"
	UserOperationCrudEditPrivateMessage   = "EditPrivateMessage"
	UserOperationCrudDeletePrivateMessage = "DeletePrivateMessage"
)
