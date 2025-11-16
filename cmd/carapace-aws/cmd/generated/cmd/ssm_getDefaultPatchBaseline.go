package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getDefaultPatchBaselineCmd = &cobra.Command{
	Use:   "get-default-patch-baseline",
	Short: "Retrieves the default patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getDefaultPatchBaselineCmd).Standalone()

	ssm_getDefaultPatchBaselineCmd.Flags().String("operating-system", "", "Returns the default patch baseline for the specified operating system.")
	ssmCmd.AddCommand(ssm_getDefaultPatchBaselineCmd)
}
