package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listSchemaExtensionsCmd = &cobra.Command{
	Use:   "list-schema-extensions",
	Short: "Lists all schema extensions applied to a Microsoft AD Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listSchemaExtensionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_listSchemaExtensionsCmd).Standalone()

		ds_listSchemaExtensionsCmd.Flags().String("directory-id", "", "The identifier of the directory from which to retrieve the schema extension information.")
		ds_listSchemaExtensionsCmd.Flags().String("limit", "", "The maximum number of items to return.")
		ds_listSchemaExtensionsCmd.Flags().String("next-token", "", "The `ListSchemaExtensions.NextToken` value from a previous call to `ListSchemaExtensions`.")
		ds_listSchemaExtensionsCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_listSchemaExtensionsCmd)
}
