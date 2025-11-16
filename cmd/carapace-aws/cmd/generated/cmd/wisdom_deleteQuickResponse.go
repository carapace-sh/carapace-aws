package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteQuickResponseCmd = &cobra.Command{
	Use:   "delete-quick-response",
	Short: "Deletes a quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteQuickResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_deleteQuickResponseCmd).Standalone()

		wisdom_deleteQuickResponseCmd.Flags().String("knowledge-base-id", "", "The knowledge base from which the quick response is deleted.")
		wisdom_deleteQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response to delete.")
		wisdom_deleteQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
		wisdom_deleteQuickResponseCmd.MarkFlagRequired("quick-response-id")
	})
	wisdomCmd.AddCommand(wisdom_deleteQuickResponseCmd)
}
