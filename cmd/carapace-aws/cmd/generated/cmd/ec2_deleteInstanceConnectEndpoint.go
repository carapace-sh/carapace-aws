package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteInstanceConnectEndpointCmd = &cobra.Command{
	Use:   "delete-instance-connect-endpoint",
	Short: "Deletes the specified EC2 Instance Connect Endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteInstanceConnectEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteInstanceConnectEndpointCmd).Standalone()

		ec2_deleteInstanceConnectEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteInstanceConnectEndpointCmd.Flags().String("instance-connect-endpoint-id", "", "The ID of the EC2 Instance Connect Endpoint to delete.")
		ec2_deleteInstanceConnectEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteInstanceConnectEndpointCmd.MarkFlagRequired("instance-connect-endpoint-id")
		ec2_deleteInstanceConnectEndpointCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteInstanceConnectEndpointCmd)
}
