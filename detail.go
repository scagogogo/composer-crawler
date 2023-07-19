package composer_crawler

import "time"

type ComposerPackageInfo struct {

	// 包的名字，用来做唯一主键
	PackageName string `json:"package_name" bson:"package_name"`
	// 某些情况下需要用到小写，我他妈也无语
	PackageNameLowercase string `json:"package_name_lowercase" bson:"package_name_lowercase"`

	Package struct {

		// 包的名字
		Name string `json:"name" bson:"name"`

		// 包描述信息
		Description string `json:"description" bson:"description"`

		// 包被创建的时间
		Time time.Time `json:"time" bson:"time"`

		// 维护者的信息
		Maintainers []*Maintainers `json:"maintainers" bson:"maintainers"`

		// 这个包的所有版本相关信息
		Versions map[string]*Version `json:"versions" bson:"versions"`

		// 包的类型
		Type string `json:"type" bson:"type"`
		// 包所对应的仓库
		Repository string `json:"repository" bson:"repository"`
		// 仓库的star有多少
		GithubStars int `json:"github_stars" bson:"github_stars"`
		// 仓库的watch有多少
		GithubWatchers int `json:"github_watchers" bson:"github_watchers"`
		// 仓库的fork有多少
		GithubForks int `json:"github_forks" bson:"github_forks"`
		// 仓库的open issue
		GithubOpenIssues int `json:"github_open_issues" bson:"github_open_issues"`
		// 仓库是啥语言？难不成还会有别的语言？
		Language string `json:"language" bson:"language"`
		// 被多少个其它包所依赖
		Dependents int `json:"dependents" bson:"dependents"`
		//
		Suggesters int `json:"suggesters" bson:"suggesters"`
		// 下载信息
		Downloads struct {
			// 历史上总共被下载过多少次
			Total int `json:"total" bson:"total"`
			// 本月被下载过多少次
			Monthly int `json:"monthly" bson:"monthly"`
			// 今日被下载过多少次
			Daily int `json:"daily" bson:"daily"`
		} `json:"downloads" bson:"downloads"`
		Favers int `json:"favers" bson:"favers"`
	} `json:"package" bson:"package"`

	// 包相关信息的md5，用来识别信息是否较之前发生了变化
	PackageInfoMd5 string `json:"package_info_md5" bson:"package_info_md5"`

	// 三个时间戳
	CreateTime *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime *time.Time `json:"update_time" bson:"update_time"`
	ChangeTime *time.Time `json:"change_time" bson:"change_time"`
}

// Maintainers 表示一个包维护者相关的信息
type Maintainers struct {
	// 维护者的名字
	Name string `json:"name" bson:"name"`
	// 维护者的头像
	AvatarURL string `json:"avatar_url" bson:"avatar_url"`
}

// Version 表示包的一个版本信息
type Version struct {
	// 版本的名字
	Name string `json:"name"`
	// 版本描述
	Description string `json:"description"`
	// 版本关键词
	Keywords []string `json:"keywords"`
	// 首页
	Homepage string `json:"homepage"`
	// 版本号
	Version string `json:"version"`
	// 版本号的格式统一化
	VersionNormalized string `json:"version_normalized"`
	// 这个版本所使用的license
	License []string `json:"license" bson:"license"`
	Authors []struct {
		Name     string `json:"name" bson:"name"`
		Email    string `json:"email" bson:"email"`
		Homepage string `json:"homepage" bson:"homepage"`
	} `json:"authors" bson:"authors"`
	Source struct {
		URL       string `json:"url" bson:"url"`
		Type      string `json:"type" bson:"type"`
		Reference string `json:"reference" bson:"reference"`
	} `json:"source" bson:"source"`
	Dist struct {
		URL       string `json:"url" bson:"url"`
		Type      string `json:"type" bson:"type"`
		Shasum    string `json:"shasum" bson:"shasum"`
		Reference string `json:"reference" bson:"reference"`
	} `json:"dist" bson:"dist"`
	Type    string `json:"type" bson:"type"`
	Support struct {
		Source string `json:"source" bson:"source"`
	} `json:"support" bson:"support"`
	Funding []struct {
		URL  string `json:"url" bson:"url"`
		Type string `json:"type" bson:"type"`
	} `json:"funding" bson:"funding"`

	// 可以看做是版本的发布时间
	Time time.Time `json:"time" bson:"time"`

	// 这个版本所依赖的其它包的其它版本
	Require map[string]string `json:"require" bson:"require"`

	// 2022-5-29 20:57: 我他妈都不清楚这些字段是啥意思，也不想去弄清楚了，直接Interface吧，后面用到的话再来搞明白啥意思
	Autoload interface{} `json:"autoload" bson:"autoload"`
	Extra    interface{} `json:"extra" bson:"extra"`
	Suggest  interface{} `json:"suggest" bson:"suggest"`
	Provide  interface{} `json:"provide" bson:"provide"`
}
