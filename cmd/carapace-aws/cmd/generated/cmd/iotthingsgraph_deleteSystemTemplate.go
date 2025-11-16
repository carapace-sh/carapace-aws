package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deleteSystemTemplateCmd = &cobra.Command{
	Use:   "delete-system-template",
	Short: "Deletes a system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deleteSystemTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_deleteSystemTemplateCmd).Standalone()

		iotthingsgraph_deleteSystemTemplateCmd.Flags().String("id", "", "The ID of the system to be deleted.")
		iotthingsgraph_deleteSystemTemplateCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deleteSystemTemplateCmd)
}
