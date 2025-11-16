package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getEnvironmentTemplateCmd = &cobra.Command{
	Use:   "get-environment-template",
	Short: "Get detailed data for an environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getEnvironmentTemplateCmd).Standalone()

	proton_getEnvironmentTemplateCmd.Flags().String("name", "", "The name of the environment template that you want to get the detailed data for.")
	proton_getEnvironmentTemplateCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_getEnvironmentTemplateCmd)
}
