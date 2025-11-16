package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an application that is specified either by its application ID, name, or ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_deleteApplicationCmd).Standalone()

	servicecatalogAppregistry_deleteApplicationCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_deleteApplicationCmd.MarkFlagRequired("application")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_deleteApplicationCmd)
}
