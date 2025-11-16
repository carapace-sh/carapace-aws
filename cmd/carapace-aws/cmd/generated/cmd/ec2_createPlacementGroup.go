package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createPlacementGroupCmd = &cobra.Command{
	Use:   "create-placement-group",
	Short: "Creates a placement group in which to launch instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createPlacementGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createPlacementGroupCmd).Standalone()

		ec2_createPlacementGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_createPlacementGroupCmd.Flags().String("group-name", "", "A name for the placement group.")
		ec2_createPlacementGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_createPlacementGroupCmd.Flags().String("partition-count", "", "The number of partitions.")
		ec2_createPlacementGroupCmd.Flags().String("spread-level", "", "Determines how placement groups spread instances.")
		ec2_createPlacementGroupCmd.Flags().String("strategy", "", "The placement strategy.")
		ec2_createPlacementGroupCmd.Flags().String("tag-specifications", "", "The tags to apply to the new placement group.")
		ec2_createPlacementGroupCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createPlacementGroupCmd)
}
