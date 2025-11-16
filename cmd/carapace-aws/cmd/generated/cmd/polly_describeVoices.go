package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var polly_describeVoicesCmd = &cobra.Command{
	Use:   "describe-voices",
	Short: "Returns the list of voices that are available for use when requesting speech synthesis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(polly_describeVoicesCmd).Standalone()

	polly_describeVoicesCmd.Flags().String("engine", "", "Specifies the engine (`standard`, `neural`, `long-form` or `generative`) used by Amazon Polly when processing input text for speech synthesis.")
	polly_describeVoicesCmd.Flags().String("include-additional-language-codes", "", "Boolean value indicating whether to return any bilingual voices that use the specified language as an additional language.")
	polly_describeVoicesCmd.Flags().String("language-code", "", "The language identification tag (ISO 639 code for the language name-ISO 3166 country code) for filtering the list of voices returned.")
	polly_describeVoicesCmd.Flags().String("next-token", "", "An opaque pagination token returned from the previous `DescribeVoices` operation.")
	pollyCmd.AddCommand(polly_describeVoicesCmd)
}
