package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listProviderServicesCmd = &cobra.Command{
	Use:   "list-provider-services",
	Short: "Returns a list of all the `ProviderServices` that are available in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listProviderServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_listProviderServicesCmd).Standalone()

		entityresolution_listProviderServicesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		entityresolution_listProviderServicesCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
		entityresolution_listProviderServicesCmd.Flags().String("provider-name", "", "The name of the provider.")
	})
	entityresolutionCmd.AddCommand(entityresolution_listProviderServicesCmd)
}
