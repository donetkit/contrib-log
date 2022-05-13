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

// WithHostInfo set hostName, ip function
func WithHostInfo(hostName, ip string) Option {
	return func(cfg *Config) {
		cfg.hostName = hostName
		cfg.ip = ip
	}
}

// WithLogLevel set logLevel function
func WithLogLevel(logLevel LogLevel) Option {
	return func(cfg *Config) {
		cfg.logLevel = logLevel
	}
}

// WithLogColor set logColor function
func WithLogColor(logColor bool) Option {
	return func(cfg *Config) {
		cfg.logColor = logColor
	}
}
