package proto

type Photo struct {
	AltDescription         interface{}   `json:"alt_description"`
	BlurHash               string        `json:"blur_hash"`
	Categories             []interface{} `json:"categories"`
	Color                  string        `json:"color"`
	CreatedAt              string        `json:"created_at"`
	CurrentUserCollections []Collection `json:"current_user_collections"`
	Description            string        `json:"description"`
	Downloads              int64         `json:"downloads"`
	Exif                   struct {
		Aperture     string `json:"aperture"`
		ExposureTime string `json:"exposure_time"`
		FocalLength  string `json:"focal_length"`
		Iso          int64  `json:"iso"`
		Make         string `json:"make"`
		Model        string `json:"model"`
	} `json:"exif"`
	Height      int64  `json:"height"`
	ID          string `json:"id"`
	LikedByUser bool   `json:"liked_by_user"`
	Likes       int64  `json:"likes"`
	Links       struct {
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
		HTML             string `json:"html"`
		Self             string `json:"self"`
	} `json:"links"`
	Location struct {
		City     interface{} `json:"city"`
		Country  interface{} `json:"country"`
		Name     interface{} `json:"name"`
		Position struct {
			Latitude  interface{} `json:"latitude"`
			Longitude interface{} `json:"longitude"`
		} `json:"position"`
		Title interface{} `json:"title"`
	} `json:"location"`
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
		TotalCollections int64       `json:"total_collections"`
		TotalLikes       int64       `json:"total_likes"`
		TotalPhotos      int64       `json:"total_photos"`
		TwitterUsername  interface{} `json:"twitter_username"`
		UpdatedAt        string      `json:"updated_at"`
		Username         string      `json:"username"`
	} `json:"user"`
	Views int64 `json:"views"`
	Width int64 `json:"width"`
}

type PhotoStatistics struct {
	Downloads struct {
		Historical struct {
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
	ID    string `json:"id"`
	Likes struct {
		Historical struct {
			Change     int64  `json:"change"`
			Quantity   int64  `json:"quantity"`
			Resolution string `json:"resolution"`
			Values     []struct {
				Date  string `json:"date"`
				Value int64  `json:"value"`
			} `json:"values"`
		} `json:"historical"`
		Total int64 `json:"total"`
	} `json:"likes"`
	Views struct {
		Historical struct {
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


type PhotoLike struct {
	Photo struct {
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
		Width int64 `json:"width"`
	} `json:"photo"`
	User struct {
		ID    string `json:"id"`
		Links struct {
			HTML   string `json:"html"`
			Likes  string `json:"likes"`
			Photos string `json:"photos"`
			Self   string `json:"self"`
		} `json:"links"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"user"`
}

