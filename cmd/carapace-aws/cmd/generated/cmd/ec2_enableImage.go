package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableImageCmd = &cobra.Command{
	Use:   "enable-image",
	Short: "Re-enables a disabled AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableImageCmd).Standalone()

	ec2_enableImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableImageCmd.Flags().String("image-id", "", "The ID of the AMI.")
	ec2_enableImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableImageCmd.MarkFlagRequired("image-id")
	ec2_enableImageCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_enableImageCmd)
}
