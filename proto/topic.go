package proto

type Topic struct {
	CoverPhoto struct {
		AltDescription         string        `json:"alt_description"`
		BlurHash               string        `json:"blur_hash"`
		Categories             []interface{} `json:"categories"`
		Color                  string        `json:"color"`
		CreatedAt              string        `json:"created_at"`
		CurrentUserCollections []interface{} `json:"current_user_collections"`
		Description            string        `json:"description"`
		Height                 int64         `json:"height"`
		ID                     string        `json:"id"`
		LikedByUser            bool          `json:"liked_by_user"`
		Likes                  int64         `json:"likes"`
		Links                  struct {
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
			HTML             string `json:"html"`
			Self             string `json:"self"`
		} `json:"links"`
		PromotedAt  string      `json:"promoted_at"`
		Sponsorship interface{} `json:"sponsorship"`
		UpdatedAt   string      `json:"updated_at"`
		Urls        struct {
			Full    string `json:"full"`
			Raw     string `json:"raw"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
		User struct {
			AcceptedTos       bool   `json:"accepted_tos"`
			Bio               string `json:"bio"`
			FirstName         string `json:"first_name"`
			ForHire           bool   `json:"for_hire"`
			ID                string `json:"id"`
			InstagramUsername string `json:"instagram_username"`
			LastName          string `json:"last_name"`
			Links             struct {
				Followers string `json:"followers"`
				Following string `json:"following"`
				HTML      string `json:"html"`
				Likes     string `json:"likes"`
				Photos    string `json:"photos"`
				Portfolio string `json:"portfolio"`
				Self      string `json:"self"`
			} `json:"links"`
			Location     string `json:"location"`
			Name         string `json:"name"`
			PortfolioURL string `json:"portfolio_url"`
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
		} `json:"user"`
		Width int64 `json:"width"`
	} `json:"cover_photo"`
	CurrentUserContributions []interface{} `json:"current_user_contributions"`
	Description              string        `json:"description"`
	EndsAt                   string        `json:"ends_at"`
	Featured                 bool          `json:"featured"`
	ID                       string        `json:"id"`
	Links                    struct {
		HTML   string `json:"html"`
		Photos string `json:"photos"`
		Self   string `json:"self"`
	} `json:"links"`
	OnlySubmissionsAfter interface{} `json:"only_submissions_after"`
	Owners               []struct {
		AcceptedTos       bool        `json:"accepted_tos"`
		Bio               string      `json:"bio"`
		FirstName         string      `json:"first_name"`
		ForHire           bool        `json:"for_hire"`
		ID                string      `json:"id"`
		InstagramUsername string      `json:"instagram_username"`
		LastName          interface{} `json:"last_name"`
		Links             struct {
			Followers string `json:"followers"`
			Following string `json:"following"`
			HTML      string `json:"html"`
			Likes     string `json:"likes"`
			Photos    string `json:"photos"`
			Portfolio string `json:"portfolio"`
			Self      string `json:"self"`
		} `json:"links"`
		Location     interface{} `json:"location"`
		Name         string      `json:"name"`
		PortfolioURL string      `json:"portfolio_url"`
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
	} `json:"owners"`
	PreviewPhotos []struct {
		BlurHash  string `json:"blur_hash"`
		CreatedAt string `json:"created_at"`
		ID        string `json:"id"`
		UpdatedAt string `json:"updated_at"`
		Urls      struct {
			Full    string `json:"full"`
			Raw     string `json:"raw"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
	} `json:"preview_photos"`
	PublishedAt                 string      `json:"published_at"`
	Slug                        string      `json:"slug"`
	StartsAt                    string      `json:"starts_at"`
	Status                      string      `json:"status"`
	Title                       string      `json:"title"`
	TotalCurrentUserSubmissions interface{} `json:"total_current_user_submissions"`
	TotalPhotos                 int64       `json:"total_photos"`
	UpdatedAt                   string      `json:"updated_at"`
}


