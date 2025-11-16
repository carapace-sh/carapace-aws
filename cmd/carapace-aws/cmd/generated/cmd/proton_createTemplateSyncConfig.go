package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createTemplateSyncConfigCmd = &cobra.Command{
	Use:   "create-template-sync-config",
	Short: "Set up a template to create new template versions automatically by tracking a linked repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createTemplateSyncConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_createTemplateSyncConfigCmd).Standalone()

		proton_createTemplateSyncConfigCmd.Flags().String("branch", "", "The repository branch for your template.")
		proton_createTemplateSyncConfigCmd.Flags().String("repository-name", "", "The repository name (for example, `myrepos/myrepo`).")
		proton_createTemplateSyncConfigCmd.Flags().String("repository-provider", "", "The provider type for your repository.")
		proton_createTemplateSyncConfigCmd.Flags().String("subdirectory", "", "A repository subdirectory path to your template bundle directory.")
		proton_createTemplateSyncConfigCmd.Flags().String("template-name", "", "The name of your registered template.")
		proton_createTemplateSyncConfigCmd.Flags().String("template-type", "", "The type of the registered template.")
		proton_createTemplateSyncConfigCmd.MarkFlagRequired("branch")
		proton_createTemplateSyncConfigCmd.MarkFlagRequired("repository-name")
		proton_createTemplateSyncConfigCmd.MarkFlagRequired("repository-provider")
		proton_createTemplateSyncConfigCmd.MarkFlagRequired("template-name")
		proton_createTemplateSyncConfigCmd.MarkFlagRequired("template-type")
	})
	protonCmd.AddCommand(proton_createTemplateSyncConfigCmd)
}
