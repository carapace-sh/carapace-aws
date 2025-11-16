package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getConsoleScreenshotCmd = &cobra.Command{
	Use:   "get-console-screenshot",
	Short: "Retrieve a JPG-format screenshot of a running instance to help with troubleshooting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getConsoleScreenshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getConsoleScreenshotCmd).Standalone()

		ec2_getConsoleScreenshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getConsoleScreenshotCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_getConsoleScreenshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getConsoleScreenshotCmd.Flags().Bool("no-wake-up", false, "When set to `true`, acts as keystroke input and wakes up an instance that's in standby or \"sleep\" mode.")
		ec2_getConsoleScreenshotCmd.Flags().Bool("wake-up", false, "When set to `true`, acts as keystroke input and wakes up an instance that's in standby or \"sleep\" mode.")
		ec2_getConsoleScreenshotCmd.MarkFlagRequired("instance-id")
		ec2_getConsoleScreenshotCmd.Flag("no-dry-run").Hidden = true
		ec2_getConsoleScreenshotCmd.Flag("no-wake-up").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getConsoleScreenshotCmd)
}
