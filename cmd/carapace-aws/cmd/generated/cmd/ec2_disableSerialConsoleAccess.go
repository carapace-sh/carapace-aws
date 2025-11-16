package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableSerialConsoleAccessCmd = &cobra.Command{
	Use:   "disable-serial-console-access",
	Short: "Disables access to the EC2 serial console of all instances for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableSerialConsoleAccessCmd).Standalone()

	ec2_disableSerialConsoleAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableSerialConsoleAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableSerialConsoleAccessCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableSerialConsoleAccessCmd)
}
