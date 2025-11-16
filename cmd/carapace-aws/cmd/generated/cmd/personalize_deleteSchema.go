package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteSchemaCmd = &cobra.Command{
	Use:   "delete-schema",
	Short: "Deletes a schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_deleteSchemaCmd).Standalone()

		personalize_deleteSchemaCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the schema to delete.")
		personalize_deleteSchemaCmd.MarkFlagRequired("schema-arn")
	})
	personalizeCmd.AddCommand(personalize_deleteSchemaCmd)
}
