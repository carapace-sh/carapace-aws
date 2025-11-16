package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTemplateVersionsCmd = &cobra.Command{
	Use:   "list-template-versions",
	Short: "Lists all the versions of the templates in the current Amazon Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTemplateVersionsCmd).Standalone()

	quicksight_listTemplateVersionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the templates that you're listing.")
	quicksight_listTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listTemplateVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listTemplateVersionsCmd.Flags().String("template-id", "", "The ID for the template.")
	quicksight_listTemplateVersionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_listTemplateVersionsCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_listTemplateVersionsCmd)
}
