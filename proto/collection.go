package proto

type Collection struct {
	CoverPhoto struct {
		BlurHash    string `json:"blur_hash"`
		Color       string `json:"color"`
		Description string `json:"description"`
		Height      int64  `json:"height"`
		ID          string `json:"id"`
		LikedByUser bool   `json:"liked_by_user"`
		Likes       int64  `json:"likes"`
		Links       struct {
			Download string `json:"download"`
			HTML     string `json:"html"`
			Self     string `json:"self"`
		} `json:"links"`
		Urls struct {
			Full    string `json:"full"`
			Raw     string `json:"raw"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
		} `json:"urls"`
		User struct {
			Bio   string `json:"bio"`
			ID    string `json:"id"`
			Links struct {
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
			Username         string `json:"username"`
		} `json:"user"`
		Width int64 `json:"width"`
	} `json:"cover_photo"`
	Description     string `json:"description"`
	ID              string  `json:"id"`
	LastCollectedAt string `json:"last_collected_at"`
	Links           struct {
		HTML    string `json:"html"`
		Photos  string `json:"photos"`
		Related string `json:"related"`
		Self    string `json:"self"`
	} `json:"links"`
	Private     bool   `json:"private"`
	PublishedAt string `json:"published_at"`
	ShareKey    string `json:"share_key"`
	Title       string `json:"title"`
	TotalPhotos int64  `json:"total_photos"`
	UpdatedAt   string `json:"updated_at"`
	User        struct {
		Bio   string `json:"bio"`
		ID    string `json:"id"`
		Links struct {
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
		UpdatedAt        string `json:"updated_at"`
		Username         string `json:"username"`
	} `json:"user"`
}

type PhotoOperationOfCollection struct{
	Photo Photo
	Collection Collection
	User User
	CreatedAt string `json:"created_at"`
}

