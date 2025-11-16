package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getIpamDiscoveredAccountsCmd = &cobra.Command{
	Use:   "get-ipam-discovered-accounts",
	Short: "Gets IPAM discovered accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getIpamDiscoveredAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getIpamDiscoveredAccountsCmd).Standalone()

		ec2_getIpamDiscoveredAccountsCmd.Flags().String("discovery-region", "", "The Amazon Web Services Region that the account information is returned from.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().Bool("dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().String("filters", "", "Discovered account filters.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().String("ipam-resource-discovery-id", "", "A resource discovery ID.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().String("max-results", "", "The maximum number of discovered accounts to return in one page of results.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		ec2_getIpamDiscoveredAccountsCmd.Flags().Bool("no-dry-run", false, "A check for whether you have the required permissions for the action without actually making the request and provides an error response.")
		ec2_getIpamDiscoveredAccountsCmd.MarkFlagRequired("discovery-region")
		ec2_getIpamDiscoveredAccountsCmd.MarkFlagRequired("ipam-resource-discovery-id")
		ec2_getIpamDiscoveredAccountsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getIpamDiscoveredAccountsCmd)
}
