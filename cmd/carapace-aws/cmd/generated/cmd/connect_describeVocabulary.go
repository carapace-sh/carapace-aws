package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeVocabularyCmd = &cobra.Command{
	Use:   "describe-vocabulary",
	Short: "Describes the specified vocabulary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeVocabularyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeVocabularyCmd).Standalone()

		connect_describeVocabularyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeVocabularyCmd.Flags().String("vocabulary-id", "", "The identifier of the custom vocabulary.")
		connect_describeVocabularyCmd.MarkFlagRequired("instance-id")
		connect_describeVocabularyCmd.MarkFlagRequired("vocabulary-id")
	})
	connectCmd.AddCommand(connect_describeVocabularyCmd)
}
