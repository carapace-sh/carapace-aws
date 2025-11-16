package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceConnectEndpointsCmd = &cobra.Command{
	Use:   "describe-instance-connect-endpoints",
	Short: "Describes the specified EC2 Instance Connect Endpoints or all EC2 Instance Connect Endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceConnectEndpointsCmd).Standalone()

	ec2_describeInstanceConnectEndpointsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceConnectEndpointsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeInstanceConnectEndpointsCmd.Flags().String("instance-connect-endpoint-ids", "", "One or more EC2 Instance Connect Endpoint IDs.")
	ec2_describeInstanceConnectEndpointsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceConnectEndpointsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceConnectEndpointsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceConnectEndpointsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceConnectEndpointsCmd)
}
