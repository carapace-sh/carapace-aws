package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_cancelSchemaExtensionCmd = &cobra.Command{
	Use:   "cancel-schema-extension",
	Short: "Cancels an in-progress schema extension to a Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_cancelSchemaExtensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_cancelSchemaExtensionCmd).Standalone()

		ds_cancelSchemaExtensionCmd.Flags().String("directory-id", "", "The identifier of the directory whose schema extension will be canceled.")
		ds_cancelSchemaExtensionCmd.Flags().String("schema-extension-id", "", "The identifier of the schema extension that will be canceled.")
		ds_cancelSchemaExtensionCmd.MarkFlagRequired("directory-id")
		ds_cancelSchemaExtensionCmd.MarkFlagRequired("schema-extension-id")
	})
	dsCmd.AddCommand(ds_cancelSchemaExtensionCmd)
}
