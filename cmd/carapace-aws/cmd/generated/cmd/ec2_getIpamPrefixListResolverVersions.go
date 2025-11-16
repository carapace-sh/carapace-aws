package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamPrefixListResolverVersionsCmd = &cobra.Command{
	Use:   "get-ipam-prefix-list-resolver-versions",
	Short: "Retrieves version information for an IPAM prefix list resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamPrefixListResolverVersionsCmd).Standalone()

	ec2_getIpamPrefixListResolverVersionsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().String("filters", "", "One or more filters to limit the results.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().String("ipam-prefix-list-resolver-id", "", "The ID of the IPAM prefix list resolver whose versions you want to retrieve.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().String("ipam-prefix-list-resolver-versions", "", "Specific version numbers to retrieve.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getIpamPrefixListResolverVersionsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
	ec2_getIpamPrefixListResolverVersionsCmd.MarkFlagRequired("ipam-prefix-list-resolver-id")
	ec2_getIpamPrefixListResolverVersionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getIpamPrefixListResolverVersionsCmd)
}
