package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptVpcEndpointConnectionsCmd = &cobra.Command{
	Use:   "accept-vpc-endpoint-connections",
	Short: "Accepts connection requests to your VPC endpoint service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptVpcEndpointConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_acceptVpcEndpointConnectionsCmd).Standalone()

		ec2_acceptVpcEndpointConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_acceptVpcEndpointConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_acceptVpcEndpointConnectionsCmd.Flags().String("service-id", "", "The ID of the VPC endpoint service.")
		ec2_acceptVpcEndpointConnectionsCmd.Flags().String("vpc-endpoint-ids", "", "The IDs of the interface VPC endpoints.")
		ec2_acceptVpcEndpointConnectionsCmd.Flag("no-dry-run").Hidden = true
		ec2_acceptVpcEndpointConnectionsCmd.MarkFlagRequired("service-id")
		ec2_acceptVpcEndpointConnectionsCmd.MarkFlagRequired("vpc-endpoint-ids")
	})
	ec2Cmd.AddCommand(ec2_acceptVpcEndpointConnectionsCmd)
}
