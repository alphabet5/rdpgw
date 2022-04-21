package config

import (
	//	"github.com/spf13/viper"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"log"
	"strings"
)

type Configuration struct {
	Server   ServerConfig   `koanf:"server"`
	OpenId   OpenIDConfig   `koanf:"openid"`
	Caps     RDGCapsConfig  `koanf:"caps"`
	Security SecurityConfig `koanf:"security"`
	Client   ClientConfig   `koanf:"client"`
}

type ServerConfig struct {
	GatewayAddress       string   `koanf:"gatewayaddress"`
	Port                 int      `koanf:"port"`
	CertFile             string   `koanf:"certfile"`
	KeyFile              string   `koanf:"keyfile"`
	Hosts                []string `koanf:"hosts"`
	RoundRobin           bool     `koanf:"roundrobin"`
	SessionKey           string   `koanf:"sessionkey"`
	SessionEncryptionKey string   `koanf:"sessionencryptionkey"`
	SendBuf              int      `koanf:"sendbuf"`
	ReceiveBuf           int      `koanf:"recievebuf"`
}

type OpenIDConfig struct {
	ProviderUrl  string `koanf:"providerurl"`
	ClientId     string `koanf:"clientid"`
	ClientSecret string `koanf:"clientsecret"`
}

type RDGCapsConfig struct {
	SmartCardAuth   bool `koanf:"smartcardauth"`
	TokenAuth       bool `koanf:"tokenauth"`
	IdleTimeout     int  `koanf:"idletimeout"`
	RedirectAll     bool `koanf:"redirectall"`
	DisableRedirect bool `koanf:"disableredirect"`
	EnableClipboard bool `koanf:"enableclipboard"`
	EnablePrinter   bool `koanf:"enableprinter"`
	EnablePort      bool `koanf:"enableport"`
	EnablePnp       bool `koanf:"enablepnp"`
	EnableDrive     bool `koanf:"enabledrive"`
}

type SecurityConfig struct {
	PAATokenEncryptionKey  string `koanf:"paatokenencryptionkey"`
	PAATokenSigningKey     string `koanf:"paatokensigningkey"`
	UserTokenEncryptionKey string `koanf:"usertokenencryptionkey"`
	UserTokenSigningKey    string `koanf:"usertokensigningkey"`
	VerifyClientIp         bool   `koanf:"verifyclientip"`
	EnableUserToken        bool   `koanf:"enableusertoken"`
}

type ClientConfig struct {
	NetworkAutoDetect   int    `koanf:"networkautodetect"`
	BandwidthAutoDetect int    `koanf:"bandwidthautodetect"`
	ConnectionType      int    `koanf:"connectiontype"`
	UsernameTemplate    string `koanf:"usernametemplate"`
	SplitUserDomain     bool   `koanf:"splituserdomain"`
	DefaultDomain       string `koanf:"defaultdomain"`
}

func Load(configFile string) Configuration {

	var k = koanf.New(".")

	k.Load(confmap.Provider(map[string]interface{}{
		"Server.CertFile":            "server.pem",
		"Server.KeyFile":             "key.pem",
		"Server.Port":                443,
		"Client.NetworkAutoDetect":   1,
		"Client.BandwidthAutoDetect": 1,
		"Security.VerifyClientIp":    true,
	}, "."), nil)

	if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
		log.Fatalf("Error loading config from file: %v", err)
	}

	var envPrefix string = "RDPGW_"

	if err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, envPrefix)), "_", ".", -1)
	}), nil); err != nil {
		log.Fatalf("Error loading config from env: %v", err)
	}

	var conf Configuration

	var koanfTag = koanf.UnmarshalConf{Tag: "koanf"}

	k.UnmarshalWithConf("Server", &conf.Server, koanfTag)
	k.UnmarshalWithConf("OpenId", &conf.OpenId, koanfTag)
	k.UnmarshalWithConf("Caps", &conf.Caps, koanfTag)
	k.UnmarshalWithConf("Security", &conf.Security, koanfTag)
	k.UnmarshalWithConf("Client", &conf.Client, koanfTag)

	return conf

}
