package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deletePlacementGroupCmd = &cobra.Command{
	Use:   "delete-placement-group",
	Short: "Deletes the specified placement group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deletePlacementGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deletePlacementGroupCmd).Standalone()

		ec2_deletePlacementGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_deletePlacementGroupCmd.Flags().String("group-name", "", "The name of the placement group.")
		ec2_deletePlacementGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_deletePlacementGroupCmd.MarkFlagRequired("group-name")
		ec2_deletePlacementGroupCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deletePlacementGroupCmd)
}
