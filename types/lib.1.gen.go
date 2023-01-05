package types

type SortType string

const (
	SortTypeActive       SortType = "Active"
	SortTypeHot          SortType = "Hot"
	SortTypeNew          SortType = "New"
	SortTypeTopDay       SortType = "TopDay"
	SortTypeTopWeek      SortType = "TopWeek"
	SortTypeTopMonth     SortType = "TopMonth"
	SortTypeTopYear      SortType = "TopYear"
	SortTypeTopAll       SortType = "TopAll"
	SortTypeMostComments SortType = "MostComments"
	SortTypeNewComments  SortType = "NewComments"
)

type ListingType string

const (
	ListingTypeAll        ListingType = "All"
	ListingTypeLocal      ListingType = "Local"
	ListingTypeSubscribed ListingType = "Subscribed"
	ListingTypeCommunity  ListingType = "Community"
)

type SearchType string

const (
	SearchTypeAll         SearchType = "All"
	SearchTypeComments    SearchType = "Comments"
	SearchTypePosts       SearchType = "Posts"
	SearchTypeCommunities SearchType = "Communities"
	SearchTypeUsers       SearchType = "Users"
	SearchTypeUrl         SearchType = "Url"
)