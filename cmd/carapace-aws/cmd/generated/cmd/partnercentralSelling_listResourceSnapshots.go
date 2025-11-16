package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listResourceSnapshotsCmd = &cobra.Command{
	Use:   "list-resource-snapshots",
	Short: "Retrieves a list of resource view snapshots based on specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listResourceSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listResourceSnapshotsCmd).Standalone()

		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("created-by", "", "Filters the response to include only snapshots of resources owned by the specified AWS account.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("engagement-identifier", "", "The unique identifier of the engagement associated with the snapshots.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("resource-identifier", "", "Filters the response to include only snapshots of the specified resource.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("resource-snapshot-template-identifier", "", "Filters the response to include only snapshots created using the specified template.")
		partnercentralSelling_listResourceSnapshotsCmd.Flags().String("resource-type", "", "Filters the response to include only snapshots of the specified resource type.")
		partnercentralSelling_listResourceSnapshotsCmd.MarkFlagRequired("catalog")
		partnercentralSelling_listResourceSnapshotsCmd.MarkFlagRequired("engagement-identifier")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listResourceSnapshotsCmd)
}
