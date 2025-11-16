package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_bundleInstanceCmd = &cobra.Command{
	Use:   "bundle-instance",
	Short: "Bundles an Amazon instance store-backed Windows instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_bundleInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_bundleInstanceCmd).Standalone()

		ec2_bundleInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_bundleInstanceCmd.Flags().String("instance-id", "", "The ID of the instance to bundle.")
		ec2_bundleInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_bundleInstanceCmd.Flags().String("storage", "", "The bucket in which to store the AMI.")
		ec2_bundleInstanceCmd.MarkFlagRequired("instance-id")
		ec2_bundleInstanceCmd.Flag("no-dry-run").Hidden = true
		ec2_bundleInstanceCmd.MarkFlagRequired("storage")
	})
	ec2Cmd.AddCommand(ec2_bundleInstanceCmd)
}
