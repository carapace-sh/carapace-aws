package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_deleteResourceSnapshotJobCmd = &cobra.Command{
	Use:   "delete-resource-snapshot-job",
	Short: "Use this action to deletes a previously created resource snapshot job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_deleteResourceSnapshotJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_deleteResourceSnapshotJobCmd).Standalone()

		partnercentralSelling_deleteResourceSnapshotJobCmd.Flags().String("catalog", "", "Specifies the catalog from which to delete the snapshot job.")
		partnercentralSelling_deleteResourceSnapshotJobCmd.Flags().String("resource-snapshot-job-identifier", "", "The unique identifier of the resource snapshot job to be deleted.")
		partnercentralSelling_deleteResourceSnapshotJobCmd.MarkFlagRequired("catalog")
		partnercentralSelling_deleteResourceSnapshotJobCmd.MarkFlagRequired("resource-snapshot-job-identifier")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_deleteResourceSnapshotJobCmd)
}
