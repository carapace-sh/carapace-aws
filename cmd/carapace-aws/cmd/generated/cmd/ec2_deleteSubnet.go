package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteSubnetCmd = &cobra.Command{
	Use:   "delete-subnet",
	Short: "Deletes the specified subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteSubnetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteSubnetCmd).Standalone()

		ec2_deleteSubnetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteSubnetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteSubnetCmd.Flags().String("subnet-id", "", "The ID of the subnet.")
		ec2_deleteSubnetCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteSubnetCmd.MarkFlagRequired("subnet-id")
	})
	ec2Cmd.AddCommand(ec2_deleteSubnetCmd)
}
