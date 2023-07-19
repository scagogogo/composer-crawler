package response

type AdvisoriesResponse struct {
	Advisories map[string][]*Advisory `json:"advisories"`
}

type Advisory struct {
	AdvisoryID         string     `json:"advisoryId"`
	PackageName        string     `json:"packageName"`
	RemoteID           string     `json:"remoteId"`
	Title              string     `json:"title"`
	Link               string     `json:"link"`
	Cve                string     `json:"cve"`
	AffectedVersions   string     `json:"affectedVersions"`
	Source             string     `json:"source"`
	ReportedAt         string     `json:"reportedAt"`
	ComposerRepository string     `json:"composerRepository"`
	Sources            []*Sources `json:"sources"`
}

type Sources struct {
	Name     string `json:"name"`
	RemoteID string `json:"remoteId"`
}
