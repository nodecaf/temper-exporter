package cmd

import (
	"fmt"
	"log"
	"net/http"

	temper "github.com/nodecaf/temper-exporter/pkg/temper"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Starts bbusters client or api server",
	Long:  "To start the client or api server, used keyword client/apiserver",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Main command was called")

		reg := prometheus.NewRegistry()

		// Create new metrics and register them using the custom registry.
		temper.NewMetrics(reg)
		go temper.Checker(hdev)
		//m.Visits.With(prometheus.Labels{"handle": "test"}).Inc()
		log.Println("Starting Prometheus /metrics endpoint")
		// Expose metrics and custom registry via an HTTP server
		// using the HandleFor function. "/metrics" is the usual endpoint for that.
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
		log.Fatal(http.ListenAndServe(":9100", nil))
	},
}

var hdev string

func init() {
	rootCmd.Flags().StringVarP(&hdev, "hidraw", "d", "1", "hidraw device")
}

func Execute() {
	rootCmd.Execute()
}
