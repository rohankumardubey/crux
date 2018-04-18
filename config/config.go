package config

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"fmt"
)

const (
	Verbosity = "verbosity"
	AlwaysSendTo = "alwayssendto"
	Storage = "storage"
	WorkDir = "workdir"
	Url = "url"
	OtherNodes = "othernodes"
	PublicKeys = "publickeys"
	PrivateKeys = "privatekeys"
	Port = "port"
	Socket = "socket"

	GenerateKeys = "generate-keys"

	BerkeleyDb = "berkeleydb"

	Tls = "tls"
	TlsServerChain = "tlsserverchain"
	TlsServerTrust = "tlsservertrust"
	TlsKnownServers = "tlsknownservers"
	TlsClientCert = "tlsclientcert"
	TlsServerCert = "tlsservercert"
	TlsKnownClients = "tlsknownclients"
	TlsClientChain = "tlsclientchain"
	TlsClientKey = "tlsclientkey"
	TlsClientTrust = "tlsclienttrust"
	TlsServerKey = "tlsserverkey"
)

func InitFlags() {
	flag.String(GenerateKeys, "", "Generate a new keypair")
	flag.String(Url, "", "The URL to advertise to other nodes (reachable by them)")
	flag.Int(Port, -1, "The local port to listen on")
	flag.String(WorkDir, ".", "The folder to put stuff in (default: .)")
	flag.String(Socket, "crux.ipc", "IPC socket to create for access to the Private API")
	flag.String(OtherNodes, "", "\"Boot nodes\" to connect to to discover the network")
	flag.String(PublicKeys, "", "Public keys hosted by this node")
	flag.String(PrivateKeys, "", "Private keys hosted by this node")
	flag.String(Storage, "crux.db", "Database storage file name")
	flag.Bool(BerkeleyDb, false,
		"Use Berkeley DB for working with an existing Constellation data store [experimental]")

	flag.Int(Verbosity, 1, "Verbosity level of logs")
	flag.String(AlwaysSendTo, "", "List of public keys for nodes to send all transactions too")

	// storage not currently supported as we use LevelDB
	// TLS is not currently supported

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "      %-25s%s\n", "crux.config", "Optional config file")
	pflag.PrintDefaults()
}

func ParseCommandLine() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func LoadConfig(configPath string) error {
	viper.SetConfigType("hcl")
	viper.SetConfigFile(configPath)
	return viper.ReadInConfig()
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}


