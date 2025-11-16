package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAssociationsCmd = &cobra.Command{
	Use:   "list-associations",
	Short: "Lists the associations in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listAssociationsCmd).Standalone()

		sagemaker_listAssociationsCmd.Flags().String("association-type", "", "A filter that returns only associations of the specified type.")
		sagemaker_listAssociationsCmd.Flags().String("created-after", "", "A filter that returns only associations created on or after the specified time.")
		sagemaker_listAssociationsCmd.Flags().String("created-before", "", "A filter that returns only associations created on or before the specified time.")
		sagemaker_listAssociationsCmd.Flags().String("destination-arn", "", "A filter that returns only associations with the specified destination Amazon Resource Name (ARN).")
		sagemaker_listAssociationsCmd.Flags().String("destination-type", "", "A filter that returns only associations with the specified destination type.")
		sagemaker_listAssociationsCmd.Flags().String("max-results", "", "The maximum number of associations to return in the response.")
		sagemaker_listAssociationsCmd.Flags().String("next-token", "", "If the previous call to `ListAssociations` didn't return the full set of associations, the call returns a token for getting the next set of associations.")
		sagemaker_listAssociationsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listAssociationsCmd.Flags().String("sort-order", "", "The sort order.")
		sagemaker_listAssociationsCmd.Flags().String("source-arn", "", "A filter that returns only associations with the specified source ARN.")
		sagemaker_listAssociationsCmd.Flags().String("source-type", "", "A filter that returns only associations with the specified source type.")
	})
	sagemakerCmd.AddCommand(sagemaker_listAssociationsCmd)
}
