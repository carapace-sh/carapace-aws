package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listApprovalTeamsCmd = &cobra.Command{
	Use:   "list-approval-teams",
	Short: "Returns a list of approval teams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listApprovalTeamsCmd).Standalone()

	mpa_listApprovalTeamsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	mpa_listApprovalTeamsCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
	mpaCmd.AddCommand(mpa_listApprovalTeamsCmd)
}
