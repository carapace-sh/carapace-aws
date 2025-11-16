package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVolumeAttributeCmd = &cobra.Command{
	Use:   "describe-volume-attribute",
	Short: "Describes the specified attribute of the specified volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVolumeAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVolumeAttributeCmd).Standalone()

		ec2_describeVolumeAttributeCmd.Flags().String("attribute", "", "The attribute of the volume.")
		ec2_describeVolumeAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumeAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVolumeAttributeCmd.Flags().String("volume-id", "", "The ID of the volume.")
		ec2_describeVolumeAttributeCmd.MarkFlagRequired("attribute")
		ec2_describeVolumeAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_describeVolumeAttributeCmd.MarkFlagRequired("volume-id")
	})
	ec2Cmd.AddCommand(ec2_describeVolumeAttributeCmd)
}
