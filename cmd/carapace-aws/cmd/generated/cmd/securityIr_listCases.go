package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_listCasesCmd = &cobra.Command{
	Use:   "list-cases",
	Short: "Lists all cases the requester has access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_listCasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_listCasesCmd).Standalone()

		securityIr_listCasesCmd.Flags().String("max-results", "", "Optional element for ListCases to limit the number of responses.")
		securityIr_listCasesCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to ListCases.")
	})
	securityIrCmd.AddCommand(securityIr_listCasesCmd)
}
