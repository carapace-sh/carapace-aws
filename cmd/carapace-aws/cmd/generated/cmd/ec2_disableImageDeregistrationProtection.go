package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableImageDeregistrationProtectionCmd = &cobra.Command{
	Use:   "disable-image-deregistration-protection",
	Short: "Disables deregistration protection for an AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableImageDeregistrationProtectionCmd).Standalone()

	ec2_disableImageDeregistrationProtectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableImageDeregistrationProtectionCmd.Flags().String("image-id", "", "The ID of the AMI.")
	ec2_disableImageDeregistrationProtectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableImageDeregistrationProtectionCmd.MarkFlagRequired("image-id")
	ec2_disableImageDeregistrationProtectionCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableImageDeregistrationProtectionCmd)
}
