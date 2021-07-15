package proto

type User struct {
	Badge struct {
		Link    string `json:"link"`
		Primary bool   `json:"primary"`
		Slug    string `json:"slug"`
		Title   string `json:"title"`
	} `json:"badge"`
	Bio               string `json:"bio"`
	Downloads         int64  `json:"downloads"`
	FirstName         string `json:"first_name"`
	FollowedByUser    bool   `json:"followed_by_user"`
	FollowersCount    int64  `json:"followers_count"`
	FollowingCount    int64  `json:"following_count"`
	ID                string `json:"id"`
	InstagramUsername string `json:"instagram_username"`
	LastName          string `json:"last_name"`
	Links             struct {
		HTML      string `json:"html"`
		Likes     string `json:"likes"`
		Photos    string `json:"photos"`
		Portfolio string `json:"portfolio"`
		Self      string `json:"self"`
	} `json:"links"`
	Location     string      `json:"location"`
	Name         string      `json:"name"`
	PortfolioURL interface{} `json:"portfolio_url"`
	ProfileImage struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"profile_image"`
	TotalCollections int64  `json:"total_collections"`
	TotalLikes       int64  `json:"total_likes"`
	TotalPhotos      int64  `json:"total_photos"`
	TwitterUsername  string `json:"twitter_username"`
	UpdatedAt        string `json:"updated_at"`
	Username         string `json:"username"`
}

type PortfolioLink struct {
	Url string
}

type UserStatistics struct {
	Downloads struct {
		Historical struct {
			Average    int64  `json:"average"`
			Change     int64  `json:"change"`
			Quantity   int64  `json:"quantity"`
			Resolution string `json:"resolution"`
			Values     []struct {
				Date  string `json:"date"`
				Value int64  `json:"value"`
			} `json:"values"`
		} `json:"historical"`
		Total int64 `json:"total"`
	} `json:"downloads"`
	Username string `json:"username"`
	Views    struct {
		Historical struct {
			Average    int64  `json:"average"`
			Change     int64  `json:"change"`
			Quantity   int64  `json:"quantity"`
			Resolution string `json:"resolution"`
			Values     []struct {
				Date  string `json:"date"`
				Value int64  `json:"value"`
			} `json:"values"`
		} `json:"historical"`
		Total int64 `json:"total"`
	} `json:"views"`
}

