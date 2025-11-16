package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listJobsByConsumableResourceCmd = &cobra.Command{
	Use:   "list-jobs-by-consumable-resource",
	Short: "Returns a list of Batch jobs that require a specific consumable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listJobsByConsumableResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_listJobsByConsumableResourceCmd).Standalone()

		batch_listJobsByConsumableResourceCmd.Flags().String("consumable-resource", "", "The name or ARN of the consumable resource.")
		batch_listJobsByConsumableResourceCmd.Flags().String("filters", "", "The filters to apply to the job list query.")
		batch_listJobsByConsumableResourceCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListJobsByConsumableResource` in paginated output.")
		batch_listJobsByConsumableResourceCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListJobsByConsumableResource` request where `maxResults` was used and the results exceeded the value of that parameter.")
		batch_listJobsByConsumableResourceCmd.MarkFlagRequired("consumable-resource")
	})
	batchCmd.AddCommand(batch_listJobsByConsumableResourceCmd)
}
