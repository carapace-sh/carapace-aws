package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamPrefixListResolverTargetCmd = &cobra.Command{
	Use:   "delete-ipam-prefix-list-resolver-target",
	Short: "Deletes an IPAM prefix list resolver target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamPrefixListResolverTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteIpamPrefixListResolverTargetCmd).Standalone()

		ec2_deleteIpamPrefixListResolverTargetCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPrefixListResolverTargetCmd.Flags().String("ipam-prefix-list-resolver-target-id", "", "The ID of the IPAM prefix list resolver target to delete.")
		ec2_deleteIpamPrefixListResolverTargetCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPrefixListResolverTargetCmd.MarkFlagRequired("ipam-prefix-list-resolver-target-id")
		ec2_deleteIpamPrefixListResolverTargetCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteIpamPrefixListResolverTargetCmd)
}
