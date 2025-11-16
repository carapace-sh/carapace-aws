package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createCustomMetadataCmd = &cobra.Command{
	Use:   "create-custom-metadata",
	Short: "Adds one or more custom properties to the specified resource (a folder, document, or version).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createCustomMetadataCmd).Standalone()

	workdocs_createCustomMetadataCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_createCustomMetadataCmd.Flags().String("custom-metadata", "", "Custom metadata in the form of name-value pairs.")
	workdocs_createCustomMetadataCmd.Flags().String("resource-id", "", "The ID of the resource.")
	workdocs_createCustomMetadataCmd.Flags().String("version-id", "", "The ID of the version, if the custom metadata is being added to a document version.")
	workdocs_createCustomMetadataCmd.MarkFlagRequired("custom-metadata")
	workdocs_createCustomMetadataCmd.MarkFlagRequired("resource-id")
	workdocsCmd.AddCommand(workdocs_createCustomMetadataCmd)
}
