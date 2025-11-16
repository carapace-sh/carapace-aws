package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a connection definition in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createConnectionCmd).Standalone()

		glue_createConnectionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which to create the connection.")
		glue_createConnectionCmd.Flags().String("connection-input", "", "A `ConnectionInput` object defining the connection to create.")
		glue_createConnectionCmd.Flags().String("tags", "", "The tags you assign to the connection.")
		glue_createConnectionCmd.MarkFlagRequired("connection-input")
	})
	glueCmd.AddCommand(glue_createConnectionCmd)
}
