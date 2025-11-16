package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateDocumentCmd = &cobra.Command{
	Use:   "update-document",
	Short: "Updates one or more values for an SSM document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateDocumentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateDocumentCmd).Standalone()

		ssm_updateDocumentCmd.Flags().String("attachments", "", "A list of key-value pairs that describe attachments to a version of a document.")
		ssm_updateDocumentCmd.Flags().String("content", "", "A valid JSON or YAML string.")
		ssm_updateDocumentCmd.Flags().String("display-name", "", "The friendly name of the SSM document that you want to update.")
		ssm_updateDocumentCmd.Flags().String("document-format", "", "Specify the document format for the new document version.")
		ssm_updateDocumentCmd.Flags().String("document-version", "", "The version of the document that you want to update.")
		ssm_updateDocumentCmd.Flags().String("name", "", "The name of the SSM document that you want to update.")
		ssm_updateDocumentCmd.Flags().String("target-type", "", "Specify a new target type for the document.")
		ssm_updateDocumentCmd.Flags().String("version-name", "", "An optional field specifying the version of the artifact you are updating with the document.")
		ssm_updateDocumentCmd.MarkFlagRequired("content")
		ssm_updateDocumentCmd.MarkFlagRequired("name")
	})
	ssmCmd.AddCommand(ssm_updateDocumentCmd)
}
