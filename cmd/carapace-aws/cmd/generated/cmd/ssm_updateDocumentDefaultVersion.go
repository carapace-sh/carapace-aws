package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateDocumentDefaultVersionCmd = &cobra.Command{
	Use:   "update-document-default-version",
	Short: "Set the default version of a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateDocumentDefaultVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateDocumentDefaultVersionCmd).Standalone()

		ssm_updateDocumentDefaultVersionCmd.Flags().String("document-version", "", "The version of a custom document that you want to set as the default version.")
		ssm_updateDocumentDefaultVersionCmd.Flags().String("name", "", "The name of a custom document that you want to set as the default version.")
		ssm_updateDocumentDefaultVersionCmd.MarkFlagRequired("document-version")
		ssm_updateDocumentDefaultVersionCmd.MarkFlagRequired("name")
	})
	ssmCmd.AddCommand(ssm_updateDocumentDefaultVersionCmd)
}
