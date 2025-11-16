package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceImageMetadataCmd = &cobra.Command{
	Use:   "describe-instance-image-metadata",
	Short: "Describes the AMI that was used to launch an instance, even if the AMI is deprecated, deregistered, made private (no longer public or shared with your account), or not allowed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceImageMetadataCmd).Standalone()

	ec2_describeInstanceImageMetadataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceImageMetadataCmd.Flags().String("filters", "", "The filters.")
	ec2_describeInstanceImageMetadataCmd.Flags().String("instance-ids", "", "The instance IDs.")
	ec2_describeInstanceImageMetadataCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceImageMetadataCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceImageMetadataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceImageMetadataCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceImageMetadataCmd)
}
