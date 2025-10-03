package structs

type Homepage struct {
	// Page metadata
	Title           string `json:"title"`
	MetaDescription string `json:"metaDescription"`

	// Personal information
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	JobTitle  string `json:"jobTitle"`
	Tagline   string `json:"tagline"`
	Location  string `json:"location"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`

	// Hero section
	HeroDescription  string `json:"heroDescription"`
	ProfileImage     string `json:"profileImage"`
	AvailableForWork bool   `json:"availableForWork"`

	// About section
	AboutTitle      string       `json:"aboutTitle"`
	AboutIntro      string       `json:"aboutIntro"`
	AboutParagraphs []string     `json:"aboutParagraphs"`
	Skills          []string     `json:"skills"`
	Experience      []Experience `json:"experience"`

	// Projects
	Projects []Project `json:"projects"`

	// Social links
	SocialLinks SocialLinks `json:"socialLinks"`

	// Navigation
	NavItems []NavItem `json:"navItems"`
}

type Experience struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Year        string `json:"year"`
	Description string `json:"description"`
}

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Image        string   `json:"image"`
	Technologies []string `json:"technologies"`
	ProjectURL   string   `json:"projectURL"`
	GithubURL    string   `json:"githubURL"`
	LiveURL      string   `json:"liveURL"`
}

type SocialLinks struct {
	Github   string `json:"github"`
	LinkedIn string `json:"linkedin"`
	Email    string `json:"email"`
}

type NavItem struct {
	Text string `json:"text"`
	Href string `json:"href"`
}
