package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamScopeCmd = &cobra.Command{
	Use:   "modify-ipam-scope",
	Short: "Modify an IPAM scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamScopeCmd).Standalone()

	ec2_modifyIpamScopeCmd.Flags().String("description", "", "The description of the scope you want to modify.")
	ec2_modifyIpamScopeCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamScopeCmd.Flags().String("external-authority-configuration", "", "The configuration that links an Amazon VPC IPAM scope to an external authority system.")
	ec2_modifyIpamScopeCmd.Flags().String("ipam-scope-id", "", "The ID of the scope you want to modify.")
	ec2_modifyIpamScopeCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamScopeCmd.Flags().Bool("no-remove-external-authority-configuration", false, "Remove the external authority configuration.")
	ec2_modifyIpamScopeCmd.Flags().Bool("remove-external-authority-configuration", false, "Remove the external authority configuration.")
	ec2_modifyIpamScopeCmd.MarkFlagRequired("ipam-scope-id")
	ec2_modifyIpamScopeCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyIpamScopeCmd.Flag("no-remove-external-authority-configuration").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyIpamScopeCmd)
}
