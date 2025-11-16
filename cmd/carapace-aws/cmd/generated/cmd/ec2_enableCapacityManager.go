package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableCapacityManagerCmd = &cobra.Command{
	Use:   "enable-capacity-manager",
	Short: "Enables EC2 Capacity Manager for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableCapacityManagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableCapacityManagerCmd).Standalone()

		ec2_enableCapacityManagerCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_enableCapacityManagerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableCapacityManagerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableCapacityManagerCmd.Flags().Bool("no-organizations-access", false, "Specifies whether to enable cross-account access for Amazon Web Services Organizations.")
		ec2_enableCapacityManagerCmd.Flags().Bool("organizations-access", false, "Specifies whether to enable cross-account access for Amazon Web Services Organizations.")
		ec2_enableCapacityManagerCmd.Flag("no-dry-run").Hidden = true
		ec2_enableCapacityManagerCmd.Flag("no-organizations-access").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableCapacityManagerCmd)
}
