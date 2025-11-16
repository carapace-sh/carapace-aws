package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deprecateSystemTemplateCmd = &cobra.Command{
	Use:   "deprecate-system-template",
	Short: "Deprecates the specified system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deprecateSystemTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_deprecateSystemTemplateCmd).Standalone()

		iotthingsgraph_deprecateSystemTemplateCmd.Flags().String("id", "", "The ID of the system to delete.")
		iotthingsgraph_deprecateSystemTemplateCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deprecateSystemTemplateCmd)
}
