package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_listApplicationSnapshotsCmd = &cobra.Command{
	Use:   "list-application-snapshots",
	Short: "Lists information about the current application snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_listApplicationSnapshotsCmd).Standalone()

	kinesisanalyticsv2_listApplicationSnapshotsCmd.Flags().String("application-name", "", "The name of an existing application.")
	kinesisanalyticsv2_listApplicationSnapshotsCmd.Flags().String("limit", "", "The maximum number of application snapshots to list.")
	kinesisanalyticsv2_listApplicationSnapshotsCmd.Flags().String("next-token", "", "Use this parameter if you receive a `NextToken` response in a previous request that indicates that there is more output available.")
	kinesisanalyticsv2_listApplicationSnapshotsCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_listApplicationSnapshotsCmd)
}
