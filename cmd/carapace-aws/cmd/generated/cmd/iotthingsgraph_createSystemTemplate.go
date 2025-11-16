package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_createSystemTemplateCmd = &cobra.Command{
	Use:   "create-system-template",
	Short: "Creates a system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_createSystemTemplateCmd).Standalone()

	iotthingsgraph_createSystemTemplateCmd.Flags().String("compatible-namespace-version", "", "The namespace version in which the system is to be created.")
	iotthingsgraph_createSystemTemplateCmd.Flags().String("definition", "", "The `DefinitionDocument` used to create the system.")
	iotthingsgraph_createSystemTemplateCmd.MarkFlagRequired("definition")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_createSystemTemplateCmd)
}
