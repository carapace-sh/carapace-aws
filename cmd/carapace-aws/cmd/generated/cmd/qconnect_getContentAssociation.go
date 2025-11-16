package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getContentAssociationCmd = &cobra.Command{
	Use:   "get-content-association",
	Short: "Returns the content association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getContentAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getContentAssociationCmd).Standalone()

		qconnect_getContentAssociationCmd.Flags().String("content-association-id", "", "The identifier of the content association.")
		qconnect_getContentAssociationCmd.Flags().String("content-id", "", "The identifier of the content.")
		qconnect_getContentAssociationCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_getContentAssociationCmd.MarkFlagRequired("content-association-id")
		qconnect_getContentAssociationCmd.MarkFlagRequired("content-id")
		qconnect_getContentAssociationCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_getContentAssociationCmd)
}
