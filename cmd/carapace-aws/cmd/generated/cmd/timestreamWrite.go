package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWriteCmd = &cobra.Command{
	Use:   "timestream-write",
	Short: "Amazon Timestream Write\n\nAmazon Timestream is a fast, scalable, fully managed time-series database service that makes it easy to store and analyze trillions of time-series data points per day.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWriteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWriteCmd).Standalone()

	})
	rootCmd.AddCommand(timestreamWriteCmd)
}
