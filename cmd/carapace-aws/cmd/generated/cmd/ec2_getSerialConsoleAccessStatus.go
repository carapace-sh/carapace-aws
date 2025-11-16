package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getSerialConsoleAccessStatusCmd = &cobra.Command{
	Use:   "get-serial-console-access-status",
	Short: "Retrieves the access status of your account to the EC2 serial console of all instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getSerialConsoleAccessStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getSerialConsoleAccessStatusCmd).Standalone()

		ec2_getSerialConsoleAccessStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getSerialConsoleAccessStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getSerialConsoleAccessStatusCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getSerialConsoleAccessStatusCmd)
}
