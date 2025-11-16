package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_unmonitorInstancesCmd = &cobra.Command{
	Use:   "unmonitor-instances",
	Short: "Disables detailed monitoring for a running instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_unmonitorInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_unmonitorInstancesCmd).Standalone()

		ec2_unmonitorInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_unmonitorInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
		ec2_unmonitorInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_unmonitorInstancesCmd.MarkFlagRequired("instance-ids")
		ec2_unmonitorInstancesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_unmonitorInstancesCmd)
}
