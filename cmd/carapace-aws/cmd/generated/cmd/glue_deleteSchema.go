package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteSchemaCmd = &cobra.Command{
	Use:   "delete-schema",
	Short: "Deletes the entire schema set, including the schema set and all of its versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteSchemaCmd).Standalone()

		glue_deleteSchemaCmd.Flags().String("schema-id", "", "This is a wrapper structure that may contain the schema name and Amazon Resource Name (ARN).")
		glue_deleteSchemaCmd.MarkFlagRequired("schema-id")
	})
	glueCmd.AddCommand(glue_deleteSchemaCmd)
}
