package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIpamPrefixListResolversCmd = &cobra.Command{
	Use:   "describe-ipam-prefix-list-resolvers",
	Short: "Describes one or more IPAM prefix list resolvers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIpamPrefixListResolversCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIpamPrefixListResolversCmd).Standalone()

		ec2_describeIpamPrefixListResolversCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPrefixListResolversCmd.Flags().String("filters", "", "One or more filters to limit the results.")
		ec2_describeIpamPrefixListResolversCmd.Flags().String("ipam-prefix-list-resolver-ids", "", "The IDs of the IPAM prefix list resolvers to describe.")
		ec2_describeIpamPrefixListResolversCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeIpamPrefixListResolversCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeIpamPrefixListResolversCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_describeIpamPrefixListResolversCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeIpamPrefixListResolversCmd)
}
