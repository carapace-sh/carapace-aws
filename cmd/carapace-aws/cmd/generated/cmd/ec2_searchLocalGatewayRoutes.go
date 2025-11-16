package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_searchLocalGatewayRoutesCmd = &cobra.Command{
	Use:   "search-local-gateway-routes",
	Short: "Searches for routes in the specified local gateway route table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_searchLocalGatewayRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_searchLocalGatewayRoutesCmd).Standalone()

		ec2_searchLocalGatewayRoutesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_searchLocalGatewayRoutesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_searchLocalGatewayRoutesCmd.Flags().String("local-gateway-route-table-id", "", "The ID of the local gateway route table.")
		ec2_searchLocalGatewayRoutesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_searchLocalGatewayRoutesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_searchLocalGatewayRoutesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_searchLocalGatewayRoutesCmd.MarkFlagRequired("local-gateway-route-table-id")
		ec2_searchLocalGatewayRoutesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_searchLocalGatewayRoutesCmd)
}
