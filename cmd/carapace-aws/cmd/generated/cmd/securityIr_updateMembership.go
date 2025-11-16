package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_updateMembershipCmd = &cobra.Command{
	Use:   "update-membership",
	Short: "Updates membership configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_updateMembershipCmd).Standalone()

	securityIr_updateMembershipCmd.Flags().String("incident-response-team", "", "Optional element for UpdateMembership to update the membership name.")
	securityIr_updateMembershipCmd.Flags().String("membership-accounts-configurations-update", "", "The `membershipAccountsConfigurationsUpdate` field in the `UpdateMembershipRequest` structure allows you to update the configuration settings for accounts within a membership.")
	securityIr_updateMembershipCmd.Flags().String("membership-id", "", "Required element for UpdateMembership to identify the membership to update.")
	securityIr_updateMembershipCmd.Flags().String("membership-name", "", "Optional element for UpdateMembership to update the membership name.")
	securityIr_updateMembershipCmd.Flags().Bool("no-undo-membership-cancellation", false, "The `undoMembershipCancellation` parameter is a boolean flag that indicates whether to reverse a previously requested membership cancellation.")
	securityIr_updateMembershipCmd.Flags().String("opt-in-features", "", "Optional element for UpdateMembership to enable or disable opt-in features for the service.")
	securityIr_updateMembershipCmd.Flags().Bool("undo-membership-cancellation", false, "The `undoMembershipCancellation` parameter is a boolean flag that indicates whether to reverse a previously requested membership cancellation.")
	securityIr_updateMembershipCmd.MarkFlagRequired("membership-id")
	securityIr_updateMembershipCmd.Flag("no-undo-membership-cancellation").Hidden = true
	securityIrCmd.AddCommand(securityIr_updateMembershipCmd)
}
