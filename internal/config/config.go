package config

type Config struct {
	General GeneralConfig `json:"general"`
	HTTP    HTTPConfig    `json:"http"`
	Filter  FilterConfig  `json:"filter"`

	Input   InputConfig   `json:"input"`
	Matcher MatcherConfig `json:"matcher"`
	Output  OutputConfig  `json:"output"`
}
type GeneralConfig struct {
	AutoCalibration           bool     `json:"autocalibration"`
	AutoCalibrationKeyword    string   `json:"autocalibration_keyword"`
	AutoCalibrationPerHost    bool     `json:"autocalibration_per_host"`
	AutoCalibrationStrategies []string `json:"autocalibration_strategies"`
	AutoCalibrationStrings    []string `json:"autocalibration_strings"`
	Colors                    bool     `json:"colors"`
	ConfigFile                string   `toml:"-" json:"config_file"`
	Delay                     string   `json:"delay"`
	Json                      bool     `json:"json"`
	MaxTime                   int      `json:"maxtime"`
	MaxTimeJob                int      `json:"maxtime_job"`
	Noninteractive            bool     `json:"noninteractive"`
	Quiet                     bool     `json:"quiet"`
	Rate                      int      `json:"rate"`
	ScraperFile               string   `json:"scraperfile"`
	Scrapers                  string   `json:"scrapers"`
	Searchhash                string   `json:"-"`
	ShowVersion               bool     `toml:"-" json:"-"`
	StopOn403                 bool     `json:"stop_on_403"`
	StopOnAll                 bool     `json:"stop_on_all"`
	StopOnErrors              bool     `json:"stop_on_errors"`
	Threads                   int      `json:"threads"`
	Verbose                   bool     `json:"verbose"`
}
type HTTPConfig struct {
	Cookies           []string `json:"cookies"`
	Data              string   `json:"data"`
	FollowRedirect    bool     `json:"follow_redirect"`
	Headers           []string `json:"headers"`
	IgnoreBody        bool     `json:"ignore_body"`
	Method            string   `json:"method"`
	ProxyURL          string   `json:"proxy_url"`
	Raw               bool     `json:"raw"`
	Recursion         bool     `json:"recursion"`
	RecursionDepth    int      `json:"recursion_depth"`
	RecursionStrategy string   `json:"recursion_strategy"`
	ReplayProxyURL    string   `json:"replay_proxy_url"`
	SNI               string   `json:"sni"`
	Timeout           int      `json:"timeout"`
	URL               string   `json:"url"`
	Http2             bool     `json:"http2"`
	ClientCert        string   `json:"client-cert"`
	ClientKey         string   `json:"client-key"`
}
type FilterConfig struct {
	Mode   string `json:"mode"`
	Line   string `json:"line"`
	Regex  string `json:"regex"`
	Size   string `json:"size"`
	Status string `json:"status"`
	Time   string `json:"time"`
	Words  string `json:"words"`
}
type InputConfig struct {
	DirSearchCompat        bool     `json:"dirsearch_compat"`
	Encoders               []string `json:"encoders"`
	Extensions             string   `json:"extensions"`
	IgnoreWordlistComments bool     `json:"ignore_wordlist_comments"`
	InputMode              string   `json:"input_mode"`
	InputNum               int      `json:"input_num"`
	InputShell             string   `json:"input_shell"`
	InputCommands          []string `json:"input_commands"`
	Request                string   `json:"request_file"`
	RequestProto           string   `json:"request_proto"`
	WordList               []string `json:"word_list"`
}
type MatcherConfig struct {
	Mode   string `json:"mode"`
	Line   string `json:"line"`
	Regex  string `json:"regex"`
	Size   string `json:"size"`
	Status string `json:"status"`
	Time   string `json:"time"`
	Words  string `json:"words"`
}
type OutputConfig struct {
	AuditLog            string `json:"audit_log"`
	DebugLog            string `json:"debug_log"`
	OutputDirectory     string `json:"output_directory"`
	OutputFile          string `json:"output_file"`
	OutputFormat        string `json:"output_format"`
	OutputSkipEmptyFile bool   `json:"output_skip_empty"`
}
