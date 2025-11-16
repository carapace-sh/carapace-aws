package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putAccessControlRuleCmd = &cobra.Command{
	Use:   "put-access-control-rule",
	Short: "Adds a new access control rule for the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putAccessControlRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_putAccessControlRuleCmd).Standalone()

		workmail_putAccessControlRuleCmd.Flags().String("actions", "", "Access protocol actions to include in the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("description", "", "The rule description.")
		workmail_putAccessControlRuleCmd.Flags().String("effect", "", "The rule effect.")
		workmail_putAccessControlRuleCmd.Flags().String("impersonation-role-ids", "", "Impersonation role IDs to include in the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("ip-ranges", "", "IPv4 CIDR ranges to include in the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("name", "", "The rule name.")
		workmail_putAccessControlRuleCmd.Flags().String("not-actions", "", "Access protocol actions to exclude from the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("not-impersonation-role-ids", "", "Impersonation role IDs to exclude from the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("not-ip-ranges", "", "IPv4 CIDR ranges to exclude from the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("not-user-ids", "", "User IDs to exclude from the rule.")
		workmail_putAccessControlRuleCmd.Flags().String("organization-id", "", "The identifier of the organization.")
		workmail_putAccessControlRuleCmd.Flags().String("user-ids", "", "User IDs to include in the rule.")
		workmail_putAccessControlRuleCmd.MarkFlagRequired("description")
		workmail_putAccessControlRuleCmd.MarkFlagRequired("effect")
		workmail_putAccessControlRuleCmd.MarkFlagRequired("name")
		workmail_putAccessControlRuleCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_putAccessControlRuleCmd)
}
