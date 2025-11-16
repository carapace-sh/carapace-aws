package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_listCaseRulesCmd = &cobra.Command{
	Use:   "list-case-rules",
	Short: "Lists all case rules in a Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_listCaseRulesCmd).Standalone()

	connectcases_listCaseRulesCmd.Flags().String("domain-id", "", "Unique identifier of a Cases domain.")
	connectcases_listCaseRulesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connectcases_listCaseRulesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connectcases_listCaseRulesCmd.MarkFlagRequired("domain-id")
	connectcasesCmd.AddCommand(connectcases_listCaseRulesCmd)
}
