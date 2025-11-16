package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates an existing application with new attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_updateApplicationCmd).Standalone()

		servicecatalogAppregistry_updateApplicationCmd.Flags().String("application", "", "The name, ID, or ARN of the application that will be updated.")
		servicecatalogAppregistry_updateApplicationCmd.Flags().String("description", "", "The new description of the application.")
		servicecatalogAppregistry_updateApplicationCmd.Flags().String("name", "", "Deprecated: The new name of the application.")
		servicecatalogAppregistry_updateApplicationCmd.MarkFlagRequired("application")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_updateApplicationCmd)
}
