package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_stopInstancesCmd = &cobra.Command{
	Use:   "stop-instances",
	Short: "Stops an Amazon EBS-backed instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_stopInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_stopInstancesCmd).Standalone()

		ec2_stopInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_stopInstancesCmd.Flags().Bool("force", false, "Forces the instance to stop.")
		ec2_stopInstancesCmd.Flags().Bool("hibernate", false, "Hibernates the instance if the instance was enabled for hibernation at launch.")
		ec2_stopInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		ec2_stopInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_stopInstancesCmd.Flags().Bool("no-force", false, "Forces the instance to stop.")
		ec2_stopInstancesCmd.Flags().Bool("no-hibernate", false, "Hibernates the instance if the instance was enabled for hibernation at launch.")
		ec2_stopInstancesCmd.Flags().Bool("no-skip-os-shutdown", false, "Specifies whether to bypass the graceful OS shutdown process when the instance is stopped.")
		ec2_stopInstancesCmd.Flags().Bool("skip-os-shutdown", false, "Specifies whether to bypass the graceful OS shutdown process when the instance is stopped.")
		ec2_stopInstancesCmd.MarkFlagRequired("instance-ids")
		ec2_stopInstancesCmd.Flag("no-dry-run").Hidden = true
		ec2_stopInstancesCmd.Flag("no-force").Hidden = true
		ec2_stopInstancesCmd.Flag("no-hibernate").Hidden = true
		ec2_stopInstancesCmd.Flag("no-skip-os-shutdown").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_stopInstancesCmd)
}
