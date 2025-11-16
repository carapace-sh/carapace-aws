package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_restoreImageFromRecycleBinCmd = &cobra.Command{
	Use:   "restore-image-from-recycle-bin",
	Short: "Restores an AMI from the Recycle Bin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_restoreImageFromRecycleBinCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_restoreImageFromRecycleBinCmd).Standalone()

		ec2_restoreImageFromRecycleBinCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreImageFromRecycleBinCmd.Flags().String("image-id", "", "The ID of the AMI to restore.")
		ec2_restoreImageFromRecycleBinCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreImageFromRecycleBinCmd.MarkFlagRequired("image-id")
		ec2_restoreImageFromRecycleBinCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_restoreImageFromRecycleBinCmd)
}
