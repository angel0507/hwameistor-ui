package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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

func RegisterShutdownChannel(done chan struct{}) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Infof("Receive %v to exit", sig)
		close(done)
	}()
}
