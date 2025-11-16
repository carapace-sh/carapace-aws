package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rebootInstancesCmd = &cobra.Command{
	Use:   "reboot-instances",
	Short: "Requests a reboot of the specified instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rebootInstancesCmd).Standalone()

	ec2_rebootInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_rebootInstancesCmd.Flags().String("instance-ids", "", "The instance IDs.")
	ec2_rebootInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_rebootInstancesCmd.MarkFlagRequired("instance-ids")
	ec2_rebootInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_rebootInstancesCmd)
}
