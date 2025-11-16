package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createInstanceConnectEndpointCmd = &cobra.Command{
	Use:   "create-instance-connect-endpoint",
	Short: "Creates an EC2 Instance Connect Endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createInstanceConnectEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createInstanceConnectEndpointCmd).Standalone()

		ec2_createInstanceConnectEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createInstanceConnectEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createInstanceConnectEndpointCmd.Flags().String("ip-address-type", "", "The IP address type of the endpoint.")
		ec2_createInstanceConnectEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createInstanceConnectEndpointCmd.Flags().Bool("no-preserve-client-ip", false, "Indicates whether the client IP address is preserved as the source.")
		ec2_createInstanceConnectEndpointCmd.Flags().Bool("preserve-client-ip", false, "Indicates whether the client IP address is preserved as the source.")
		ec2_createInstanceConnectEndpointCmd.Flags().String("security-group-ids", "", "One or more security groups to associate with the endpoint.")
		ec2_createInstanceConnectEndpointCmd.Flags().String("subnet-id", "", "The ID of the subnet in which to create the EC2 Instance Connect Endpoint.")
		ec2_createInstanceConnectEndpointCmd.Flags().String("tag-specifications", "", "The tags to apply to the EC2 Instance Connect Endpoint during creation.")
		ec2_createInstanceConnectEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_createInstanceConnectEndpointCmd.Flag("no-preserve-client-ip").Hidden = true
		ec2_createInstanceConnectEndpointCmd.MarkFlagRequired("subnet-id")
	})
	ec2Cmd.AddCommand(ec2_createInstanceConnectEndpointCmd)
}
