package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeEffectivePatchesForPatchBaselineCmd = &cobra.Command{
	Use:   "describe-effective-patches-for-patch-baseline",
	Short: "Retrieves the current effective patches (the patch and the approval state) for the specified patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeEffectivePatchesForPatchBaselineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeEffectivePatchesForPatchBaselineCmd).Standalone()

		ssm_describeEffectivePatchesForPatchBaselineCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to retrieve the effective patches for.")
		ssm_describeEffectivePatchesForPatchBaselineCmd.Flags().String("max-results", "", "The maximum number of patches to return (per page).")
		ssm_describeEffectivePatchesForPatchBaselineCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeEffectivePatchesForPatchBaselineCmd.MarkFlagRequired("baseline-id")
	})
	ssmCmd.AddCommand(ssm_describeEffectivePatchesForPatchBaselineCmd)
}
