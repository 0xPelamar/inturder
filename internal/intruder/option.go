package intruder

import (
	"context"
	"fmt"
	"github.com/0xPelamar/intruder/internal/helper"
	"runtime"
	"strings"
)

type Options struct {
	General GeneralOptions `json:"general"`
	HTTP    HTTPOptions    `json:"http"`
	Filter  FilterOptions  `json:"filter"`
	Input   InputOptions   `json:"input"`
	Matcher MatcherOptions `json:"matcher"`
	Output  OutputOptions  `json:"output"`
}
type GeneralOptions struct {
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
type HTTPOptions struct {
	Cookies           []string `json:"-"`
	Data              string   `json:"data"`
	FollowRedirects   bool     `json:"follow_redirects"`
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
	HTTP2             bool     `json:"http2"`
	ClientCert        string   `json:"client-cert"`
	ClientKey         string   `json:"client-key"`
}
type FilterOptions struct {
	Mode   string `json:"mode"`
	Lines  string `json:"lines"`
	Regexp string `json:"regexp"`
	Size   string `json:"size"`
	Status string `json:"status"`
	Time   string `json:"time"`
	Words  string `json:"words"`
}
type InputOptions struct {
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
type MatcherOptions struct {
	Mode   string `json:"mode"`
	Lines  string `json:"lines"`
	Regexp string `json:"regexp"`
	Size   string `json:"size"`
	Status string `json:"status"`
	Time   string `json:"time"`
	Words  string `json:"words"`
}
type OutputOptions struct {
	AuditLog            string `json:"audit_log"`
	DebugLog            string `json:"debug_log"`
	OutputDirectory     string `json:"output_directory"`
	OutputFile          string `json:"output_file"`
	OutputFormat        string `json:"output_format"`
	OutputSkipEmptyFile bool   `json:"output_skip_empty"`
}

func NewOptions() *Options {
	o := &Options{}
	o.Filter.Mode = "or"
	o.Filter.Lines = ""
	o.Filter.Regexp = ""
	o.Filter.Size = ""
	o.Filter.Status = ""
	o.Filter.Time = ""
	o.Filter.Words = ""

	o.General.AutoCalibration = false
	o.General.AutoCalibrationKeyword = "FUZZ"
	o.General.AutoCalibrationStrategies = []string{"basic"}
	o.General.Colors = false
	o.General.Delay = ""
	o.General.Json = false
	o.General.MaxTime = 0
	o.General.MaxTimeJob = 0
	o.General.Noninteractive = false
	o.General.Quiet = false
	o.General.Rate = 0
	o.General.Searchhash = ""
	o.General.ScraperFile = ""
	o.General.Scrapers = "all"
	o.General.ShowVersion = false
	o.General.StopOn403 = false
	o.General.StopOnAll = false
	o.General.StopOnErrors = false
	o.General.Threads = 40
	o.General.Verbose = false

	o.HTTP.Data = ""
	o.HTTP.FollowRedirects = false
	o.HTTP.IgnoreBody = false
	o.HTTP.Method = ""
	o.HTTP.ProxyURL = ""
	o.HTTP.Raw = false
	o.HTTP.Recursion = false
	o.HTTP.RecursionDepth = 0
	o.HTTP.RecursionStrategy = "default"
	o.HTTP.ReplayProxyURL = ""
	o.HTTP.Timeout = 10
	o.HTTP.SNI = ""
	o.HTTP.URL = ""
	o.HTTP.HTTP2 = false

	o.Input.DirSearchCompat = false
	o.Input.Encoders = []string{}
	o.Input.Extensions = ""
	o.Input.IgnoreWordlistComments = false
	o.Input.InputMode = "clusterbomb"
	o.Input.InputNum = 100
	o.Input.Request = ""
	o.Input.RequestProto = "https"
	o.Input.WordList = []string{}
	o.Input.InputCommands = []string{}
	o.Input.InputShell = ""

	o.Matcher.Mode = "or"
	o.Matcher.Lines = ""
	o.Matcher.Regexp = ""
	o.Matcher.Size = ""
	o.Matcher.Status = "200-299,301,302,307,401,403,405,500"
	o.Matcher.Time = ""
	o.Matcher.Words = ""

	o.Output.AuditLog = ""
	o.Output.DebugLog = ""
	o.Output.OutputDirectory = ""
	o.Output.OutputFile = ""
	o.Output.OutputFormat = "json"
	o.Output.OutputSkipEmptyFile = false

	return o
}

func GetConfigFromOptions(opt *Options, ctx context.Context, cancelFunc context.CancelFunc) (*Config, error) {
	config := NewConfig(ctx, cancelFunc)
	errs := helper.NewMultiError()
	if len(opt.HTTP.URL) == 0 && len(opt.Input.Request) == 0 {
		errs.Add(fmt.Errorf("-u flag or -request flag is required"))
	}

	// Prepare extensions
	if opt.Input.Extensions != "" {
		extensions := strings.Split(opt.Input.Extensions, ",")
		config.Extensions = extensions
	}

	if len(opt.HTTP.Cookies) > 0 {
		opt.HTTP.Headers = append(opt.HTTP.Headers, fmt.Sprintf("Cookie: %s", strings.Join(opt.HTTP.Cookies, ";")))
	}

	config.InputMode = opt.Input.InputMode

	var validMode bool
	for _, mode := range []string{"clusterbomb", "pitchfork", "sniper"} {
		if config.InputMode == mode {
			validMode = true
		}
	}
	if !validMode {
		errs.Add(fmt.Errorf("input mode (-mode) %s not recognized", config.InputMode))
	}

	var template string
	// sniper mode needs some additional checking
	if config.InputMode == "sniper" {
		template = "§"

		if len(opt.Input.WordList) > 1 {
			errs.Add(fmt.Errorf("sniper mode only supports one wordlist"))
		}

		if len(opt.Input.InputCommands) > 1 {
			errs.Add(fmt.Errorf("sniper mode only supports one input command"))
		}
	}
	tempEncoders := make(map[string]string)
	for _, e := range opt.Input.Encoders {
		if strings.Contains(e, ":") {
			key := strings.Split(e, ":")[0]
			val := strings.Split(e, ":")[1]
			tempEncoders[key] = val
		}
	}
	tempWordlist := make([]string, 0)
	for _, word := range opt.Input.WordList {
		var wl []string
		if runtime.GOOS == "windows" {
			// Try to ensure that Windows file paths like C:\path\to\wordlist.txt:KEYWORD are treated properly
			if helper.FileExists(word) {
				// The wordlist was supplied without a keyword parameter
				wl = []string{word}
			} else {
				filepart := word
				if strings.Contains(filepart, ":") {
					filepart = word[:strings.LastIndex(filepart, ":")]
				}

				if helper.FileExists(filepart) {
					wl = []string{filepart, word[strings.LastIndex(word, ":")+1:]}
				} else {
					// The file was not found. Use full wordlist parameter value for more concise error message down the line
					wl = []string{word}
				}

			}
		} else {
			wl = strings.SplitN(word, ":", 2)
		}
	}

}
