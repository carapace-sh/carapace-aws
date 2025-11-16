package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listLineageGroupsCmd = &cobra.Command{
	Use:   "list-lineage-groups",
	Short: "A list of lineage groups shared with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listLineageGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listLineageGroupsCmd).Standalone()

		sagemaker_listLineageGroupsCmd.Flags().String("created-after", "", "A timestamp to filter against lineage groups created after a certain point in time.")
		sagemaker_listLineageGroupsCmd.Flags().String("created-before", "", "A timestamp to filter against lineage groups created before a certain point in time.")
		sagemaker_listLineageGroupsCmd.Flags().String("max-results", "", "The maximum number of endpoints to return in the response.")
		sagemaker_listLineageGroupsCmd.Flags().String("next-token", "", "If the response is truncated, SageMaker returns this token.")
		sagemaker_listLineageGroupsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemaker_listLineageGroupsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listLineageGroupsCmd)
}
