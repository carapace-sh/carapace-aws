package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getResourceSnapshotJobCmd = &cobra.Command{
	Use:   "get-resource-snapshot-job",
	Short: "Use this action to retrieves information about a specific resource snapshot job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getResourceSnapshotJobCmd).Standalone()

	partnercentralSelling_getResourceSnapshotJobCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
	partnercentralSelling_getResourceSnapshotJobCmd.Flags().String("resource-snapshot-job-identifier", "", "The unique identifier of the resource snapshot job to be retrieved.")
	partnercentralSelling_getResourceSnapshotJobCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getResourceSnapshotJobCmd.MarkFlagRequired("resource-snapshot-job-identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getResourceSnapshotJobCmd)
}
