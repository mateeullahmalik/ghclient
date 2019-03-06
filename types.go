package ghclient

import (
	"time"
)

// Event is triggered by the GitHub app.
type Event struct {
	Action       string       `json:"action"`
	CheckRun     CheckRun     `json:"check_run"`
	Installation Installation `json:"installation"`
	Number       int          `json:"number"`
	Organization Organization `json:"organization"`
	PullRequest  PullRequest  `json:"pull_request"`
	Repository   Repository   `json:"repository"`
	Repositories []Repository `json:"repositories"`
	Sender       Sender       `json:"sender"`
}

// InstallationEvent is triggered when a GitHub app is either installed or removed.
type InstallationEvent struct {
	Action       string       `json:"action"`
	Installation Installation `json:"installation"`
	Repositories []Repository `json:"repositories"`
	Sender       Sender       `json:"sender"`
}

// CheckRunEvent is triggered when a Check Run is created, rerequested, completed or has a requested action.
type CheckRunEvent struct {
	Action       string       `json:"action"`
	CheckRun     CheckRun     `json:"check_run"`
	Repository   Repository   `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       Sender       `json:"sender"`
	Installation Installation `json:"installation"`
}

// PullRequestEvent is triggered when a Pull Request is assigned, unassigned, labeled, unlabeled, opened, edited, closed, reopened, synchronized
type PullRequestEvent struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      Sender      `json:"sender"`
}

// Output used provide information back to the requester
type Output struct {
	Title            string `json:"title"`
	Summary          string `json:"summary"`
	Text             string `json:"text"`
	AnnotationsCount int    `json:"annotations_count"`
	AnnotationsURL   string `json:"annotations_url"`
}

// CheckRun is used to run different types of checks (quality, security, dependency) against a repository.
type CheckRun struct {
	ID           int           `json:"id"`
	HeadSha      string        `json:"head_sha"`
	ExternalID   string        `json:"external_id"`
	URL          string        `json:"url"`
	HTMLURL      string        `json:"html_url"`
	Status       string        `json:"status"`
	Conclusion   string        `json:"conclusion"`
	StartedAt    time.Time     `json:"started_at"`
	CompletedAt  time.Time     `json:"completed_at"`
	Output       Output        `json:"output"`
	Name         string        `json:"name"`
	CheckSuite   CheckSuite    `json:"check_suite"`
	App          App           `json:"app"`
	PullRequests []interface{} `json:"pull_requests"`
}

// CheckSuite contains one or more Check Runs
type CheckSuite struct {
	ID           int           `json:"id"`
	HeadBranch   string        `json:"head_branch"`
	HeadSha      string        `json:"head_sha"`
	Status       string        `json:"status"`
	Conclusion   string        `json:"conclusion"`
	URL          string        `json:"url"`
	Before       string        `json:"before"`
	After        string        `json:"after"`
	PullRequests []interface{} `json:"pull_requests"`
	App          App           `json:"app"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// App contains information about the GitHub Application
type App struct {
	ID          int         `json:"id"`
	NodeID      string      `json:"node_id"`
	Owner       Owner       `json:"owner"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	ExternalURL string      `json:"external_url"`
	HTMLURL     string      `json:"html_url"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

// Repository contains information about the repository, can be used to reference a specific commit (last, current head, etc.)
type Repository struct {
	ID               int         `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Owner            Owner       `json:"owner"`
	Private          bool        `json:"private"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         interface{} `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
}

// Owner contains details about the Repository or App owner
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Organization contains details about the GitHub Organization
type Organization struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Installation contains details about the Installation
type Installation struct {
	ID                  int         `json:"id"`
	Account             Account     `json:"account"`
	RepositorySelection string      `json:"repository_selection"`
	AccessTokensURL     string      `json:"access_tokens_url"`
	RepositoriesURL     string      `json:"repositories_url"`
	HTMLURL             string      `json:"html_url"`
	AppID               int         `json:"app_id"`
	TargetID            int         `json:"target_id"`
	TargetType          string      `json:"target_type"`
	Permissions         Permissions `json:"permissions"`
	Events              []string    `json:"events"`
	CreatedAt           int         `json:"created_at"`
	UpdatedAt           int         `json:"updated_at"`
	SingleFileName      string      `json:"single_file_name"`
}

// Sender provides details about the person or service triggering the event
type Sender struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// PullRequest provides details about the pull request
type PullRequest struct {
	URL                 string        `json:"url"`
	ID                  int           `json:"id"`
	NodeID              string        `json:"node_id"`
	HTMLURL             string        `json:"html_url"`
	DiffURL             string        `json:"diff_url"`
	PatchURL            string        `json:"patch_url"`
	IssueURL            string        `json:"issue_url"`
	Number              int           `json:"number"`
	State               string        `json:"state"`
	Locked              bool          `json:"locked"`
	Title               string        `json:"title"`
	User                User          `json:"user"`
	Body                string        `json:"body"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	ClosedAt            time.Time     `json:"closed_at"`
	MergedAt            interface{}   `json:"merged_at"`
	MergeCommitSha      string        `json:"merge_commit_sha"`
	Assignee            interface{}   `json:"assignee"`
	Assignees           []interface{} `json:"assignees"`
	RequestedReviewers  []interface{} `json:"requested_reviewers"`
	RequestedTeams      []interface{} `json:"requested_teams"`
	Labels              []interface{} `json:"labels"`
	Milestone           interface{}   `json:"milestone"`
	CommitsURL          string        `json:"commits_url"`
	ReviewCommentsURL   string        `json:"review_comments_url"`
	ReviewCommentURL    string        `json:"review_comment_url"`
	CommentsURL         string        `json:"comments_url"`
	StatusesURL         string        `json:"statuses_url"`
	Head                Head          `json:"head"`
	Base                Base          `json:"base"`
	Links               Links         `json:"_links"`
	AuthorAssociation   string        `json:"author_association"`
	Merged              bool          `json:"merged"`
	Mergeable           bool          `json:"mergeable"`
	Rebaseable          bool          `json:"rebaseable"`
	MergeableState      string        `json:"mergeable_state"`
	MergedBy            interface{}   `json:"merged_by"`
	Comments            int           `json:"comments"`
	ReviewComments      int           `json:"review_comments"`
	MaintainerCanModify bool          `json:"maintainer_can_modify"`
	Commits             int           `json:"commits"`
	Additions           int           `json:"additions"`
	Deletions           int           `json:"deletions"`
	ChangedFiles        int           `json:"changed_files"`
}

// Links to related data for Pull Request
type Links struct {
	Self           Self           `json:"self"`
	HTML           HTML           `json:"html"`
	Issue          Issue          `json:"issue"`
	Comments       Comments       `json:"comments"`
	ReviewComments ReviewComments `json:"review_comments"`
	ReviewComment  ReviewComment  `json:"review_comment"`
	Commits        Commits        `json:"commits"`
	Statuses       Statuses       `json:"statuses"`
}

// Self href reference
type Self struct {
	Href string `json:"href"`
}

// HTML href reference
type HTML struct {
	Href string `json:"href"`
}

// Issue href reference
type Issue struct {
	Href string `json:"href"`
}

// Comments href reference
type Comments struct {
	Href string `json:"href"`
}

// ReviewComments href reference
type ReviewComments struct {
	Href string `json:"href"`
}

// ReviewComment href reference
type ReviewComment struct {
	Href string `json:"href"`
}

// Commits href reference
type Commits struct {
	Href string `json:"href"`
}

// Statuses href reference
type Statuses struct {
	Href string `json:"href"`
}

// Base commit reference
type Base struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo   `json:"repo"`
}

// Head commit reference
type Head struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo   `json:"repo"`
}

// User details about the authticated person making the request
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Repo contains the details about a specific repository
type Repo struct {
	ID               int         `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Owner            Owner       `json:"owner"`
	Private          bool        `json:"private"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         interface{} `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
}

// Permissions provides permission details a specific installation
type Permissions struct {
	Metadata string `json:"metadata"`
	Contents string `json:"contents"`
	Issues   string `json:"issues"`
}

// Account provides details about a specific account where the app is installed
type Account struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
