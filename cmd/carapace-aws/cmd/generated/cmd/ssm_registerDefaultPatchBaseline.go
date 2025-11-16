package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_registerDefaultPatchBaselineCmd = &cobra.Command{
	Use:   "register-default-patch-baseline",
	Short: "Defines the default patch baseline for the relevant operating system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_registerDefaultPatchBaselineCmd).Standalone()

	ssm_registerDefaultPatchBaselineCmd.Flags().String("baseline-id", "", "The ID of the patch baseline that should be the default patch baseline.")
	ssm_registerDefaultPatchBaselineCmd.MarkFlagRequired("baseline-id")
	ssmCmd.AddCommand(ssm_registerDefaultPatchBaselineCmd)
}
