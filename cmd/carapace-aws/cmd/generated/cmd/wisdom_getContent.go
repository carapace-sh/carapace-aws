package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getContentCmd = &cobra.Command{
	Use:   "get-content",
	Short: "Retrieves content, including a pre-signed URL to download the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_getContentCmd).Standalone()

		wisdom_getContentCmd.Flags().String("content-id", "", "The identifier of the content.")
		wisdom_getContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_getContentCmd.MarkFlagRequired("content-id")
		wisdom_getContentCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_getContentCmd)
}
