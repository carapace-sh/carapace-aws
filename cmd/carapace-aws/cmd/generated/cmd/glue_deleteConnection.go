package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes a connection from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteConnectionCmd).Standalone()

	glue_deleteConnectionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the connection resides.")
	glue_deleteConnectionCmd.Flags().String("connection-name", "", "The name of the connection to delete.")
	glue_deleteConnectionCmd.MarkFlagRequired("connection-name")
	glueCmd.AddCommand(glue_deleteConnectionCmd)
}
