package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateEnvironmentTemplateCmd = &cobra.Command{
	Use:   "update-environment-template",
	Short: "Update an environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateEnvironmentTemplateCmd).Standalone()

	proton_updateEnvironmentTemplateCmd.Flags().String("description", "", "A description of the environment template update.")
	proton_updateEnvironmentTemplateCmd.Flags().String("display-name", "", "The name of the environment template to update as displayed in the developer interface.")
	proton_updateEnvironmentTemplateCmd.Flags().String("name", "", "The name of the environment template to update.")
	proton_updateEnvironmentTemplateCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_updateEnvironmentTemplateCmd)
}
