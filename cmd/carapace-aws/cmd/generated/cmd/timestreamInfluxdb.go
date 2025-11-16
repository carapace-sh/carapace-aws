package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdbCmd = &cobra.Command{
	Use:   "timestream-influxdb",
	Short: "Amazon Timestream for InfluxDB is a managed time-series database engine that makes it easy for application developers and DevOps teams to run InfluxDB databases on Amazon Web Services for near real-time time-series applications using open-source APIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdbCmd).Standalone()

	rootCmd.AddCommand(timestreamInfluxdbCmd)
}
