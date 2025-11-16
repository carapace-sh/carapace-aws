package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVolumeAttributeCmd = &cobra.Command{
	Use:   "modify-volume-attribute",
	Short: "Modifies a volume attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVolumeAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVolumeAttributeCmd).Standalone()

		ec2_modifyVolumeAttributeCmd.Flags().String("auto-enable-io", "", "Indicates whether the volume should be auto-enabled for I/O operations.")
		ec2_modifyVolumeAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVolumeAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVolumeAttributeCmd.Flags().String("volume-id", "", "The ID of the volume.")
		ec2_modifyVolumeAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVolumeAttributeCmd.MarkFlagRequired("volume-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVolumeAttributeCmd)
}
