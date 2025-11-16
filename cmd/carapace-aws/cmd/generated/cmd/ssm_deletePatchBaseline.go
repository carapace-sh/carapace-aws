package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deletePatchBaselineCmd = &cobra.Command{
	Use:   "delete-patch-baseline",
	Short: "Deletes a patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deletePatchBaselineCmd).Standalone()

	ssm_deletePatchBaselineCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to delete.")
	ssm_deletePatchBaselineCmd.MarkFlagRequired("baseline-id")
	ssmCmd.AddCommand(ssm_deletePatchBaselineCmd)
}
