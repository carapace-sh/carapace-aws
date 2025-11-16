package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_attachInternetGatewayCmd = &cobra.Command{
	Use:   "attach-internet-gateway",
	Short: "Attaches an internet gateway or a virtual private gateway to a VPC, enabling connectivity between the internet and the VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_attachInternetGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_attachInternetGatewayCmd).Standalone()

		ec2_attachInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachInternetGatewayCmd.Flags().String("internet-gateway-id", "", "The ID of the internet gateway.")
		ec2_attachInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_attachInternetGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_attachInternetGatewayCmd.MarkFlagRequired("internet-gateway-id")
		ec2_attachInternetGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_attachInternetGatewayCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_attachInternetGatewayCmd)
}
