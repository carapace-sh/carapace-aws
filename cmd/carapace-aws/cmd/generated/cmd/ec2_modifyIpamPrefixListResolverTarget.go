package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIpamPrefixListResolverTargetCmd = &cobra.Command{
	Use:   "modify-ipam-prefix-list-resolver-target",
	Short: "Modifies an IPAM prefix list resolver target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIpamPrefixListResolverTargetCmd).Standalone()

	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().String("desired-version", "", "The desired version of the prefix list to target.")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().String("ipam-prefix-list-resolver-target-id", "", "The ID of the IPAM prefix list resolver target to modify.")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flags().String("track-latest-version", "", "Indicates whether the resolver target should automatically track the latest version of the prefix list.")
	ec2_modifyIpamPrefixListResolverTargetCmd.MarkFlagRequired("ipam-prefix-list-resolver-target-id")
	ec2_modifyIpamPrefixListResolverTargetCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyIpamPrefixListResolverTargetCmd)
}
