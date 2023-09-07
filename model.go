package fullstory

type Session struct {
	// ID is the fullstory session id
	ID string `json:"id"`
}

type Browser struct {
	URL            string `json:"url,omitempty"`
	UserAgent      string `json:"user_agent,omitempty"`
	InitialReferer string `json:"initial_referer,omitempty"`
}

type Mobile struct {
	AppID        string `json:"app_id,omitempty"`
	AppVersion   string `json:"app_version,omitempty"`
	AppName      string `json:"app_name,omitempty"`
	BuildVariant string `json:"build_variant,omitempty"`
}

type Device struct {
	Manufacturer   string `json:"manufacturer,omitempty"`
	Model          string `json:"model,omitempty"`
	ScreenWidth    int32  `json:"screen_width,omitempty"`
	ScreenHeight   int32  `json:"screen_height,omitempty"`
	ViewportWidth  int32  `json:"viewport_width,omitempty"`
	ViewportHeight int32  `json:"viewport_height,omitempty"`
}

type Location struct{}

type Context struct {
	Browser  Browser  `json:"browser,omitempty"`
	Mobile   Mobile   `json:"mobile,omitempty"`
	Device   Device   `json:"device,omitempty"`
	Location Location `json:"location,omitempty"`
}
