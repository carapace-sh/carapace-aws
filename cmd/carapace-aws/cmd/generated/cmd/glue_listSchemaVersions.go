package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listSchemaVersionsCmd = &cobra.Command{
	Use:   "list-schema-versions",
	Short: "Returns a list of schema versions that you have created, with minimal information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listSchemaVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listSchemaVersionsCmd).Standalone()

		glue_listSchemaVersionsCmd.Flags().String("max-results", "", "Maximum number of results required per page.")
		glue_listSchemaVersionsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_listSchemaVersionsCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
		glue_listSchemaVersionsCmd.MarkFlagRequired("schema-id")
	})
	glueCmd.AddCommand(glue_listSchemaVersionsCmd)
}
