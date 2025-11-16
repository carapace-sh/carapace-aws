package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createDataProductCmd = &cobra.Command{
	Use:   "create-data-product",
	Short: "Creates a data product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createDataProductCmd).Standalone()

	datazone_createDataProductCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_createDataProductCmd.Flags().String("description", "", "The description of the data product.")
	datazone_createDataProductCmd.Flags().String("domain-identifier", "", "The ID of the domain where the data product is created.")
	datazone_createDataProductCmd.Flags().String("forms-input", "", "The metadata forms of the data product.")
	datazone_createDataProductCmd.Flags().String("glossary-terms", "", "The glossary terms of the data product.")
	datazone_createDataProductCmd.Flags().String("items", "", "The data assets of the data product.")
	datazone_createDataProductCmd.Flags().String("name", "", "The name of the data product.")
	datazone_createDataProductCmd.Flags().String("owning-project-identifier", "", "The ID of the owning project of the data product.")
	datazone_createDataProductCmd.MarkFlagRequired("domain-identifier")
	datazone_createDataProductCmd.MarkFlagRequired("name")
	datazone_createDataProductCmd.MarkFlagRequired("owning-project-identifier")
	datazoneCmd.AddCommand(datazone_createDataProductCmd)
}
