package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamPrefixListResolverVersionEntriesCmd = &cobra.Command{
	Use:   "get-ipam-prefix-list-resolver-version-entries",
	Short: "Retrieves the CIDR entries for a specific version of an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamPrefixListResolverVersionEntriesCmd).Standalone()

	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver whose version entries you want to retrieve.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().String("ipam-prefix-list-resolver-version", "", "The version number of the resolver for which to retrieve CIDR entries.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.MarkFlagRequired("ipam-prefix-list-resolver-version")
	ec2_getIpamPrefixListResolverVersionEntriesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getIpamPrefixListResolverVersionEntriesCmd)
}
