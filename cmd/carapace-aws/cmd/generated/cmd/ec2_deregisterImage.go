package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deregisterImageCmd = &cobra.Command{
	Use:   "deregister-image",
	Short: "Deregisters the specified AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deregisterImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deregisterImageCmd).Standalone()

		ec2_deregisterImageCmd.Flags().Bool("delete-associated-snapshots", false, "Specifies whether to delete the snapshots associated with the AMI during deregistration.")
		ec2_deregisterImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterImageCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_deregisterImageCmd.Flags().Bool("no-delete-associated-snapshots", false, "Specifies whether to delete the snapshots associated with the AMI during deregistration.")
		ec2_deregisterImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterImageCmd.MarkFlagRequired("image-id")
		ec2_deregisterImageCmd.Flag("no-delete-associated-snapshots").Hidden = true
		ec2_deregisterImageCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deregisterImageCmd)
}
