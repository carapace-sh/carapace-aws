package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Retrieves metadata information about one of your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_getApplicationCmd).Standalone()

	servicecatalogAppregistry_getApplicationCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_getApplicationCmd.MarkFlagRequired("application")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_getApplicationCmd)
}
