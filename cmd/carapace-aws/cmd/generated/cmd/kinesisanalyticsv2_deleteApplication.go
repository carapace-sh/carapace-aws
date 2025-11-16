package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_deleteApplicationCmd).Standalone()

		kinesisanalyticsv2_deleteApplicationCmd.Flags().String("application-name", "", "The name of the application to delete.")
		kinesisanalyticsv2_deleteApplicationCmd.Flags().String("create-timestamp", "", "Use the `DescribeApplication` operation to get this value.")
		kinesisanalyticsv2_deleteApplicationCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_deleteApplicationCmd.MarkFlagRequired("create-timestamp")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationCmd)
}
