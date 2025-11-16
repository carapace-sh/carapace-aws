package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listImageVersionsCmd = &cobra.Command{
	Use:   "list-image-versions",
	Short: "Lists the versions of a specified image and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listImageVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listImageVersionsCmd).Standalone()

		sagemaker_listImageVersionsCmd.Flags().String("creation-time-after", "", "A filter that returns only versions created on or after the specified time.")
		sagemaker_listImageVersionsCmd.Flags().String("creation-time-before", "", "A filter that returns only versions created on or before the specified time.")
		sagemaker_listImageVersionsCmd.Flags().String("image-name", "", "The name of the image to list the versions of.")
		sagemaker_listImageVersionsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only versions modified on or after the specified time.")
		sagemaker_listImageVersionsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only versions modified on or before the specified time.")
		sagemaker_listImageVersionsCmd.Flags().String("max-results", "", "The maximum number of versions to return in the response.")
		sagemaker_listImageVersionsCmd.Flags().String("next-token", "", "If the previous call to `ListImageVersions` didn't return the full set of versions, the call returns a token for getting the next set of versions.")
		sagemaker_listImageVersionsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listImageVersionsCmd.Flags().String("sort-order", "", "The sort order.")
		sagemaker_listImageVersionsCmd.MarkFlagRequired("image-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listImageVersionsCmd)
}
