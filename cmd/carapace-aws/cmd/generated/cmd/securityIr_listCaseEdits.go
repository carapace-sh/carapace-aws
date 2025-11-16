package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_listCaseEditsCmd = &cobra.Command{
	Use:   "list-case-edits",
	Short: "Views the case history for edits made to a designated case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_listCaseEditsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_listCaseEditsCmd).Standalone()

		securityIr_listCaseEditsCmd.Flags().String("case-id", "", "Required element used with ListCaseEdits to identify the case to query.")
		securityIr_listCaseEditsCmd.Flags().String("max-results", "", "Optional element to identify how many results to obtain.")
		securityIr_listCaseEditsCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to ListCaseEdits.")
		securityIr_listCaseEditsCmd.MarkFlagRequired("case-id")
	})
	securityIrCmd.AddCommand(securityIr_listCaseEditsCmd)
}
