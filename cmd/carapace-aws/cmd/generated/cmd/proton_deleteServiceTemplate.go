package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteServiceTemplateCmd = &cobra.Command{
	Use:   "delete-service-template",
	Short: "If no other major or minor versions of the service template exist, delete the service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteServiceTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteServiceTemplateCmd).Standalone()

		proton_deleteServiceTemplateCmd.Flags().String("name", "", "The name of the service template to delete.")
		proton_deleteServiceTemplateCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_deleteServiceTemplateCmd)
}
