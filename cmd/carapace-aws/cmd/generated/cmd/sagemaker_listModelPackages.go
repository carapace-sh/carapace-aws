package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelPackagesCmd = &cobra.Command{
	Use:   "list-model-packages",
	Short: "Lists the model packages that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelPackagesCmd).Standalone()

	sagemaker_listModelPackagesCmd.Flags().String("creation-time-after", "", "A filter that returns only model packages created after the specified time (timestamp).")
	sagemaker_listModelPackagesCmd.Flags().String("creation-time-before", "", "A filter that returns only model packages created before the specified time (timestamp).")
	sagemaker_listModelPackagesCmd.Flags().String("max-results", "", "The maximum number of model packages to return in the response.")
	sagemaker_listModelPackagesCmd.Flags().String("model-approval-status", "", "A filter that returns only the model packages with the specified approval status.")
	sagemaker_listModelPackagesCmd.Flags().String("model-package-group-name", "", "A filter that returns only model versions that belong to the specified model group.")
	sagemaker_listModelPackagesCmd.Flags().String("model-package-type", "", "A filter that returns only the model packages of the specified type.")
	sagemaker_listModelPackagesCmd.Flags().String("name-contains", "", "A string in the model package name.")
	sagemaker_listModelPackagesCmd.Flags().String("next-token", "", "If the response to a previous `ListModelPackages` request was truncated, the response includes a `NextToken`.")
	sagemaker_listModelPackagesCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
	sagemaker_listModelPackagesCmd.Flags().String("sort-order", "", "The sort order for the results.")
	sagemakerCmd.AddCommand(sagemaker_listModelPackagesCmd)
}
