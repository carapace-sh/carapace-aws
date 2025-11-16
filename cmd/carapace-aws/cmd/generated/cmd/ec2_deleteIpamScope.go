package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamScopeCmd = &cobra.Command{
	Use:   "delete-ipam-scope",
	Short: "Delete the scope for an IPAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteIpamScopeCmd).Standalone()

		ec2_deleteIpamScopeCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamScopeCmd.Flags().String("ipam-scope-id", "", "The ID of the scope to delete.")
		ec2_deleteIpamScopeCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamScopeCmd.MarkFlagRequired("ipam-scope-id")
		ec2_deleteIpamScopeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteIpamScopeCmd)
}
