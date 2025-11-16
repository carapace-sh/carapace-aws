package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTransitGatewayMulticastDomainsCmd = &cobra.Command{
	Use:   "describe-transit-gateway-multicast-domains",
	Short: "Describes one or more transit gateway multicast domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTransitGatewayMulticastDomainsCmd).Standalone()

	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flags().String("transit-gateway-multicast-domain-ids", "", "The ID of the transit gateway multicast domain.")
	ec2_describeTransitGatewayMulticastDomainsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeTransitGatewayMulticastDomainsCmd)
}
