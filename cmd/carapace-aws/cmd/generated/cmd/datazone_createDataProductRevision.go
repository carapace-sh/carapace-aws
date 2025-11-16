package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createDataProductRevisionCmd = &cobra.Command{
	Use:   "create-data-product-revision",
	Short: "Creates a data product revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createDataProductRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createDataProductRevisionCmd).Standalone()

		datazone_createDataProductRevisionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createDataProductRevisionCmd.Flags().String("description", "", "The description of the data product revision.")
		datazone_createDataProductRevisionCmd.Flags().String("domain-identifier", "", "The ID of the domain where the data product revision is created.")
		datazone_createDataProductRevisionCmd.Flags().String("forms-input", "", "The metadata forms of the data product revision.")
		datazone_createDataProductRevisionCmd.Flags().String("glossary-terms", "", "The glossary terms of the data product revision.")
		datazone_createDataProductRevisionCmd.Flags().String("identifier", "", "The ID of the data product revision.")
		datazone_createDataProductRevisionCmd.Flags().String("items", "", "The data assets of the data product revision.")
		datazone_createDataProductRevisionCmd.Flags().String("name", "", "The name of the data product revision.")
		datazone_createDataProductRevisionCmd.MarkFlagRequired("domain-identifier")
		datazone_createDataProductRevisionCmd.MarkFlagRequired("identifier")
		datazone_createDataProductRevisionCmd.MarkFlagRequired("name")
	})
	datazoneCmd.AddCommand(datazone_createDataProductRevisionCmd)
}
