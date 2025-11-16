package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableImageDeregistrationProtectionCmd = &cobra.Command{
	Use:   "enable-image-deregistration-protection",
	Short: "Enables deregistration protection for an AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableImageDeregistrationProtectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableImageDeregistrationProtectionCmd).Standalone()

		ec2_enableImageDeregistrationProtectionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableImageDeregistrationProtectionCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_enableImageDeregistrationProtectionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableImageDeregistrationProtectionCmd.Flags().Bool("no-with-cooldown", false, "If `true`, enforces deregistration protection for 24 hours after deregistration protection is disabled.")
		ec2_enableImageDeregistrationProtectionCmd.Flags().Bool("with-cooldown", false, "If `true`, enforces deregistration protection for 24 hours after deregistration protection is disabled.")
		ec2_enableImageDeregistrationProtectionCmd.MarkFlagRequired("image-id")
		ec2_enableImageDeregistrationProtectionCmd.Flag("no-dry-run").Hidden = true
		ec2_enableImageDeregistrationProtectionCmd.Flag("no-with-cooldown").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableImageDeregistrationProtectionCmd)
}
