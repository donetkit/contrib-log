package glog

import (
	"net"
	"strings"
)

// GetHostIp ip
func getHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err == nil {
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		ip := strings.Split(localAddr.String(), ":")[0]
		return ip
	}
	return ""
}

// Option for queue system
type Option func(*Config)

// WithFile set file function
func WithFile(log2File bool) Option {
	return func(cfg *Config) {
		cfg.log2File = log2File
	}
}

// WithHostName set hostName, ip function
func WithHostName(hostName string) Option {
	return func(cfg *Config) {
		cfg.hostName = hostName
	}
}

// WithHostIP set hostName, ip function
func WithHostIP(ip string) Option {
	return func(cfg *Config) {
		cfg.ip = ip
	}
}

// WithLevel set level function
func WithLevel(level Level) Option {
	return func(cfg *Config) {
		cfg.level = level
	}
}

// WithNoFieldsColors set noFieldsColors function
func WithNoFieldsColors(noFieldsColors bool) Option {
	return func(cfg *Config) {
		cfg.noFieldsColors = noFieldsColors
	}
}

// WithNoColors set noColors function
func WithNoColors(noColors bool) Option {
	return func(cfg *Config) {
		cfg.noColors = noColors
	}
}

// WithClassName set className function
func WithClassName(className string) Option {
	return func(cfg *Config) {
		cfg.className = className
	}
}

// WithFields set fields function
func WithFields(fields map[string]interface{}) Option {
	return func(cfg *Config) {
		cfg.fields = fields
	}
}

// WithDisplayFields set displayFields function
func WithDisplayFields(displayFields bool) Option {
	return func(cfg *Config) {
		cfg.displayFields = displayFields
	}
}

// WithReportCaller set reportCaller function
func WithReportCaller(reportCaller bool) Option {
	return func(cfg *Config) {
		cfg.reportCaller = reportCaller
	}
}

// WithFileName set fileName function
func WithFileName(fileName string) Option {
	return func(cfg *Config) {
		cfg.fileName = fileName
	}
}

// WithMaxSize set maxSize function
func WithMaxSize(maxSize int) Option {
	return func(cfg *Config) {
		cfg.maxSize = maxSize
	}
}

// WithMaxBackups set maxBackups function
func WithMaxBackups(maxBackups int) Option {
	return func(cfg *Config) {
		cfg.maxBackups = maxBackups
	}
}

// WithMaxAge set maxAge function
func WithMaxAge(maxAge int) Option {
	return func(cfg *Config) {
		cfg.maxAge = maxAge
	}
}

// WithCompress set compress function
func WithCompress(compress bool) Option {
	return func(cfg *Config) {
		cfg.compress = compress
	}
}

// WithLocalTime set localTime function
func WithLocalTime(localTime bool) Option {
	return func(cfg *Config) {
		cfg.localTime = localTime
	}
}
