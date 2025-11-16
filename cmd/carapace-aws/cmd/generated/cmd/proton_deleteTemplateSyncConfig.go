package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteTemplateSyncConfigCmd = &cobra.Command{
	Use:   "delete-template-sync-config",
	Short: "Delete a template sync configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteTemplateSyncConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteTemplateSyncConfigCmd).Standalone()

		proton_deleteTemplateSyncConfigCmd.Flags().String("template-name", "", "The template name.")
		proton_deleteTemplateSyncConfigCmd.Flags().String("template-type", "", "The template type.")
		proton_deleteTemplateSyncConfigCmd.MarkFlagRequired("template-name")
		proton_deleteTemplateSyncConfigCmd.MarkFlagRequired("template-type")
	})
	protonCmd.AddCommand(proton_deleteTemplateSyncConfigCmd)
}
