package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeThesaurusCmd = &cobra.Command{
	Use:   "describe-thesaurus",
	Short: "Gets information about an Amazon Kendra thesaurus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeThesaurusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_describeThesaurusCmd).Standalone()

		kendra_describeThesaurusCmd.Flags().String("id", "", "The identifier of the thesaurus you want to get information on.")
		kendra_describeThesaurusCmd.Flags().String("index-id", "", "The identifier of the index for the thesaurus.")
		kendra_describeThesaurusCmd.MarkFlagRequired("id")
		kendra_describeThesaurusCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_describeThesaurusCmd)
}
