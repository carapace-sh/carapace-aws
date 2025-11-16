package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importInstanceCmd = &cobra.Command{
	Use:   "import-instance",
	Short: "We recommend that you use the [`ImportImage`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImage.html) API instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_importInstanceCmd).Standalone()

		ec2_importInstanceCmd.Flags().String("description", "", "A description for the instance being imported.")
		ec2_importInstanceCmd.Flags().String("disk-images", "", "The disk image.")
		ec2_importInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importInstanceCmd.Flags().String("launch-specification", "", "The launch specification.")
		ec2_importInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importInstanceCmd.Flags().String("platform", "", "The instance operating system.")
		ec2_importInstanceCmd.Flag("no-dry-run").Hidden = true
		ec2_importInstanceCmd.MarkFlagRequired("platform")
	})
	ec2Cmd.AddCommand(ec2_importInstanceCmd)
}
