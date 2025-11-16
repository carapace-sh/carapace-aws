package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_getSnapshotsCmd = &cobra.Command{
	Use:   "get-snapshots",
	Short: "Retrieves search metrics data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_getSnapshotsCmd).Standalone()

	kendra_getSnapshotsCmd.Flags().String("index-id", "", "The identifier of the index to get search metrics data.")
	kendra_getSnapshotsCmd.Flags().String("interval", "", "The time interval or time window to get search metrics data.")
	kendra_getSnapshotsCmd.Flags().String("max-results", "", "The maximum number of returned data for the metric.")
	kendra_getSnapshotsCmd.Flags().String("metric-type", "", "The metric you want to retrieve.")
	kendra_getSnapshotsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_getSnapshotsCmd.MarkFlagRequired("index-id")
	kendra_getSnapshotsCmd.MarkFlagRequired("interval")
	kendra_getSnapshotsCmd.MarkFlagRequired("metric-type")
	kendraCmd.AddCommand(kendra_getSnapshotsCmd)
}
