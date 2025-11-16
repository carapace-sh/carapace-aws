package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamPrefixListResolverRulesCmd = &cobra.Command{
	Use:   "get-ipam-prefix-list-resolver-rules",
	Short: "Retrieves the CIDR selection rules for an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamPrefixListResolverRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamPrefixListResolverRulesCmd).Standalone()

		ec2_getIpamPrefixListResolverRulesCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamPrefixListResolverRulesCmd.Flags().String("filters", "", "One or more filters to limit the results.")
		ec2_getIpamPrefixListResolverRulesCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver whose rules you want to retrieve.")
		ec2_getIpamPrefixListResolverRulesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_getIpamPrefixListResolverRulesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getIpamPrefixListResolverRulesCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamPrefixListResolverRulesCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
		ec2_getIpamPrefixListResolverRulesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getIpamPrefixListResolverRulesCmd)
}
