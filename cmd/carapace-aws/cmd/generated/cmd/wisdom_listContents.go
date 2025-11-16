package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listContentsCmd = &cobra.Command{
	Use:   "list-contents",
	Short: "Lists the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listContentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_listContentsCmd).Standalone()

		wisdom_listContentsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_listContentsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		wisdom_listContentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		wisdom_listContentsCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_listContentsCmd)
}
