package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getAccessControlEffectCmd = &cobra.Command{
	Use:   "get-access-control-effect",
	Short: "Gets the effects of an organization's access control rules as they apply to a specified IPv4 address, access protocol action, and user ID or impersonation role ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getAccessControlEffectCmd).Standalone()

	workmail_getAccessControlEffectCmd.Flags().String("action", "", "The access protocol action.")
	workmail_getAccessControlEffectCmd.Flags().String("impersonation-role-id", "", "The impersonation role ID.")
	workmail_getAccessControlEffectCmd.Flags().String("ip-address", "", "The IPv4 address.")
	workmail_getAccessControlEffectCmd.Flags().String("organization-id", "", "The identifier for the organization.")
	workmail_getAccessControlEffectCmd.Flags().String("user-id", "", "The user ID.")
	workmail_getAccessControlEffectCmd.MarkFlagRequired("action")
	workmail_getAccessControlEffectCmd.MarkFlagRequired("ip-address")
	workmail_getAccessControlEffectCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_getAccessControlEffectCmd)
}
