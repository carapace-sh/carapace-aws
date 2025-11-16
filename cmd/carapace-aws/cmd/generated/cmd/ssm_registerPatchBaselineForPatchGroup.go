package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_registerPatchBaselineForPatchGroupCmd = &cobra.Command{
	Use:   "register-patch-baseline-for-patch-group",
	Short: "Registers a patch baseline for a patch group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_registerPatchBaselineForPatchGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_registerPatchBaselineForPatchGroupCmd).Standalone()

		ssm_registerPatchBaselineForPatchGroupCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to register with the patch group.")
		ssm_registerPatchBaselineForPatchGroupCmd.Flags().String("patch-group", "", "The name of the patch group to be registered with the patch baseline.")
		ssm_registerPatchBaselineForPatchGroupCmd.MarkFlagRequired("baseline-id")
		ssm_registerPatchBaselineForPatchGroupCmd.MarkFlagRequired("patch-group")
	})
	ssmCmd.AddCommand(ssm_registerPatchBaselineForPatchGroupCmd)
}
