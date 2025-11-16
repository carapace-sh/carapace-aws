package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceTemplateCmd = &cobra.Command{
	Use:   "get-service-template",
	Short: "Get detailed data for a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceTemplateCmd).Standalone()

	proton_getServiceTemplateCmd.Flags().String("name", "", "The name of the service template that you want to get detailed data for.")
	proton_getServiceTemplateCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_getServiceTemplateCmd)
}
