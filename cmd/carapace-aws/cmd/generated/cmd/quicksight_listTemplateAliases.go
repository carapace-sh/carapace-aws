package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTemplateAliasesCmd = &cobra.Command{
	Use:   "list-template-aliases",
	Short: "Lists all the aliases of a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTemplateAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listTemplateAliasesCmd).Standalone()

		quicksight_listTemplateAliasesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template aliases that you're listing.")
		quicksight_listTemplateAliasesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listTemplateAliasesCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listTemplateAliasesCmd.Flags().String("template-id", "", "The ID for the template.")
		quicksight_listTemplateAliasesCmd.MarkFlagRequired("aws-account-id")
		quicksight_listTemplateAliasesCmd.MarkFlagRequired("template-id")
	})
	quicksightCmd.AddCommand(quicksight_listTemplateAliasesCmd)
}
