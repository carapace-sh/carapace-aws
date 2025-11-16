package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listResourceSnapshotJobsCmd = &cobra.Command{
	Use:   "list-resource-snapshot-jobs",
	Short: "Lists resource snapshot jobs owned by the customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listResourceSnapshotJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listResourceSnapshotJobsCmd).Standalone()

		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("engagement-identifier", "", "The identifier of the engagement to filter the response.")
		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("sort", "", "Configures the sorting of the response.")
		partnercentralSelling_listResourceSnapshotJobsCmd.Flags().String("status", "", "The status of the jobs to filter the response.")
		partnercentralSelling_listResourceSnapshotJobsCmd.MarkFlagRequired("catalog")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listResourceSnapshotJobsCmd)
}
