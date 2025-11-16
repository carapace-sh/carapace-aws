package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFastLaunchImagesCmd = &cobra.Command{
	Use:   "describe-fast-launch-images",
	Short: "Describe details for Windows AMIs that are configured for Windows fast launch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFastLaunchImagesCmd).Standalone()

	ec2_describeFastLaunchImagesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFastLaunchImagesCmd.Flags().String("filters", "", "Use the following filters to streamline results.")
	ec2_describeFastLaunchImagesCmd.Flags().String("image-ids", "", "Specify one or more Windows AMI image IDs for the request.")
	ec2_describeFastLaunchImagesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeFastLaunchImagesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeFastLaunchImagesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFastLaunchImagesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeFastLaunchImagesCmd)
}
