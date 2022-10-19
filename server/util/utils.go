package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

// disk class
const (
	DiskClassNameHDD  = "HDD"
	DiskClassNameSSD  = "SSD"
	DiskClassNameNVMe = "NVMe"
)

// consts
const (
	PoolNamePrefix  = "LocalStorage_Pool"
	PoolNameForHDD  = PoolNamePrefix + DiskClassNameHDD
	PoolNameForSSD  = PoolNamePrefix + DiskClassNameSSD
	PoolNameForNVMe = PoolNamePrefix + DiskClassNameNVMe
)

// GetNodeName gets the node name from env, else
// returns an error
func GetNodeName() string {
	nodeName, ok := os.LookupEnv("NODENAME")
	if !ok {
		log.Errorf("Failed to get NODENAME from ENV")
		return ""
	}

	return nodeName
}

// GetNamespace get Namespace from env, else it returns error
func GetNamespace() string {
	ns, ok := os.LookupEnv("NAMESPACE")
	if !ok {
		log.Errorf("Failed to get NameSpace from ENV")
		return ""
	}

	return ns
}

func GetRequiredEnv(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", fmt.Errorf("can't get required environment variable, env %v wasn't set", key)
	}
	return env, nil
}
