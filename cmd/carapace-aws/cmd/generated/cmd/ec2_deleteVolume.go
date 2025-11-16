package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVolumeCmd = &cobra.Command{
	Use:   "delete-volume",
	Short: "Deletes the specified EBS volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVolumeCmd).Standalone()

	ec2_deleteVolumeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVolumeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVolumeCmd.Flags().String("volume-id", "", "The ID of the volume.")
	ec2_deleteVolumeCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteVolumeCmd.MarkFlagRequired("volume-id")
	ec2Cmd.AddCommand(ec2_deleteVolumeCmd)
}
