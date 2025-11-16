package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getTemplateSyncConfigCmd = &cobra.Command{
	Use:   "get-template-sync-config",
	Short: "Get detail data for a template sync configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getTemplateSyncConfigCmd).Standalone()

	proton_getTemplateSyncConfigCmd.Flags().String("template-name", "", "The template name.")
	proton_getTemplateSyncConfigCmd.Flags().String("template-type", "", "The template type.")
	proton_getTemplateSyncConfigCmd.MarkFlagRequired("template-name")
	proton_getTemplateSyncConfigCmd.MarkFlagRequired("template-type")
	protonCmd.AddCommand(proton_getTemplateSyncConfigCmd)
}
