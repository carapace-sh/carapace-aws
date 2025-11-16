package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getProviderServiceCmd = &cobra.Command{
	Use:   "get-provider-service",
	Short: "Returns the `ProviderService` of a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getProviderServiceCmd).Standalone()

	entityresolution_getProviderServiceCmd.Flags().String("provider-name", "", "The name of the provider.")
	entityresolution_getProviderServiceCmd.Flags().String("provider-service-name", "", "The ARN (Amazon Resource Name) of the product that the provider service provides.")
	entityresolution_getProviderServiceCmd.MarkFlagRequired("provider-name")
	entityresolution_getProviderServiceCmd.MarkFlagRequired("provider-service-name")
	entityresolutionCmd.AddCommand(entityresolution_getProviderServiceCmd)
}
