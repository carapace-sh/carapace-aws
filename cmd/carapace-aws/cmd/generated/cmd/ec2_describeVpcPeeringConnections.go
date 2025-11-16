package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcPeeringConnectionsCmd = &cobra.Command{
	Use:   "describe-vpc-peering-connections",
	Short: "Describes your VPC peering connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcPeeringConnectionsCmd).Standalone()

	ec2_describeVpcPeeringConnectionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcPeeringConnectionsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeVpcPeeringConnectionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeVpcPeeringConnectionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeVpcPeeringConnectionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcPeeringConnectionsCmd.Flags().String("vpc-peering-connection-ids", "", "The IDs of the VPC peering connections.")
	ec2_describeVpcPeeringConnectionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeVpcPeeringConnectionsCmd)
}
