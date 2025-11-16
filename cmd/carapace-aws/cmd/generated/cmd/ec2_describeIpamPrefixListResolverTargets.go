package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamPrefixListResolverTargetsCmd = &cobra.Command{
	Use:   "describe-ipam-prefix-list-resolver-targets",
	Short: "Describes one or more IPAM prefix list resolver Targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamPrefixListResolverTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpamPrefixListResolverTargetsCmd).Standalone()

		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().String("filters", "", "One or more filters to limit the results.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver to filter targets by.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().String("ipam-prefix-list-resolver-target-ids", "", "The IDs of the IPAM prefix list resolver Targets to describe.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPrefixListResolverTargetsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpamPrefixListResolverTargetsCmd)
}
