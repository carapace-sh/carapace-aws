package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectVpcEndpointConnectionsCmd = &cobra.Command{
	Use:   "reject-vpc-endpoint-connections",
	Short: "Rejects VPC endpoint connection requests to your VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectVpcEndpointConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_rejectVpcEndpointConnectionsCmd).Standalone()

		ec2_rejectVpcEndpointConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectVpcEndpointConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_rejectVpcEndpointConnectionsCmd.Flags().String("service-id", "", "The ID of the service.")
		ec2_rejectVpcEndpointConnectionsCmd.Flags().String("vpc-endpoint-ids", "", "The IDs of the VPC endpoints.")
		ec2_rejectVpcEndpointConnectionsCmd.Flag("no-dry-run").Hidden = true
		ec2_rejectVpcEndpointConnectionsCmd.MarkFlagRequired("service-id")
		ec2_rejectVpcEndpointConnectionsCmd.MarkFlagRequired("vpc-endpoint-ids")
	})
	ec2Cmd.AddCommand(ec2_rejectVpcEndpointConnectionsCmd)
}
