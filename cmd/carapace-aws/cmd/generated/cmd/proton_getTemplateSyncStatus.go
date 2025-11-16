package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getTemplateSyncStatusCmd = &cobra.Command{
	Use:   "get-template-sync-status",
	Short: "Get the status of a template sync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getTemplateSyncStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getTemplateSyncStatusCmd).Standalone()

		proton_getTemplateSyncStatusCmd.Flags().String("template-name", "", "The template name.")
		proton_getTemplateSyncStatusCmd.Flags().String("template-type", "", "The template type.")
		proton_getTemplateSyncStatusCmd.Flags().String("template-version", "", "The template major version.")
		proton_getTemplateSyncStatusCmd.MarkFlagRequired("template-name")
		proton_getTemplateSyncStatusCmd.MarkFlagRequired("template-type")
		proton_getTemplateSyncStatusCmd.MarkFlagRequired("template-version")
	})
	protonCmd.AddCommand(proton_getTemplateSyncStatusCmd)
}
