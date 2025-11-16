package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_putLoggingOptionsCmd = &cobra.Command{
	Use:   "put-logging-options",
	Short: "Creates or updates the logging option.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_putLoggingOptionsCmd).Standalone()

	iotfleetwise_putLoggingOptionsCmd.Flags().String("cloud-watch-log-delivery", "", "Creates or updates the log delivery option to Amazon CloudWatch Logs.")
	iotfleetwise_putLoggingOptionsCmd.MarkFlagRequired("cloud-watch-log-delivery")
	iotfleetwiseCmd.AddCommand(iotfleetwise_putLoggingOptionsCmd)
}
