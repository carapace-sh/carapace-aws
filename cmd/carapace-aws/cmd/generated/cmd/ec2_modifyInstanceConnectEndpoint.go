package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceConnectEndpointCmd = &cobra.Command{
	Use:   "modify-instance-connect-endpoint",
	Short: "Modifies the specified EC2 Instance Connect Endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceConnectEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyInstanceConnectEndpointCmd).Standalone()

		ec2_modifyInstanceConnectEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().String("instance-connect-endpoint-id", "", "The ID of the EC2 Instance Connect Endpoint to modify.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().String("ip-address-type", "", "The new IP address type for the EC2 Instance Connect Endpoint.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().Bool("no-preserve-client-ip", false, "Indicates whether the client IP address is preserved as the source when you connect to a resource.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().Bool("preserve-client-ip", false, "Indicates whether the client IP address is preserved as the source when you connect to a resource.")
		ec2_modifyInstanceConnectEndpointCmd.Flags().String("security-group-ids", "", "Changes the security groups for the EC2 Instance Connect Endpoint.")
		ec2_modifyInstanceConnectEndpointCmd.MarkFlagRequired("instance-connect-endpoint-id")
		ec2_modifyInstanceConnectEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyInstanceConnectEndpointCmd.Flag("no-preserve-client-ip").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyInstanceConnectEndpointCmd)
}
