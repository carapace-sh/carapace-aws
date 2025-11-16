package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listFaqsCmd = &cobra.Command{
	Use:   "list-faqs",
	Short: "Gets a list of FAQs associated with an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listFaqsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_listFaqsCmd).Standalone()

		kendra_listFaqsCmd.Flags().String("index-id", "", "The index for the FAQs.")
		kendra_listFaqsCmd.Flags().String("max-results", "", "The maximum number of FAQs to return in the response.")
		kendra_listFaqsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
		kendra_listFaqsCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_listFaqsCmd)
}
