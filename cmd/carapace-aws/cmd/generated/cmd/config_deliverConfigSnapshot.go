package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deliverConfigSnapshotCmd = &cobra.Command{
	Use:   "deliver-config-snapshot",
	Short: "Schedules delivery of a configuration snapshot to the Amazon S3 bucket in the specified delivery channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deliverConfigSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deliverConfigSnapshotCmd).Standalone()

		config_deliverConfigSnapshotCmd.Flags().String("delivery-channel-name", "", "The name of the delivery channel through which the snapshot is delivered.")
		config_deliverConfigSnapshotCmd.MarkFlagRequired("delivery-channel-name")
	})
	configCmd.AddCommand(config_deliverConfigSnapshotCmd)
}
