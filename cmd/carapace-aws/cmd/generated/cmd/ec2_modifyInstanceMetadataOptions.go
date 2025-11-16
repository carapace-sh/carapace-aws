package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceMetadataOptionsCmd = &cobra.Command{
	Use:   "modify-instance-metadata-options",
	Short: "Modify the instance metadata parameters on a running or stopped instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceMetadataOptionsCmd).Standalone()

	ec2_modifyInstanceMetadataOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("http-endpoint", "", "Enables or disables the HTTP metadata endpoint on your instances.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("http-protocol-ipv6", "", "Enables or disables the IPv6 endpoint for the instance metadata service.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("http-put-response-hop-limit", "", "The desired HTTP PUT response hop limit for instance metadata requests.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("http-tokens", "", "Indicates whether IMDSv2 is required.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().String("instance-metadata-tags", "", "Set to `enabled` to allow access to instance tags from the instance metadata.")
	ec2_modifyInstanceMetadataOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceMetadataOptionsCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceMetadataOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyInstanceMetadataOptionsCmd)
}
