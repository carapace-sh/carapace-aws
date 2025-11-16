package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_searchTransitGatewayMulticastGroupsCmd = &cobra.Command{
	Use:   "search-transit-gateway-multicast-groups",
	Short: "Searches one or more transit gateway multicast groups and returns the group membership information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_searchTransitGatewayMulticastGroupsCmd).Standalone()

	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
	ec2_searchTransitGatewayMulticastGroupsCmd.Flag("no-dry-run").Hidden = true
	ec2_searchTransitGatewayMulticastGroupsCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	ec2Cmd.AddCommand(ec2_searchTransitGatewayMulticastGroupsCmd)
}
