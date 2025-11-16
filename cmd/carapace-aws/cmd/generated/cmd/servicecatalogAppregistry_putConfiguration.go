package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_putConfigurationCmd = &cobra.Command{
	Use:   "put-configuration",
	Short: "Associates a `TagKey` configuration to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_putConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_putConfigurationCmd).Standalone()

		servicecatalogAppregistry_putConfigurationCmd.Flags().String("configuration", "", "Associates a `TagKey` configuration to an account.")
		servicecatalogAppregistry_putConfigurationCmd.MarkFlagRequired("configuration")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_putConfigurationCmd)
}
