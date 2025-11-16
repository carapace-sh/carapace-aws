package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes a Recycle Bin retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_deleteRuleCmd).Standalone()

	rbin_deleteRuleCmd.Flags().String("identifier", "", "The unique ID of the retention rule.")
	rbin_deleteRuleCmd.MarkFlagRequired("identifier")
	rbinCmd.AddCommand(rbin_deleteRuleCmd)
}
