package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates a connection definition in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateConnectionCmd).Standalone()

	glue_updateConnectionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the connection resides.")
	glue_updateConnectionCmd.Flags().String("connection-input", "", "A `ConnectionInput` object that redefines the connection in question.")
	glue_updateConnectionCmd.Flags().String("name", "", "The name of the connection definition to update.")
	glue_updateConnectionCmd.MarkFlagRequired("connection-input")
	glue_updateConnectionCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_updateConnectionCmd)
}
