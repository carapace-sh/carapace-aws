package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_startApplicationCmd = &cobra.Command{
	Use:   "start-application",
	Short: "Starts the specified Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_startApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_startApplicationCmd).Standalone()

		kinesisanalyticsv2_startApplicationCmd.Flags().String("application-name", "", "The name of the application.")
		kinesisanalyticsv2_startApplicationCmd.Flags().String("run-configuration", "", "Identifies the run configuration (start parameters) of a Managed Service for Apache Flink application.")
		kinesisanalyticsv2_startApplicationCmd.MarkFlagRequired("application-name")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_startApplicationCmd)
}
