package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_updateServiceSettingsCmd = &cobra.Command{
	Use:   "update-service-settings",
	Short: "Updates settings configured for Quick Setup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_updateServiceSettingsCmd).Standalone()

	ssmQuicksetup_updateServiceSettingsCmd.Flags().String("explorer-enabling-role-arn", "", "The IAM role used to enable Explorer.")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_updateServiceSettingsCmd)
}
