package main

import ()

// ----------------------------------------------------------------------------
// Configuration keys (for use with Viper)
// ----------------------------------------------------------------------------

// These keys match what is in the JSON/YAML configuration file
const (
	ETCD_CONFIG_ADVERTISE_CLIENT_URLS       = "advertiseClientUrls"
	ETCD_CONFIG_DEBUG                       = "debug"
	ETCD_CONFIG_INITIAL_ADVERTISE_PEER_URLS = "initialAdvertisePeerUrls"
	ETCD_CONFIG_INITIAL_CLUSTER             = "initialCluster"
	ETCD_CONFIG_INITIAL_CLUSTER_STATE       = "initialClusterState"
	ETCD_CONFIG_INITIAL_CLUSTER_TOKEN       = "initialClusterToken"
	ETCD_CONFIG_LISTEN_CLIENT_URLS          = "listenClientUrls"
	ETCD_CONFIG_LISTEN_PEER_URLS            = "listenPeerUrls"
	ETCD_CONFIG_NAME                        = "name"
	ETCD_CONFIG_TRACE                       = "trace"
)

// List of keys used in context.Values(...)
func GetConfigKeys() []string {
	return []string{
		ETCD_CONFIG_ADVERTISE_CLIENT_URLS,
		ETCD_CONFIG_DEBUG,
		ETCD_CONFIG_INITIAL_ADVERTISE_PEER_URLS,
		ETCD_CONFIG_INITIAL_CLUSTER,
		ETCD_CONFIG_INITIAL_CLUSTER_STATE,
		ETCD_CONFIG_INITIAL_CLUSTER_TOKEN,
		ETCD_CONFIG_LISTEN_CLIENT_URLS,
		ETCD_CONFIG_LISTEN_PEER_URLS,
		ETCD_CONFIG_NAME,
		ETCD_CONFIG_TRACE,
	}
}

// ----------------------------------------------------------------------------
// Commandline options
// ----------------------------------------------------------------------------

const (
	ETCD_OPTION_ADVERTISE_CLIENT_URLS       = "--advertise-client-urls"
	ETCD_OPTION_CONFIGURATION               = "--configuration"
	ETCD_OPTION_DEBUG                       = "--debug"
	ETCD_OPTION_INITIAL_ADVERTISE_PEER_URLS = "--initial-advertise-peer-urls"
	ETCD_OPTION_INITIAL_CLUSTER             = "--initial-cluster"
	ETCD_OPTION_INITIAL_CLUSTER_STATE       = "--initial-cluster-state"
	ETCD_OPTION_INITIAL_CLUSTER_TOKEN       = "--initial-cluster-token"
	ETCD_OPTION_LISTEN_CLIENT_URLS          = "--listen-client-urls"
	ETCD_OPTION_LISTEN_PEER_URLS            = "--listen-peer-urls"
	ETCD_OPTION_NAME                        = "--name"
	ETCD_OPTION_TRACE                       = "--trace"
)

// ----------------------------------------------------------------------------
// Operating System Environment variables
// ----------------------------------------------------------------------------

const (
	ETCD_OS_ADVERTISE_CLIENT_URLS       = "ETCD_ADVERTISE_CLIENT_URLS"
	ETCD_OS_DEBUG                       = "ETCD_DEBUG"
	ETCD_OS_INITIAL_ADVERTISE_PEER_URLS = "ETCD_INITIAL_ADVERTISE_PEER_URLS"
	ETCD_OS_INITIAL_CLUSTER             = "ETCD_INITIAL_CLUSTER"
	ETCD_OS_INITIAL_CLUSTER_STATE       = "ETCD_INITIAL_CLUSTER_STATE"
	ETCD_OS_INITIAL_CLUSTER_TOKEN       = "ETCD_INITIAL_CLUSTER_TOKEN"
	ETCD_OS_LISTEN_CLIENT_URLS          = "ETCD_LISTEN_CLIENT_URLS"
	ETCD_OS_LISTEN_PEER_URLS            = "ETCD_LISTEN_PEER_URLS"
	ETCD_OS_NAME                        = "ETCD_NAME"
	ETCD_OS_TRACE                       = "ETCD_TRACE"
)

// List of OS Environment variable used by BixAgent.
func GetOsEnvKeys() []string {
	return []string{
		ETCD_OS_ADVERTISE_CLIENT_URLS,
		ETCD_OS_DEBUG,
		ETCD_OS_INITIAL_ADVERTISE_PEER_URLS,
		ETCD_OS_INITIAL_CLUSTER,
		ETCD_OS_INITIAL_CLUSTER_STATE,
		ETCD_OS_INITIAL_CLUSTER_TOKEN,
		ETCD_OS_LISTEN_CLIENT_URLS,
		ETCD_OS_LISTEN_PEER_URLS,
		ETCD_OS_NAME,
		ETCD_OS_TRACE,
	}
}
