package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImageReferencesCmd = &cobra.Command{
	Use:   "describe-image-references",
	Short: "Describes your Amazon Web Services resources that are referencing the specified images.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImageReferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeImageReferencesCmd).Standalone()

		ec2_describeImageReferencesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeImageReferencesCmd.Flags().String("image-ids", "", "The IDs of the images to check for resource references.")
		ec2_describeImageReferencesCmd.Flags().Bool("include-all-resource-types", false, "Specifies whether to check all supported Amazon Web Services resource types for image references.")
		ec2_describeImageReferencesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeImageReferencesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeImageReferencesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeImageReferencesCmd.Flags().Bool("no-include-all-resource-types", false, "Specifies whether to check all supported Amazon Web Services resource types for image references.")
		ec2_describeImageReferencesCmd.Flags().String("resource-types", "", "The Amazon Web Services resource types to check for image references.")
		ec2_describeImageReferencesCmd.MarkFlagRequired("image-ids")
		ec2_describeImageReferencesCmd.Flag("no-dry-run").Hidden = true
		ec2_describeImageReferencesCmd.Flag("no-include-all-resource-types").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeImageReferencesCmd)
}
