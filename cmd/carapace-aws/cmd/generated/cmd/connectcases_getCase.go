package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getCaseCmd = &cobra.Command{
	Use:   "get-case",
	Short: "Returns information about a specific case if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getCaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_getCaseCmd).Standalone()

		connectcases_getCaseCmd.Flags().String("case-id", "", "A unique identifier of the case.")
		connectcases_getCaseCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_getCaseCmd.Flags().String("fields", "", "A list of unique field identifiers.")
		connectcases_getCaseCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connectcases_getCaseCmd.MarkFlagRequired("case-id")
		connectcases_getCaseCmd.MarkFlagRequired("domain-id")
		connectcases_getCaseCmd.MarkFlagRequired("fields")
	})
	connectcasesCmd.AddCommand(connectcases_getCaseCmd)
}
