package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getMembersCmd = &cobra.Command{
	Use:   "get-members",
	Short: "Returns the details for the Security Hub member accounts for the specified account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getMembersCmd).Standalone()

	securityhub_getMembersCmd.Flags().String("account-ids", "", "The list of account IDs for the Security Hub member accounts to return the details for.")
	securityhub_getMembersCmd.MarkFlagRequired("account-ids")
	securityhubCmd.AddCommand(securityhub_getMembersCmd)
}
