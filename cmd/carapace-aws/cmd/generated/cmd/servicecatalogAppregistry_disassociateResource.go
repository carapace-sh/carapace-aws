package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_disassociateResourceCmd = &cobra.Command{
	Use:   "disassociate-resource",
	Short: "Disassociates a resource from application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_disassociateResourceCmd).Standalone()

	servicecatalogAppregistry_disassociateResourceCmd.Flags().String("application", "", "The name or ID of the application.")
	servicecatalogAppregistry_disassociateResourceCmd.Flags().String("resource", "", "The name or ID of the resource.")
	servicecatalogAppregistry_disassociateResourceCmd.Flags().String("resource-type", "", "The type of the resource that is being disassociated.")
	servicecatalogAppregistry_disassociateResourceCmd.MarkFlagRequired("application")
	servicecatalogAppregistry_disassociateResourceCmd.MarkFlagRequired("resource")
	servicecatalogAppregistry_disassociateResourceCmd.MarkFlagRequired("resource-type")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_disassociateResourceCmd)
}
