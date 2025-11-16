package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getQuickResponseCmd = &cobra.Command{
	Use:   "get-quick-response",
	Short: "Retrieves the quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getQuickResponseCmd).Standalone()

	qconnect_getQuickResponseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_getQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response.")
	qconnect_getQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_getQuickResponseCmd.MarkFlagRequired("quick-response-id")
	qconnectCmd.AddCommand(qconnect_getQuickResponseCmd)
}
