package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTopicsCmd = &cobra.Command{
	Use:   "list-topics",
	Short: "Lists all of the topics within an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTopicsCmd).Standalone()

	quicksight_listTopicsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topics that you want to list.")
	quicksight_listTopicsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listTopicsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listTopicsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listTopicsCmd)
}
