package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createIpamPrefixListResolverTargetCmd = &cobra.Command{
	Use:   "create-ipam-prefix-list-resolver-target",
	Short: "Creates an IPAM prefix list resolver target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createIpamPrefixListResolverTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createIpamPrefixListResolverTargetCmd).Standalone()

		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("desired-version", "", "The specific version of the prefix list to target.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver that will manage the synchronization of CIDRs to the target prefix list.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().Bool("no-track-latest-version", false, "Indicates whether the resolver target should automatically track the latest version of the prefix list.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("prefix-list-id", "", "The ID of the managed prefix list that will be synchronized with CIDRs selected by the IPAM prefix list resolver.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("prefix-list-region", "", "The Amazon Web Services Region where the prefix list is located.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().String("tag-specifications", "", "The tags to apply to the IPAM prefix list resolver target during creation.")
		ec2_createIpamPrefixListResolverTargetCmd.Flags().Bool("track-latest-version", false, "Indicates whether the resolver target should automatically track the latest version of the prefix list.")
		ec2_createIpamPrefixListResolverTargetCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
		ec2_createIpamPrefixListResolverTargetCmd.Flag("no-dry-run").Hidden = true
		ec2_createIpamPrefixListResolverTargetCmd.Flag("no-track-latest-version").Hidden = true
		ec2_createIpamPrefixListResolverTargetCmd.MarkFlagRequired("no-track-latest-version")
		ec2_createIpamPrefixListResolverTargetCmd.MarkFlagRequired("prefix-list-id")
		ec2_createIpamPrefixListResolverTargetCmd.MarkFlagRequired("prefix-list-region")
		ec2_createIpamPrefixListResolverTargetCmd.MarkFlagRequired("track-latest-version")
	})
	ec2Cmd.AddCommand(ec2_createIpamPrefixListResolverTargetCmd)
}
