package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates a new application that is the top-level node in a hierarchy of related cloud resource abstractions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_createApplicationCmd).Standalone()

		servicecatalogAppregistry_createApplicationCmd.Flags().String("client-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalogAppregistry_createApplicationCmd.Flags().String("description", "", "The description of the application.")
		servicecatalogAppregistry_createApplicationCmd.Flags().String("name", "", "The name of the application.")
		servicecatalogAppregistry_createApplicationCmd.Flags().String("tags", "", "Key-value pairs you can use to associate with the application.")
		servicecatalogAppregistry_createApplicationCmd.MarkFlagRequired("client-token")
		servicecatalogAppregistry_createApplicationCmd.MarkFlagRequired("name")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_createApplicationCmd)
}
