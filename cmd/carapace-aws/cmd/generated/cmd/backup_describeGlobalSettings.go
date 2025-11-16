package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeGlobalSettingsCmd = &cobra.Command{
	Use:   "describe-global-settings",
	Short: "Describes whether the Amazon Web Services account is opted in to cross-account backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeGlobalSettingsCmd).Standalone()

	backupCmd.AddCommand(backup_describeGlobalSettingsCmd)
}
