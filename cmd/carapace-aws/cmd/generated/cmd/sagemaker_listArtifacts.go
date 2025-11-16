package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listArtifactsCmd = &cobra.Command{
	Use:   "list-artifacts",
	Short: "Lists the artifacts in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listArtifactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listArtifactsCmd).Standalone()

		sagemaker_listArtifactsCmd.Flags().String("artifact-type", "", "A filter that returns only artifacts of the specified type.")
		sagemaker_listArtifactsCmd.Flags().String("created-after", "", "A filter that returns only artifacts created on or after the specified time.")
		sagemaker_listArtifactsCmd.Flags().String("created-before", "", "A filter that returns only artifacts created on or before the specified time.")
		sagemaker_listArtifactsCmd.Flags().String("max-results", "", "The maximum number of artifacts to return in the response.")
		sagemaker_listArtifactsCmd.Flags().String("next-token", "", "If the previous call to `ListArtifacts` didn't return the full set of artifacts, the call returns a token for getting the next set of artifacts.")
		sagemaker_listArtifactsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listArtifactsCmd.Flags().String("sort-order", "", "The sort order.")
		sagemaker_listArtifactsCmd.Flags().String("source-uri", "", "A filter that returns only artifacts with the specified source URI.")
	})
	sagemakerCmd.AddCommand(sagemaker_listArtifactsCmd)
}
