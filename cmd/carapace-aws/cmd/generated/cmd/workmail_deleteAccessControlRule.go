package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteAccessControlRuleCmd = &cobra.Command{
	Use:   "delete-access-control-rule",
	Short: "Deletes an access control rule for the specified WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteAccessControlRuleCmd).Standalone()

	workmail_deleteAccessControlRuleCmd.Flags().String("name", "", "The name of the access control rule.")
	workmail_deleteAccessControlRuleCmd.Flags().String("organization-id", "", "The identifier for the organization.")
	workmail_deleteAccessControlRuleCmd.MarkFlagRequired("name")
	workmail_deleteAccessControlRuleCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteAccessControlRuleCmd)
}
