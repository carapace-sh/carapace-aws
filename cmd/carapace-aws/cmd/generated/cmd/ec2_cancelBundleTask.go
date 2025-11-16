package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelBundleTaskCmd = &cobra.Command{
	Use:   "cancel-bundle-task",
	Short: "Cancels a bundling operation for an instance store-backed Windows instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelBundleTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelBundleTaskCmd).Standalone()

		ec2_cancelBundleTaskCmd.Flags().String("bundle-id", "", "The ID of the bundle task.")
		ec2_cancelBundleTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelBundleTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelBundleTaskCmd.MarkFlagRequired("bundle-id")
		ec2_cancelBundleTaskCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_cancelBundleTaskCmd)
}
