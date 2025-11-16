package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeDocumentCmd = &cobra.Command{
	Use:   "describe-document",
	Short: "Describes the specified Amazon Web Services Systems Manager document (SSM document).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeDocumentCmd).Standalone()

	ssm_describeDocumentCmd.Flags().String("document-version", "", "The document version for which you want information.")
	ssm_describeDocumentCmd.Flags().String("name", "", "The name of the SSM document.")
	ssm_describeDocumentCmd.Flags().String("version-name", "", "An optional field specifying the version of the artifact associated with the document.")
	ssm_describeDocumentCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_describeDocumentCmd)
}
