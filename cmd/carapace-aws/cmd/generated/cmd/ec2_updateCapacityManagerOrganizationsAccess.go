package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_updateCapacityManagerOrganizationsAccessCmd = &cobra.Command{
	Use:   "update-capacity-manager-organizations-access",
	Short: "Updates the Organizations access setting for EC2 Capacity Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_updateCapacityManagerOrganizationsAccessCmd).Standalone()

	ec2_updateCapacityManagerOrganizationsAccessCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_updateCapacityManagerOrganizationsAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_updateCapacityManagerOrganizationsAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_updateCapacityManagerOrganizationsAccessCmd.Flags().String("organizations-access", "", "Specifies whether to enable or disable cross-account access for Amazon Web Services Organizations.")
	ec2_updateCapacityManagerOrganizationsAccessCmd.Flag("no-dry-run").Hidden = true
	ec2_updateCapacityManagerOrganizationsAccessCmd.MarkFlagRequired("organizations-access")
	ec2Cmd.AddCommand(ec2_updateCapacityManagerOrganizationsAccessCmd)
}
