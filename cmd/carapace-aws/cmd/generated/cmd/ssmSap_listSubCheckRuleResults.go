package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listSubCheckRuleResultsCmd = &cobra.Command{
	Use:   "list-sub-check-rule-results",
	Short: "Lists the rules of a specified sub-check belonging to a configuration check operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listSubCheckRuleResultsCmd).Standalone()

	ssmSap_listSubCheckRuleResultsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ssmSap_listSubCheckRuleResultsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ssmSap_listSubCheckRuleResultsCmd.Flags().String("sub-check-result-id", "", "The ID of the sub check result.")
	ssmSap_listSubCheckRuleResultsCmd.MarkFlagRequired("sub-check-result-id")
	ssmSapCmd.AddCommand(ssmSap_listSubCheckRuleResultsCmd)
}
