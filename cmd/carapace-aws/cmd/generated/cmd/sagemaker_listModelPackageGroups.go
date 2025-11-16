package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelPackageGroupsCmd = &cobra.Command{
	Use:   "list-model-package-groups",
	Short: "Gets a list of the model groups in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelPackageGroupsCmd).Standalone()

	sagemaker_listModelPackageGroupsCmd.Flags().String("creation-time-after", "", "A filter that returns only model groups created after the specified time.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("creation-time-before", "", "A filter that returns only model groups created before the specified time.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("cross-account-filter-option", "", "A filter that returns either model groups shared with you or model groups in your own account.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("name-contains", "", "A string in the model group name.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("next-token", "", "If the result of the previous `ListModelPackageGroups` request was truncated, the response includes a `NextToken`.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("sort-by", "", "The field to sort results by.")
	sagemaker_listModelPackageGroupsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemakerCmd.AddCommand(sagemaker_listModelPackageGroupsCmd)
}
