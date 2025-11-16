package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getDocumentCmd = &cobra.Command{
	Use:   "get-document",
	Short: "Gets the contents of the specified Amazon Web Services Systems Manager document (SSM document).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getDocumentCmd).Standalone()

	ssm_getDocumentCmd.Flags().String("document-format", "", "Returns the document in the specified format.")
	ssm_getDocumentCmd.Flags().String("document-version", "", "The document version for which you want information.")
	ssm_getDocumentCmd.Flags().String("name", "", "The name of the SSM document.")
	ssm_getDocumentCmd.Flags().String("version-name", "", "An optional field specifying the version of the artifact associated with the document.")
	ssm_getDocumentCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_getDocumentCmd)
}
