package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listAccessControlRulesCmd = &cobra.Command{
	Use:   "list-access-control-rules",
	Short: "Lists the access control rules for the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listAccessControlRulesCmd).Standalone()

	workmail_listAccessControlRulesCmd.Flags().String("organization-id", "", "The identifier for the organization.")
	workmail_listAccessControlRulesCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listAccessControlRulesCmd)
}
