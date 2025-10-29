package intruder

import (
	"context"
)

type Config struct {
	AuditLog                  string                `json:"auditlog"`
	AutoCalibration           bool                  `json:"autocalibration"`
	AutoCalibrationKeyword    string                `json:"autocalibration_keyword"`
	AutoCalibrationPerHost    bool                  `json:"autocalibration_perhost"`
	AutoCalibrationStrategies []string              `json:"autocalibration_strategies"`
	AutoCalibrationStrings    []string              `json:"autocalibration_strings"`
	Cancel                    context.CancelFunc    `json:"-"`
	Colors                    bool                  `json:"colors"`
	CommandKeywords           []string              `json:"-"`
	CommandLine               string                `json:"cmdline"`
	ConfigFile                string                `json:"configfile"`
	Context                   context.Context       `json:"-"`
	Data                      string                `json:"postdata"`
	Debuglog                  string                `json:"debuglog"`
	Delay                     RangeOpt              `json:"delay"`
	DirSearchCompat           bool                  `json:"dirsearch_compatibility"`
	Encoders                  []string              `json:"encoders"`
	Extensions                []string              `json:"extensions"`
	FilterMode                string                `json:"fmode"`
	FollowRedirects           bool                  `json:"follow_redirects"`
	Headers                   map[string]string     `json:"headers"`
	IgnoreBody                bool                  `json:"ignorebody"`
	IgnoreWordlistComments    bool                  `json:"ignore_wordlist_comments"`
	InputMode                 string                `json:"inputmode"`
	InputNum                  int                   `json:"cmd_inputnum"`
	InputProviders            []InputProviderConfig `json:"inputproviders"`
	InputShell                string                `json:"inputshell"`
	Json                      bool                  `json:"json"`
	MatcherManager            MatcherManager        `json:"matchers"`
	MatcherMode               string                `json:"mmode"`
	MaxTime                   int                   `json:"maxtime"`
	MaxTimeJob                int                   `json:"maxtime_job"`
	Method                    string                `json:"method"`
	Noninteractive            bool                  `json:"noninteractive"`
	OutputDirectory           string                `json:"outputdirectory"`
	OutputFile                string                `json:"outputfile"`
	OutputFormat              string                `json:"outputformat"`
	OutputSkipEmptyFile       bool                  `json:"OutputSkipEmptyFile"`
	ProgressFrequency         int                   `json:"-"`
	ProxyURL                  string                `json:"proxyurl"`
	Quiet                     bool                  `json:"quiet"`
	Rate                      int64                 `json:"rate"`
	Raw                       bool                  `json:"raw"`
	Recursion                 bool                  `json:"recursion"`
	RecursionDepth            int                   `json:"recursion_depth"`
	RecursionStrategy         string                `json:"recursion_strategy"`
	ReplayProxyURL            string                `json:"replayproxyurl"`
	RequestFile               string                `json:"requestfile"`
	RequestProto              string                `json:"requestproto"`
	ScraperFile               string                `json:"scraperfile"`
	Scrapers                  string                `json:"scrapers"`
	SNI                       string                `json:"sni"`
	StopOn403                 bool                  `json:"stop_403"`
	StopOnAll                 bool                  `json:"stop_all"`
	StopOnErrors              bool                  `json:"stop_errors"`
	Threads                   int                   `json:"threads"`
	Timeout                   int                   `json:"timeout"`
	Url                       string                `json:"url"`
	Verbose                   bool                  `json:"verbose"`
	WordList                  []string              `json:"word_list"`
	Http2                     bool                  `json:"http2"`
	ClientCert                string                `json:"client-cert"`
	ClientKey                 string                `json:"client-key"`
}

type InputProviderConfig struct {
	Name     string `json:"name"`
	Keyword  string `json:"keyword"`
	Value    string `json:"value"`
	Encoders string `json:"encoders"`
	Template string `json:"template"` // the templating string used for sniper mode (usually "§")
}

func NewConfig(ctx context.Context, cancelFunc context.CancelFunc) Config {
	var config Config
	config.AutoCalibrationKeyword = "FUZZ"
	config.AutoCalibrationStrategies = []string{"basic"}
	config.AutoCalibrationStrings = make([]string, 0)
	config.CommandKeywords = make([]string, 0)
	config.Context = ctx
	config.Cancel = cancelFunc
	config.Data = ""
	config.Debuglog = ""
	config.Delay = config.RangeOpt{0, 0, false, false}
	config.DirSearchCompat = false
	config.Encoders = make([]string, 0)
	config.Extensions = make([]string, 0)
	config.FilterMode = "or"
	config.FollowRedirects = false
	config.Headers = make(map[string]string)
	config.IgnoreWordlistComments = false
	config.InputMode = "clusterbomb"
	config.InputNum = 0
	config.InputShell = ""
	config.InputProviders = make([]InputProviderConfig, 0)
	config.Json = false
	config.MatcherMode = "or"
	config.MaxTime = 0
	config.MaxTimeJob = 0
	config.Method = "GET"
	config.Noninteractive = false
	config.ProgressFrequency = 125
	config.ProxyURL = ""
	config.Quiet = false
	config.Rate = 0
	config.Raw = false
	config.Recursion = false
	config.RecursionDepth = 0
	config.RecursionStrategy = "default"
	config.RequestFile = ""
	config.RequestProto = "https"
	config.SNI = ""
	config.ScraperFile = ""
	config.Scrapers = "all"
	config.StopOn403 = false
	config.StopOnAll = false
	config.StopOnErrors = false
	config.Timeout = 10
	config.Url = ""
	config.Verbose = false
	config.WordList = []string{}
	config.Http2 = true
	return config
}

func (c *Config) SetContext(ctx context.Context, cancel context.CancelFunc) {
	c.Context = ctx
	c.Cancel = cancel
}
