package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/coreos/etcd/embed"
	"github.com/docktermj/go-logger/logger"
	"github.com/docopt/docopt-go"
	"github.com/spf13/viper"
)

// Values updated via "go install -ldflags" parameters.

var (
	programName    string = "go-hello-etcd"
	buildVersion   string = "0.0.0"
	buildIteration string = "0"
)

// Given a string like "http://10.1.1.1,http://10.2.2.2" create a list of []url.URL
func createUrlList(urlListString string) []url.URL {
	var result []url.URL
	if len(urlListString) > 0 {
		splits := strings.Split(urlListString, ",")
		for _, urlString := range splits {
			aUrl, err := url.Parse(urlString)
			if err == nil {
				result = append(result, *aUrl)
			} else {
				logger.Warnf("Could not parse URL: %s Reason: %v", urlString, err)
			}
		}
	}
	return result
}

// Get configuration values from environment
// https://github.com/coreos/etcd/blob/master/embed/config.go
func getEtcdConfig() *embed.Config {
	result := embed.NewConfig()
	result.Dir = "default.etcd"

	// Load URL lists.

	result.ACUrls = createUrlList(viper.GetString(ETCD_CONFIG_ADVERTISE_CLIENT_URLS))
	result.APUrls = createUrlList(viper.GetString(ETCD_CONFIG_INITIAL_ADVERTISE_PEER_URLS))
	result.LCUrls = createUrlList(viper.GetString(ETCD_CONFIG_LISTEN_CLIENT_URLS))
	result.LPUrls = createUrlList(viper.GetString(ETCD_CONFIG_LISTEN_PEER_URLS))

	value := viper.GetString(ETCD_CONFIG_INITIAL_CLUSTER)
	if len(value) > 0 {
		result.InitialCluster = value
	}

	value = viper.GetString(ETCD_CONFIG_INITIAL_CLUSTER_STATE)
	if len(value) > 0 {
		result.ClusterState = value
	}

	value = viper.GetString(ETCD_CONFIG_INITIAL_CLUSTER_TOKEN)
	if len(value) > 0 {
		result.InitialClusterToken = value
	}

	value = viper.GetString(ETCD_CONFIG_NAME)
	if len(value) > 0 {
		result.Name = value
	}

	logger.Infof("%+v\n", result)

	return result
}

func main() {

	usage := `
Usage:
    go-hello-etcd [<command>] [options]

Options:
   -h, --help                           Show this help
   --name=<id>                          xxx
   --advertise-client-urls=<urls>       xxx
   --listen-client-urls=<urls>          xxx
   --initial-advertise-peer-urls=<urls> xxx
   --listen-peer-urls=<urls>            xxx
   --initial-cluster-token=<token>      xxx
   --initial-cluster=<namedurls>        xxx
   --initial-cluster-state=<state>      xxx   
`

	// DocOpt processing.

	commandVersion := fmt.Sprintf("%s %s-%s", programName, buildVersion, buildIteration)
	args, _ := docopt.Parse(usage, nil, true, commandVersion, false)

	// Configure output log.

	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	logger.SetLevel(logger.LevelInfo)

	// Load configuration from commandline, OS environment variables, configuration file, defaults.

	LoadConfig(args)
	if viper.GetBool(ETCD_CONFIG_DEBUG) {
		logger.SetLevel(logger.LevelDebug)
	}
	if viper.GetBool(ETCD_CONFIG_TRACE) {
		logger.SetLevel(logger.LevelTrace)
	}

	// Configuration
	// https://github.com/coreos/etcd/blob/master/embed/config.go

	etcdConfig := getEtcdConfig()

	// Start Etcd server.

	etcdService, err := embed.StartEtcd(etcdConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdService.Close()

	// Monitor etcd service.

	select {
	case <-etcdService.Server.ReadyNotify():
		log.Printf("Server is ready!")
	case <-time.After(60 * time.Second):
		etcdService.Server.Stop() // trigger a shutdown
		log.Printf("Server took too long to start!")
	}
	log.Fatal(<-etcdService.Err())
}
