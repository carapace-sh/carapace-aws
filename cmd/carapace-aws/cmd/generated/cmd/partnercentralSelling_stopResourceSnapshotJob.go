package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_stopResourceSnapshotJobCmd = &cobra.Command{
	Use:   "stop-resource-snapshot-job",
	Short: "Stops a resource snapshot job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_stopResourceSnapshotJobCmd).Standalone()

	partnercentralSelling_stopResourceSnapshotJobCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
	partnercentralSelling_stopResourceSnapshotJobCmd.Flags().String("resource-snapshot-job-identifier", "", "The identifier of the job to stop.")
	partnercentralSelling_stopResourceSnapshotJobCmd.MarkFlagRequired("catalog")
	partnercentralSelling_stopResourceSnapshotJobCmd.MarkFlagRequired("resource-snapshot-job-identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_stopResourceSnapshotJobCmd)
}
