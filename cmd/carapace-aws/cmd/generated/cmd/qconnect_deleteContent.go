package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteContentCmd = &cobra.Command{
	Use:   "delete-content",
	Short: "Deletes the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteContentCmd).Standalone()

	qconnect_deleteContentCmd.Flags().String("content-id", "", "The identifier of the content.")
	qconnect_deleteContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_deleteContentCmd.MarkFlagRequired("content-id")
	qconnect_deleteContentCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_deleteContentCmd)
}
