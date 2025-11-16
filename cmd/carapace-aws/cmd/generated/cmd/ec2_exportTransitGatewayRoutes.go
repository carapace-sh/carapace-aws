package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_exportTransitGatewayRoutesCmd = &cobra.Command{
	Use:   "export-transit-gateway-routes",
	Short: "Exports routes from the specified transit gateway route table to the specified S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_exportTransitGatewayRoutesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_exportTransitGatewayRoutesCmd).Standalone()

		ec2_exportTransitGatewayRoutesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_exportTransitGatewayRoutesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_exportTransitGatewayRoutesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_exportTransitGatewayRoutesCmd.Flags().String("s3-bucket", "", "The name of the S3 bucket.")
		ec2_exportTransitGatewayRoutesCmd.Flags().String("transit-gateway-route-table-id", "", "The ID of the route table.")
		ec2_exportTransitGatewayRoutesCmd.Flag("no-dry-run").Hidden = true
		ec2_exportTransitGatewayRoutesCmd.MarkFlagRequired("s3-bucket")
		ec2_exportTransitGatewayRoutesCmd.MarkFlagRequired("transit-gateway-route-table-id")
	})
	ec2Cmd.AddCommand(ec2_exportTransitGatewayRoutesCmd)
}
