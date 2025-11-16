package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstancePatchStatesCmd = &cobra.Command{
	Use:   "describe-instance-patch-states",
	Short: "Retrieves the high-level patch state of one or more managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstancePatchStatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeInstancePatchStatesCmd).Standalone()

		ssm_describeInstancePatchStatesCmd.Flags().String("instance-ids", "", "The ID of the managed node for which patch state information should be retrieved.")
		ssm_describeInstancePatchStatesCmd.Flags().String("max-results", "", "The maximum number of managed nodes to return (per page).")
		ssm_describeInstancePatchStatesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeInstancePatchStatesCmd.MarkFlagRequired("instance-ids")
	})
	ssmCmd.AddCommand(ssm_describeInstancePatchStatesCmd)
}
