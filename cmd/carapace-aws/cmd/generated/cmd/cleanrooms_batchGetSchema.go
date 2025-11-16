package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_batchGetSchemaCmd = &cobra.Command{
	Use:   "batch-get-schema",
	Short: "Retrieves multiple schemas by their identifiers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_batchGetSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_batchGetSchemaCmd).Standalone()

		cleanrooms_batchGetSchemaCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the schemas belong to.")
		cleanrooms_batchGetSchemaCmd.Flags().String("names", "", "The names for the schema objects to retrieve.")
		cleanrooms_batchGetSchemaCmd.MarkFlagRequired("collaboration-identifier")
		cleanrooms_batchGetSchemaCmd.MarkFlagRequired("names")
	})
	cleanroomsCmd.AddCommand(cleanrooms_batchGetSchemaCmd)
}
