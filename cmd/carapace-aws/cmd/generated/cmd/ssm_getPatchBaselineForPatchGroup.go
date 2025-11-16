package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getPatchBaselineForPatchGroupCmd = &cobra.Command{
	Use:   "get-patch-baseline-for-patch-group",
	Short: "Retrieves the patch baseline that should be used for the specified patch group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getPatchBaselineForPatchGroupCmd).Standalone()

	ssm_getPatchBaselineForPatchGroupCmd.Flags().String("operating-system", "", "Returns the operating system rule specified for patch groups using the patch baseline.")
	ssm_getPatchBaselineForPatchGroupCmd.Flags().String("patch-group", "", "The name of the patch group whose patch baseline should be retrieved.")
	ssm_getPatchBaselineForPatchGroupCmd.MarkFlagRequired("patch-group")
	ssmCmd.AddCommand(ssm_getPatchBaselineForPatchGroupCmd)
}
