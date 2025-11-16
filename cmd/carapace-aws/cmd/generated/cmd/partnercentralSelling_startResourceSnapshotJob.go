package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_startResourceSnapshotJobCmd = &cobra.Command{
	Use:   "start-resource-snapshot-job",
	Short: "Starts a resource snapshot job that has been previously created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_startResourceSnapshotJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_startResourceSnapshotJobCmd).Standalone()

		partnercentralSelling_startResourceSnapshotJobCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
		partnercentralSelling_startResourceSnapshotJobCmd.Flags().String("resource-snapshot-job-identifier", "", "The identifier of the resource snapshot job to start.")
		partnercentralSelling_startResourceSnapshotJobCmd.MarkFlagRequired("catalog")
		partnercentralSelling_startResourceSnapshotJobCmd.MarkFlagRequired("resource-snapshot-job-identifier")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_startResourceSnapshotJobCmd)
}
