package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateTemplateSyncConfigCmd = &cobra.Command{
	Use:   "update-template-sync-config",
	Short: "Update template sync configuration parameters, except for the `templateName` and `templateType`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateTemplateSyncConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateTemplateSyncConfigCmd).Standalone()

		proton_updateTemplateSyncConfigCmd.Flags().String("branch", "", "The repository branch for your template.")
		proton_updateTemplateSyncConfigCmd.Flags().String("repository-name", "", "The repository name (for example, `myrepos/myrepo`).")
		proton_updateTemplateSyncConfigCmd.Flags().String("repository-provider", "", "The repository provider.")
		proton_updateTemplateSyncConfigCmd.Flags().String("subdirectory", "", "A subdirectory path to your template bundle version.")
		proton_updateTemplateSyncConfigCmd.Flags().String("template-name", "", "The synced template name.")
		proton_updateTemplateSyncConfigCmd.Flags().String("template-type", "", "The synced template type.")
		proton_updateTemplateSyncConfigCmd.MarkFlagRequired("branch")
		proton_updateTemplateSyncConfigCmd.MarkFlagRequired("repository-name")
		proton_updateTemplateSyncConfigCmd.MarkFlagRequired("repository-provider")
		proton_updateTemplateSyncConfigCmd.MarkFlagRequired("template-name")
		proton_updateTemplateSyncConfigCmd.MarkFlagRequired("template-type")
	})
	protonCmd.AddCommand(proton_updateTemplateSyncConfigCmd)
}
