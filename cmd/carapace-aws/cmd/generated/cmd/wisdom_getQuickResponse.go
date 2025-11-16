package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getQuickResponseCmd = &cobra.Command{
	Use:   "get-quick-response",
	Short: "Retrieves the quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getQuickResponseCmd).Standalone()

	wisdom_getQuickResponseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_getQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response.")
	wisdom_getQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_getQuickResponseCmd.MarkFlagRequired("quick-response-id")
	wisdomCmd.AddCommand(wisdom_getQuickResponseCmd)
}
