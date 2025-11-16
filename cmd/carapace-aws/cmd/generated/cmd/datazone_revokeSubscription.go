package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_revokeSubscriptionCmd = &cobra.Command{
	Use:   "revoke-subscription",
	Short: "Revokes a specified subscription in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_revokeSubscriptionCmd).Standalone()

	datazone_revokeSubscriptionCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain where you want to revoke a subscription.")
	datazone_revokeSubscriptionCmd.Flags().String("identifier", "", "The identifier of the revoked subscription.")
	datazone_revokeSubscriptionCmd.Flags().Bool("no-retain-permissions", false, "Specifies whether permissions are retained when the subscription is revoked.")
	datazone_revokeSubscriptionCmd.Flags().Bool("retain-permissions", false, "Specifies whether permissions are retained when the subscription is revoked.")
	datazone_revokeSubscriptionCmd.MarkFlagRequired("domain-identifier")
	datazone_revokeSubscriptionCmd.MarkFlagRequired("identifier")
	datazone_revokeSubscriptionCmd.Flag("no-retain-permissions").Hidden = true
	datazoneCmd.AddCommand(datazone_revokeSubscriptionCmd)
}
