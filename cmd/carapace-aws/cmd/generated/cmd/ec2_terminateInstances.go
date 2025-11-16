package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_terminateInstancesCmd = &cobra.Command{
	Use:   "terminate-instances",
	Short: "Terminates (deletes) the specified instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_terminateInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_terminateInstancesCmd).Standalone()

		ec2_terminateInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_terminateInstancesCmd.Flags().Bool("force", false, "Forces the instances to terminate.")
		ec2_terminateInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		ec2_terminateInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_terminateInstancesCmd.Flags().Bool("no-force", false, "Forces the instances to terminate.")
		ec2_terminateInstancesCmd.Flags().Bool("no-skip-os-shutdown", false, "Specifies whether to bypass the graceful OS shutdown process when the instance is terminated.")
		ec2_terminateInstancesCmd.Flags().Bool("skip-os-shutdown", false, "Specifies whether to bypass the graceful OS shutdown process when the instance is terminated.")
		ec2_terminateInstancesCmd.MarkFlagRequired("instance-ids")
		ec2_terminateInstancesCmd.Flag("no-dry-run").Hidden = true
		ec2_terminateInstancesCmd.Flag("no-force").Hidden = true
		ec2_terminateInstancesCmd.Flag("no-skip-os-shutdown").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_terminateInstancesCmd)
}
