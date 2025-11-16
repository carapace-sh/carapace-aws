package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteCustomMetadataCmd = &cobra.Command{
	Use:   "delete-custom-metadata",
	Short: "Deletes custom metadata from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteCustomMetadataCmd).Standalone()

	workdocs_deleteCustomMetadataCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deleteCustomMetadataCmd.Flags().String("delete-all", "", "Flag to indicate removal of all custom metadata properties from the specified resource.")
	workdocs_deleteCustomMetadataCmd.Flags().String("keys", "", "List of properties to remove.")
	workdocs_deleteCustomMetadataCmd.Flags().String("resource-id", "", "The ID of the resource, either a document or folder.")
	workdocs_deleteCustomMetadataCmd.Flags().String("version-id", "", "The ID of the version, if the custom metadata is being deleted from a document version.")
	workdocs_deleteCustomMetadataCmd.MarkFlagRequired("resource-id")
	workdocsCmd.AddCommand(workdocs_deleteCustomMetadataCmd)
}
