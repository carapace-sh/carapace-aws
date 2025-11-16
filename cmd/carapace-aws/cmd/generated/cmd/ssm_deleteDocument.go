package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteDocumentCmd = &cobra.Command{
	Use:   "delete-document",
	Short: "Deletes the Amazon Web Services Systems Manager document (SSM document) and all managed node associations to the document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteDocumentCmd).Standalone()

	ssm_deleteDocumentCmd.Flags().String("document-version", "", "The version of the document that you want to delete.")
	ssm_deleteDocumentCmd.Flags().Bool("force", false, "Some SSM document types require that you specify a `Force` flag before you can delete the document.")
	ssm_deleteDocumentCmd.Flags().String("name", "", "The name of the document.")
	ssm_deleteDocumentCmd.Flags().Bool("no-force", false, "Some SSM document types require that you specify a `Force` flag before you can delete the document.")
	ssm_deleteDocumentCmd.Flags().String("version-name", "", "The version name of the document that you want to delete.")
	ssm_deleteDocumentCmd.MarkFlagRequired("name")
	ssm_deleteDocumentCmd.Flag("no-force").Hidden = true
	ssmCmd.AddCommand(ssm_deleteDocumentCmd)
}
