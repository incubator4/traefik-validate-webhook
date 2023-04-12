package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

var parameters WhSvrParameters
var server ServerAddr

func configInit() {
	tc := viper.New()
	tc.SetEnvPrefix("traefik")
	tc.BindEnv("protocol")
	tc.BindEnv("host")
	tc.BindEnv("port")
	tc.AutomaticEnv()

	tc.SetDefault("protocol", "http")

	pflag.StringVar(&server.Protocol, "traefik.protocol", tc.GetString("protocol"), "traefik api protocol")
	pflag.StringVar(&server.Host, "traefik.host", tc.GetString("host"), "traefik api host")
	pflag.IntVar(&server.Port, "traefik.port", tc.GetInt("port"), "traefik api port")

	// get command line parameters
	pflag.IntVar(&parameters.port, "port", 443, "Webhook server port.")
	pflag.StringVar(&parameters.certFile, "tlsCertFile", "/etc/webhook/certs/cert.pem", "File containing the x509 Certificate for HTTPS.")
	pflag.StringVar(&parameters.keyFile, "tlsKeyFile", "/etc/webhook/certs/key.pem", "File containing the x509 private key to --tlsCertFile.")
	pflag.Parse()
}

func main() {
	configInit()

	pair, err := tls.LoadX509KeyPair(parameters.certFile, parameters.keyFile)
	if err != nil {
		glog.Errorf("Failed to load key pair: %v", err)
	}

	whsvr := &WebhookServer{
		server: &http.Server{
			Addr:      fmt.Sprintf(":%v", parameters.port),
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}},
		},
	}

	// define http server and server handler
	mux := http.NewServeMux()
	//mux.HandleFunc("/mutate", whsvr.serve)
	mux.HandleFunc("/validate", whsvr.serve)
	whsvr.server.Handler = mux

	// start webhook server in new routine
	go func() {
		if err := whsvr.server.ListenAndServeTLS("", ""); err != nil {
			glog.Errorf("Failed to listen and serve webhook server: %v", err)
		}
	}()

	glog.Info("Server started")

	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	glog.Infof("Got OS shutdown signal, shutting down webhook server gracefully...")
	whsvr.server.Shutdown(context.Background())
}
