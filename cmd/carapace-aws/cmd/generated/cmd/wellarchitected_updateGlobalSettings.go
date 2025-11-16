package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateGlobalSettingsCmd = &cobra.Command{
	Use:   "update-global-settings",
	Short: "Update whether the Amazon Web Services account is opted into organization sharing and discovery integration features.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateGlobalSettingsCmd).Standalone()

	wellarchitected_updateGlobalSettingsCmd.Flags().String("discovery-integration-status", "", "The status of discovery support settings.")
	wellarchitected_updateGlobalSettingsCmd.Flags().String("jira-configuration", "", "The status of Jira integration settings.")
	wellarchitected_updateGlobalSettingsCmd.Flags().String("organization-sharing-status", "", "The status of organization sharing settings.")
	wellarchitectedCmd.AddCommand(wellarchitected_updateGlobalSettingsCmd)
}
