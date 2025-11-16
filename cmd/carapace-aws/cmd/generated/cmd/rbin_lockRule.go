package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_lockRuleCmd = &cobra.Command{
	Use:   "lock-rule",
	Short: "Locks a Region-level retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_lockRuleCmd).Standalone()

	rbin_lockRuleCmd.Flags().String("identifier", "", "The unique ID of the retention rule.")
	rbin_lockRuleCmd.Flags().String("lock-configuration", "", "Information about the retention rule lock configuration.")
	rbin_lockRuleCmd.MarkFlagRequired("identifier")
	rbin_lockRuleCmd.MarkFlagRequired("lock-configuration")
	rbinCmd.AddCommand(rbin_lockRuleCmd)
}
