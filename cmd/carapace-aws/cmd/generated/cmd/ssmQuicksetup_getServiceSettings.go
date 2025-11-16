package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_getServiceSettingsCmd = &cobra.Command{
	Use:   "get-service-settings",
	Short: "Returns settings configured for Quick Setup in the requesting Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_getServiceSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_getServiceSettingsCmd).Standalone()

	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_getServiceSettingsCmd)
}
