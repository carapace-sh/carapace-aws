package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteContentAssociationCmd = &cobra.Command{
	Use:   "delete-content-association",
	Short: "Deletes the content association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteContentAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteContentAssociationCmd).Standalone()

		qconnect_deleteContentAssociationCmd.Flags().String("content-association-id", "", "The identifier of the content association.")
		qconnect_deleteContentAssociationCmd.Flags().String("content-id", "", "The identifier of the content.")
		qconnect_deleteContentAssociationCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_deleteContentAssociationCmd.MarkFlagRequired("content-association-id")
		qconnect_deleteContentAssociationCmd.MarkFlagRequired("content-id")
		qconnect_deleteContentAssociationCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_deleteContentAssociationCmd)
}
