package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getGlobalSettingsCmd = &cobra.Command{
	Use:   "get-global-settings",
	Short: "Global settings for all workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getGlobalSettingsCmd).Standalone()

	wellarchitectedCmd.AddCommand(wellarchitected_getGlobalSettingsCmd)
}
