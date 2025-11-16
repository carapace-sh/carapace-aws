package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_exportSchemaCmd = &cobra.Command{
	Use:   "export-schema",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_exportSchemaCmd).Standalone()

	schemas_exportSchemaCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_exportSchemaCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_exportSchemaCmd.Flags().String("schema-version", "", "Specifying this limits the results to only this schema version.")
	schemas_exportSchemaCmd.Flags().String("type", "", "")
	schemas_exportSchemaCmd.MarkFlagRequired("registry-name")
	schemas_exportSchemaCmd.MarkFlagRequired("schema-name")
	schemas_exportSchemaCmd.MarkFlagRequired("type")
	schemasCmd.AddCommand(schemas_exportSchemaCmd)
}
