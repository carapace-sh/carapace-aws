package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_monitorInstancesCmd = &cobra.Command{
	Use:   "monitor-instances",
	Short: "Enables detailed monitoring for a running instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_monitorInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_monitorInstancesCmd).Standalone()

		ec2_monitorInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_monitorInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		ec2_monitorInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_monitorInstancesCmd.MarkFlagRequired("instance-ids")
		ec2_monitorInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_monitorInstancesCmd)
}
