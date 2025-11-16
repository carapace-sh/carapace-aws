package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImagesCmd = &cobra.Command{
	Use:   "describe-images",
	Short: "Describes the specified images (AMIs, AKIs, and ARIs) available to you or all of the images available to you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImagesCmd).Standalone()

	ec2_describeImagesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImagesCmd.Flags().String("executable-users", "", "Scopes the images by users with explicit launch permissions.")
	ec2_describeImagesCmd.Flags().String("filters", "", "The filters.")
	ec2_describeImagesCmd.Flags().String("image-ids", "", "The image IDs.")
	ec2_describeImagesCmd.Flags().Bool("include-deprecated", false, "Specifies whether to include deprecated AMIs.")
	ec2_describeImagesCmd.Flags().Bool("include-disabled", false, "Specifies whether to include disabled AMIs.")
	ec2_describeImagesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeImagesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeImagesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImagesCmd.Flags().Bool("no-include-deprecated", false, "Specifies whether to include deprecated AMIs.")
	ec2_describeImagesCmd.Flags().Bool("no-include-disabled", false, "Specifies whether to include disabled AMIs.")
	ec2_describeImagesCmd.Flags().String("owners", "", "Scopes the results to images with the specified owners.")
	ec2_describeImagesCmd.Flag("no-dry-run").Hidden = true
	ec2_describeImagesCmd.Flag("no-include-deprecated").Hidden = true
	ec2_describeImagesCmd.Flag("no-include-disabled").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImagesCmd)
}
