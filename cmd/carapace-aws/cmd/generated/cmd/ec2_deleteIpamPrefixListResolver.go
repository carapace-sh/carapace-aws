package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteIpamPrefixListResolverCmd = &cobra.Command{
	Use:   "delete-ipam-prefix-list-resolver",
	Short: "Deletes an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteIpamPrefixListResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteIpamPrefixListResolverCmd).Standalone()

		ec2_deleteIpamPrefixListResolverCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPrefixListResolverCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver to delete.")
		ec2_deleteIpamPrefixListResolverCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_deleteIpamPrefixListResolverCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
		ec2_deleteIpamPrefixListResolverCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteIpamPrefixListResolverCmd)
}
