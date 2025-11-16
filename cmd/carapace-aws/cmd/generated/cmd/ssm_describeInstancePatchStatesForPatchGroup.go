package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstancePatchStatesForPatchGroupCmd = &cobra.Command{
	Use:   "describe-instance-patch-states-for-patch-group",
	Short: "Retrieves the high-level patch state for the managed nodes in the specified patch group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstancePatchStatesForPatchGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeInstancePatchStatesForPatchGroupCmd).Standalone()

		ssm_describeInstancePatchStatesForPatchGroupCmd.Flags().String("filters", "", "Each entry in the array is a structure containing:")
		ssm_describeInstancePatchStatesForPatchGroupCmd.Flags().String("max-results", "", "The maximum number of patches to return (per page).")
		ssm_describeInstancePatchStatesForPatchGroupCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeInstancePatchStatesForPatchGroupCmd.Flags().String("patch-group", "", "The name of the patch group for which the patch state information should be retrieved.")
		ssm_describeInstancePatchStatesForPatchGroupCmd.MarkFlagRequired("patch-group")
	})
	ssmCmd.AddCommand(ssm_describeInstancePatchStatesForPatchGroupCmd)
}
