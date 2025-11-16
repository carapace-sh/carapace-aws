package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createDocumentCmd = &cobra.Command{
	Use:   "create-document",
	Short: "Creates a Amazon Web Services Systems Manager (SSM document).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createDocumentCmd).Standalone()

	ssm_createDocumentCmd.Flags().String("attachments", "", "A list of key-value pairs that describe attachments to a version of a document.")
	ssm_createDocumentCmd.Flags().String("content", "", "The content for the new SSM document in JSON or YAML format.")
	ssm_createDocumentCmd.Flags().String("display-name", "", "An optional field where you can specify a friendly name for the SSM document.")
	ssm_createDocumentCmd.Flags().String("document-format", "", "Specify the document format for the request.")
	ssm_createDocumentCmd.Flags().String("document-type", "", "The type of document to create.")
	ssm_createDocumentCmd.Flags().String("name", "", "A name for the SSM document.")
	ssm_createDocumentCmd.Flags().String("requires", "", "A list of SSM documents required by a document.")
	ssm_createDocumentCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
	ssm_createDocumentCmd.Flags().String("target-type", "", "Specify a target type to define the kinds of resources the document can run on.")
	ssm_createDocumentCmd.Flags().String("version-name", "", "An optional field specifying the version of the artifact you are creating with the document.")
	ssm_createDocumentCmd.MarkFlagRequired("content")
	ssm_createDocumentCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_createDocumentCmd)
}
