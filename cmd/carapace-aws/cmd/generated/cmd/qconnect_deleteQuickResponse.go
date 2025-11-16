package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteQuickResponseCmd = &cobra.Command{
	Use:   "delete-quick-response",
	Short: "Deletes a quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteQuickResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteQuickResponseCmd).Standalone()

		qconnect_deleteQuickResponseCmd.Flags().String("knowledge-base-id", "", "The knowledge base from which the quick response is deleted.")
		qconnect_deleteQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response to delete.")
		qconnect_deleteQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_deleteQuickResponseCmd.MarkFlagRequired("quick-response-id")
	})
	qconnectCmd.AddCommand(qconnect_deleteQuickResponseCmd)
}
