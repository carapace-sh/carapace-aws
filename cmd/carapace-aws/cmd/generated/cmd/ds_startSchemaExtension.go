package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_startSchemaExtensionCmd = &cobra.Command{
	Use:   "start-schema-extension",
	Short: "Applies a schema extension to a Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_startSchemaExtensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_startSchemaExtensionCmd).Standalone()

		ds_startSchemaExtensionCmd.Flags().String("create-snapshot-before-schema-extension", "", "If true, creates a snapshot of the directory before applying the schema extension.")
		ds_startSchemaExtensionCmd.Flags().String("description", "", "A description of the schema extension.")
		ds_startSchemaExtensionCmd.Flags().String("directory-id", "", "The identifier of the directory for which the schema extension will be applied to.")
		ds_startSchemaExtensionCmd.Flags().String("ldif-content", "", "The LDIF file represented as a string.")
		ds_startSchemaExtensionCmd.MarkFlagRequired("create-snapshot-before-schema-extension")
		ds_startSchemaExtensionCmd.MarkFlagRequired("description")
		ds_startSchemaExtensionCmd.MarkFlagRequired("directory-id")
		ds_startSchemaExtensionCmd.MarkFlagRequired("ldif-content")
	})
	dsCmd.AddCommand(ds_startSchemaExtensionCmd)
}
