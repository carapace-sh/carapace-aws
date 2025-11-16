package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Lists all the templates in the current Amazon Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTemplatesCmd).Standalone()

	quicksight_listTemplatesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the templates that you're listing.")
	quicksight_listTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listTemplatesCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listTemplatesCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listTemplatesCmd)
}
