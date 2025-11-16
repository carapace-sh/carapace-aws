package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_createMembershipCmd = &cobra.Command{
	Use:   "create-membership",
	Short: "Creates a new membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_createMembershipCmd).Standalone()

	securityIr_createMembershipCmd.Flags().String("client-token", "", "The `clientToken` field is an idempotency key used to ensure that repeated attempts for a single action will be ignored by the server during retries.")
	securityIr_createMembershipCmd.Flags().Bool("cover-entire-organization", false, "The `coverEntireOrganization` parameter is a boolean flag that determines whether the membership should be applied to the entire Amazon Web Services Organization.")
	securityIr_createMembershipCmd.Flags().String("incident-response-team", "", "Required element used in combination with CreateMembership to add customer incident response team members and trusted partners to the membership.")
	securityIr_createMembershipCmd.Flags().String("membership-name", "", "Required element used in combination with CreateMembership to create a name for the membership.")
	securityIr_createMembershipCmd.Flags().Bool("no-cover-entire-organization", false, "The `coverEntireOrganization` parameter is a boolean flag that determines whether the membership should be applied to the entire Amazon Web Services Organization.")
	securityIr_createMembershipCmd.Flags().String("opt-in-features", "", "Optional element to enable the monitoring and investigation opt-in features for the service.")
	securityIr_createMembershipCmd.Flags().String("tags", "", "Optional element for customer configured tags.")
	securityIr_createMembershipCmd.MarkFlagRequired("incident-response-team")
	securityIr_createMembershipCmd.MarkFlagRequired("membership-name")
	securityIr_createMembershipCmd.Flag("no-cover-entire-organization").Hidden = true
	securityIrCmd.AddCommand(securityIr_createMembershipCmd)
}
