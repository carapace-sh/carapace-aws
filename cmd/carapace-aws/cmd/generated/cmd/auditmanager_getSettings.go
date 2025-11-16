package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getSettingsCmd = &cobra.Command{
	Use:   "get-settings",
	Short: "Gets the settings for a specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getSettingsCmd).Standalone()

		auditmanager_getSettingsCmd.Flags().String("attribute", "", "The list of setting attribute enum values.")
		auditmanager_getSettingsCmd.MarkFlagRequired("attribute")
	})
	auditmanagerCmd.AddCommand(auditmanager_getSettingsCmd)
}
