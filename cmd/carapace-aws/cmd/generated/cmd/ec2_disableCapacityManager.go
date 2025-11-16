package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableCapacityManagerCmd = &cobra.Command{
	Use:   "disable-capacity-manager",
	Short: "Disables EC2 Capacity Manager for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableCapacityManagerCmd).Standalone()

	ec2_disableCapacityManagerCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_disableCapacityManagerCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableCapacityManagerCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableCapacityManagerCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableCapacityManagerCmd)
}
