package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayPrefixListReferencesCmd = &cobra.Command{
	Use:   "get-transit-gateway-prefix-list-references",
	Short: "Gets information about the prefix list references in a specified transit gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayPrefixListReferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayPrefixListReferencesCmd).Standalone()

		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the transit gateway route table.")
		ec2_getTransitGatewayPrefixListReferencesCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayPrefixListReferencesCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayPrefixListReferencesCmd)
}
