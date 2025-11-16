package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_unlockRuleCmd = &cobra.Command{
	Use:   "unlock-rule",
	Short: "Unlocks a retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_unlockRuleCmd).Standalone()

	rbin_unlockRuleCmd.Flags().String("identifier", "", "The unique ID of the retention rule.")
	rbin_unlockRuleCmd.MarkFlagRequired("identifier")
	rbinCmd.AddCommand(rbin_unlockRuleCmd)
}
