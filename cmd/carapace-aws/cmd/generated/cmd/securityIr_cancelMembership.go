package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_cancelMembershipCmd = &cobra.Command{
	Use:   "cancel-membership",
	Short: "Cancels an existing membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_cancelMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_cancelMembershipCmd).Standalone()

		securityIr_cancelMembershipCmd.Flags().String("membership-id", "", "Required element used in combination with CancelMembershipRequest to identify the membership ID to cancel.")
		securityIr_cancelMembershipCmd.MarkFlagRequired("membership-id")
	})
	securityIrCmd.AddCommand(securityIr_cancelMembershipCmd)
}
