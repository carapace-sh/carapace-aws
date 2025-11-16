package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getSystemTemplateCmd = &cobra.Command{
	Use:   "get-system-template",
	Short: "Gets a system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getSystemTemplateCmd).Standalone()

	iotthingsgraph_getSystemTemplateCmd.Flags().String("id", "", "The ID of the system to get.")
	iotthingsgraph_getSystemTemplateCmd.Flags().String("revision-number", "", "The number that specifies the revision of the system to get.")
	iotthingsgraph_getSystemTemplateCmd.MarkFlagRequired("id")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getSystemTemplateCmd)
}
