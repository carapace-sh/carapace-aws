package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Lists details about all member accounts for the current Security Hub administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listMembersCmd).Standalone()

		securityhub_listMembersCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		securityhub_listMembersCmd.Flags().String("next-token", "", "The token that is required for pagination.")
		securityhub_listMembersCmd.Flags().Bool("no-only-associated", false, "Specifies which member accounts to include in the response based on their relationship status with the administrator account.")
		securityhub_listMembersCmd.Flags().Bool("only-associated", false, "Specifies which member accounts to include in the response based on their relationship status with the administrator account.")
		securityhub_listMembersCmd.Flag("no-only-associated").Hidden = true
	})
	securityhubCmd.AddCommand(securityhub_listMembersCmd)
}
