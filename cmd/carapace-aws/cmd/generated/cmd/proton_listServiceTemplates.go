package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServiceTemplatesCmd = &cobra.Command{
	Use:   "list-service-templates",
	Short: "List service templates with detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServiceTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServiceTemplatesCmd).Standalone()

		proton_listServiceTemplatesCmd.Flags().String("max-results", "", "The maximum number of service templates to list.")
		proton_listServiceTemplatesCmd.Flags().String("next-token", "", "A token that indicates the location of the next service template in the array of service templates, after the list of service templates previously requested.")
	})
	protonCmd.AddCommand(proton_listServiceTemplatesCmd)
}
