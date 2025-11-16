package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelCardVersionsCmd = &cobra.Command{
	Use:   "list-model-card-versions",
	Short: "List existing versions of an Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelCardVersionsCmd).Standalone()

	sagemaker_listModelCardVersionsCmd.Flags().String("creation-time-after", "", "Only list model card versions that were created after the time specified.")
	sagemaker_listModelCardVersionsCmd.Flags().String("creation-time-before", "", "Only list model card versions that were created before the time specified.")
	sagemaker_listModelCardVersionsCmd.Flags().String("max-results", "", "The maximum number of model card versions to list.")
	sagemaker_listModelCardVersionsCmd.Flags().String("model-card-name", "", "List model card versions for the model card with the specified name or Amazon Resource Name (ARN).")
	sagemaker_listModelCardVersionsCmd.Flags().String("model-card-status", "", "Only list model card versions with the specified approval status.")
	sagemaker_listModelCardVersionsCmd.Flags().String("next-token", "", "If the response to a previous `ListModelCardVersions` request was truncated, the response includes a `NextToken`.")
	sagemaker_listModelCardVersionsCmd.Flags().String("sort-by", "", "Sort listed model card versions by version.")
	sagemaker_listModelCardVersionsCmd.Flags().String("sort-order", "", "Sort model card versions by ascending or descending order.")
	sagemaker_listModelCardVersionsCmd.MarkFlagRequired("model-card-name")
	sagemakerCmd.AddCommand(sagemaker_listModelCardVersionsCmd)
}
