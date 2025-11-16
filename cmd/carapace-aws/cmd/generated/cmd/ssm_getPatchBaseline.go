package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getPatchBaselineCmd = &cobra.Command{
	Use:   "get-patch-baseline",
	Short: "Retrieves information about a patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getPatchBaselineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getPatchBaselineCmd).Standalone()

		ssm_getPatchBaselineCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to retrieve.")
		ssm_getPatchBaselineCmd.MarkFlagRequired("baseline-id")
	})
	ssmCmd.AddCommand(ssm_getPatchBaselineCmd)
}
