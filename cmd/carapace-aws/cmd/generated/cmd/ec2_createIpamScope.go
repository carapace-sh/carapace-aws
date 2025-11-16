package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamScopeCmd = &cobra.Command{
	Use:   "create-ipam-scope",
	Short: "Create an IPAM scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createIpamScopeCmd).Standalone()

		ec2_createIpamScopeCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createIpamScopeCmd.Flags().String("description", "", "A description for the scope you're creating.")
		ec2_createIpamScopeCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamScopeCmd.Flags().String("external-authority-configuration", "", "The configuration that links an Amazon VPC IPAM scope to an external authority system.")
		ec2_createIpamScopeCmd.Flags().String("ipam-id", "", "The ID of the IPAM for which you're creating this scope.")
		ec2_createIpamScopeCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamScopeCmd.Flags().String("tag-specifications", "", "The key/value combination of a tag assigned to the resource.")
		ec2_createIpamScopeCmd.MarkFlagRequired("ipam-id")
		ec2_createIpamScopeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createIpamScopeCmd)
}
