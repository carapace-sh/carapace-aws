package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableIpamOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "enable-ipam-organization-admin-account",
	Short: "Enable an Organizations member account as the IPAM admin account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableIpamOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableIpamOrganizationAdminAccountCmd).Standalone()

		ec2_enableIpamOrganizationAdminAccountCmd.Flags().String("delegated-admin-account-id", "", "The Organizations member account ID that you want to enable as the IPAM account.")
		ec2_enableIpamOrganizationAdminAccountCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_enableIpamOrganizationAdminAccountCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_enableIpamOrganizationAdminAccountCmd.MarkFlagRequired("delegated-admin-account-id")
		ec2_enableIpamOrganizationAdminAccountCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableIpamOrganizationAdminAccountCmd)
}
