package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_listQuickSetupTypesCmd = &cobra.Command{
	Use:   "list-quick-setup-types",
	Short: "Returns the available Quick Setup types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_listQuickSetupTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_listQuickSetupTypesCmd).Standalone()

	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_listQuickSetupTypesCmd)
}
