package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteEnvironmentTemplateCmd = &cobra.Command{
	Use:   "delete-environment-template",
	Short: "If no other major or minor versions of an environment template exist, delete the environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteEnvironmentTemplateCmd).Standalone()

	proton_deleteEnvironmentTemplateCmd.Flags().String("name", "", "The name of the environment template to delete.")
	proton_deleteEnvironmentTemplateCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_deleteEnvironmentTemplateCmd)
}
