package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_detachInternetGatewayCmd = &cobra.Command{
	Use:   "detach-internet-gateway",
	Short: "Detaches an internet gateway from a VPC, disabling connectivity between the internet and the VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_detachInternetGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_detachInternetGatewayCmd).Standalone()

		ec2_detachInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachInternetGatewayCmd.Flags().String("internet-gateway-id", "", "The ID of the internet gateway.")
		ec2_detachInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_detachInternetGatewayCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_detachInternetGatewayCmd.MarkFlagRequired("internet-gateway-id")
		ec2_detachInternetGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_detachInternetGatewayCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_detachInternetGatewayCmd)
}
