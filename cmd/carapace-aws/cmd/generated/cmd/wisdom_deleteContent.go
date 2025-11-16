package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteContentCmd = &cobra.Command{
	Use:   "delete-content",
	Short: "Deletes the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteContentCmd).Standalone()

	wisdom_deleteContentCmd.Flags().String("content-id", "", "The identifier of the content.")
	wisdom_deleteContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_deleteContentCmd.MarkFlagRequired("content-id")
	wisdom_deleteContentCmd.MarkFlagRequired("knowledge-base-id")
	wisdomCmd.AddCommand(wisdom_deleteContentCmd)
}
