package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_exportImageCmd = &cobra.Command{
	Use:   "export-image",
	Short: "Exports an Amazon Machine Image (AMI) to a VM file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_exportImageCmd).Standalone()

	ec2_exportImageCmd.Flags().String("client-token", "", "Token to enable idempotency for export image requests.")
	ec2_exportImageCmd.Flags().String("description", "", "A description of the image being exported.")
	ec2_exportImageCmd.Flags().String("disk-image-format", "", "The disk image format.")
	ec2_exportImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportImageCmd.Flags().String("image-id", "", "The ID of the image.")
	ec2_exportImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportImageCmd.Flags().String("role-name", "", "The name of the role that grants VM Import/Export permission to export images to your Amazon S3 bucket.")
	ec2_exportImageCmd.Flags().String("s3-export-location", "", "The Amazon S3 bucket for the destination image.")
	ec2_exportImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the export image task during creation.")
	ec2_exportImageCmd.MarkFlagRequired("disk-image-format")
	ec2_exportImageCmd.MarkFlagRequired("image-id")
	ec2_exportImageCmd.Flag("no-dry-run").Hidden = true
	ec2_exportImageCmd.MarkFlagRequired("s3-export-location")
	ec2Cmd.AddCommand(ec2_exportImageCmd)
}
