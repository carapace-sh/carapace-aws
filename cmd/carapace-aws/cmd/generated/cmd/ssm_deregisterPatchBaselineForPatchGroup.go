package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deregisterPatchBaselineForPatchGroupCmd = &cobra.Command{
	Use:   "deregister-patch-baseline-for-patch-group",
	Short: "Removes a patch group from a patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deregisterPatchBaselineForPatchGroupCmd).Standalone()

	ssm_deregisterPatchBaselineForPatchGroupCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to deregister the patch group from.")
	ssm_deregisterPatchBaselineForPatchGroupCmd.Flags().String("patch-group", "", "The name of the patch group that should be deregistered from the patch baseline.")
	ssm_deregisterPatchBaselineForPatchGroupCmd.MarkFlagRequired("baseline-id")
	ssm_deregisterPatchBaselineForPatchGroupCmd.MarkFlagRequired("patch-group")
	ssmCmd.AddCommand(ssm_deregisterPatchBaselineForPatchGroupCmd)
}
