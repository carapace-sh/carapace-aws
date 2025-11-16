package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_getConfigurationCmd = &cobra.Command{
	Use:   "get-configuration",
	Short: "Retrieves a `TagKey` configuration from an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_getConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_getConfigurationCmd).Standalone()

	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_getConfigurationCmd)
}
