package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableImageCmd = &cobra.Command{
	Use:   "disable-image",
	Short: "Sets the AMI state to `disabled` and removes all launch permissions from the AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableImageCmd).Standalone()

		ec2_disableImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableImageCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_disableImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableImageCmd.MarkFlagRequired("image-id")
		ec2_disableImageCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disableImageCmd)
}
