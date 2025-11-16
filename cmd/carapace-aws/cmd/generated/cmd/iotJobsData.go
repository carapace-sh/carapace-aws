package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsDataCmd = &cobra.Command{
	Use:   "iot-jobs-data",
	Short: "IoT Jobs is a service that allows you to define a set of jobs — remote operations that are sent to and executed on one or more devices connected to Amazon Web Services IoT Core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotJobsDataCmd).Standalone()

	})
	rootCmd.AddCommand(iotJobsDataCmd)
}
