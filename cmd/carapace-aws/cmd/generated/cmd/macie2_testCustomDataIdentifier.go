package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_testCustomDataIdentifierCmd = &cobra.Command{
	Use:   "test-custom-data-identifier",
	Short: "Tests criteria for a custom data identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_testCustomDataIdentifierCmd).Standalone()

	macie2_testCustomDataIdentifierCmd.Flags().String("ignore-words", "", "An array that lists specific character sequences (*ignore words*) to exclude from the results.")
	macie2_testCustomDataIdentifierCmd.Flags().String("keywords", "", "An array that lists specific character sequences (*keywords*), one of which must precede and be within proximity (maximumMatchDistance) of the regular expression to match.")
	macie2_testCustomDataIdentifierCmd.Flags().String("maximum-match-distance", "", "The maximum number of characters that can exist between the end of at least one complete character sequence specified by the keywords array and the end of the text that matches the regex pattern.")
	macie2_testCustomDataIdentifierCmd.Flags().String("regex", "", "The regular expression (*regex*) that defines the pattern to match.")
	macie2_testCustomDataIdentifierCmd.Flags().String("sample-text", "", "The sample text to inspect by using the custom data identifier.")
	macie2_testCustomDataIdentifierCmd.MarkFlagRequired("regex")
	macie2_testCustomDataIdentifierCmd.MarkFlagRequired("sample-text")
	macie2Cmd.AddCommand(macie2_testCustomDataIdentifierCmd)
}
