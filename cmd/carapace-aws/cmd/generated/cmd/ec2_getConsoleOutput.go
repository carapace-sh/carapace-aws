package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getConsoleOutputCmd = &cobra.Command{
	Use:   "get-console-output",
	Short: "Gets the console output for the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getConsoleOutputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getConsoleOutputCmd).Standalone()

		ec2_getConsoleOutputCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getConsoleOutputCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_getConsoleOutputCmd.Flags().Bool("latest", false, "When enabled, retrieves the latest console output for the instance.")
		ec2_getConsoleOutputCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getConsoleOutputCmd.Flags().Bool("no-latest", false, "When enabled, retrieves the latest console output for the instance.")
		ec2_getConsoleOutputCmd.MarkFlagRequired("instance-id")
		ec2_getConsoleOutputCmd.Flag("no-dry-run").Hidden = true
		ec2_getConsoleOutputCmd.Flag("no-latest").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getConsoleOutputCmd)
}
