package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelCardsCmd = &cobra.Command{
	Use:   "list-model-cards",
	Short: "List existing model cards.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelCardsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listModelCardsCmd).Standalone()

		sagemaker_listModelCardsCmd.Flags().String("creation-time-after", "", "Only list model cards that were created after the time specified.")
		sagemaker_listModelCardsCmd.Flags().String("creation-time-before", "", "Only list model cards that were created before the time specified.")
		sagemaker_listModelCardsCmd.Flags().String("max-results", "", "The maximum number of model cards to list.")
		sagemaker_listModelCardsCmd.Flags().String("model-card-status", "", "Only list model cards with the specified approval status.")
		sagemaker_listModelCardsCmd.Flags().String("name-contains", "", "Only list model cards with names that contain the specified string.")
		sagemaker_listModelCardsCmd.Flags().String("next-token", "", "If the response to a previous `ListModelCards` request was truncated, the response includes a `NextToken`.")
		sagemaker_listModelCardsCmd.Flags().String("sort-by", "", "Sort model cards by either name or creation time.")
		sagemaker_listModelCardsCmd.Flags().String("sort-order", "", "Sort model cards by ascending or descending order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listModelCardsCmd)
}
