package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableIpamOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "disable-ipam-organization-admin-account",
	Short: "Disable the IPAM account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableIpamOrganizationAdminAccountCmd).Standalone()

	ec2_disableIpamOrganizationAdminAccountCmd.Flags().String("delegated-admin-account-id", "", "The Organizations member account ID that you want to disable as IPAM account.")
	ec2_disableIpamOrganizationAdminAccountCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_disableIpamOrganizationAdminAccountCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_disableIpamOrganizationAdminAccountCmd.MarkFlagRequired("delegated-admin-account-id")
	ec2_disableIpamOrganizationAdminAccountCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableIpamOrganizationAdminAccountCmd)
}
