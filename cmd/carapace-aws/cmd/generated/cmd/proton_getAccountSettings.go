package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Get detail data for Proton account-wide settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getAccountSettingsCmd).Standalone()

	})
	protonCmd.AddCommand(proton_getAccountSettingsCmd)
}
