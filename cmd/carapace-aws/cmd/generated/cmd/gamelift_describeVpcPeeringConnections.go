package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeVpcPeeringConnectionsCmd = &cobra.Command{
	Use:   "describe-vpc-peering-connections",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves information on VPC peering connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeVpcPeeringConnectionsCmd).Standalone()

	gamelift_describeVpcPeeringConnectionsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet.")
	gameliftCmd.AddCommand(gamelift_describeVpcPeeringConnectionsCmd)
}
