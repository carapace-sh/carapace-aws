package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchDeleteConnectionCmd = &cobra.Command{
	Use:   "batch-delete-connection",
	Short: "Deletes a list of connection definitions from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchDeleteConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchDeleteConnectionCmd).Standalone()

		glue_batchDeleteConnectionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the connections reside.")
		glue_batchDeleteConnectionCmd.Flags().String("connection-name-list", "", "A list of names of the connections to delete.")
		glue_batchDeleteConnectionCmd.MarkFlagRequired("connection-name-list")
	})
	glueCmd.AddCommand(glue_batchDeleteConnectionCmd)
}
