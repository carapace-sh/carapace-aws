package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getSchemaCmd = &cobra.Command{
	Use:   "get-schema",
	Short: "Retrieves the schema for a relation within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getSchemaCmd).Standalone()

	cleanrooms_getSchemaCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the schema belongs to.")
	cleanrooms_getSchemaCmd.Flags().String("name", "", "The name of the relation to retrieve the schema for.")
	cleanrooms_getSchemaCmd.MarkFlagRequired("collaboration-identifier")
	cleanrooms_getSchemaCmd.MarkFlagRequired("name")
	cleanroomsCmd.AddCommand(cleanrooms_getSchemaCmd)
}
