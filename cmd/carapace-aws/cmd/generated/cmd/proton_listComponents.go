package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "List components with summary data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listComponentsCmd).Standalone()

		proton_listComponentsCmd.Flags().String("environment-name", "", "The name of an environment for result list filtering.")
		proton_listComponentsCmd.Flags().String("max-results", "", "The maximum number of components to list.")
		proton_listComponentsCmd.Flags().String("next-token", "", "A token that indicates the location of the next component in the array of components, after the list of components that was previously requested.")
		proton_listComponentsCmd.Flags().String("service-instance-name", "", "The name of a service instance for result list filtering.")
		proton_listComponentsCmd.Flags().String("service-name", "", "The name of a service for result list filtering.")
	})
	protonCmd.AddCommand(proton_listComponentsCmd)
}
