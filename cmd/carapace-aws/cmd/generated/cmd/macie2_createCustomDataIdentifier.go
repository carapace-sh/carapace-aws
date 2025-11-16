package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createCustomDataIdentifierCmd = &cobra.Command{
	Use:   "create-custom-data-identifier",
	Short: "Creates and defines the criteria and other settings for a custom data identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createCustomDataIdentifierCmd).Standalone()

	macie2_createCustomDataIdentifierCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
	macie2_createCustomDataIdentifierCmd.Flags().String("description", "", "A custom description of the custom data identifier.")
	macie2_createCustomDataIdentifierCmd.Flags().String("ignore-words", "", "An array that lists specific character sequences (*ignore words*) to exclude from the results.")
	macie2_createCustomDataIdentifierCmd.Flags().String("keywords", "", "An array that lists specific character sequences (*keywords*), one of which must precede and be within proximity (maximumMatchDistance) of the regular expression to match.")
	macie2_createCustomDataIdentifierCmd.Flags().String("maximum-match-distance", "", "The maximum number of characters that can exist between the end of at least one complete character sequence specified by the keywords array and the end of the text that matches the regex pattern.")
	macie2_createCustomDataIdentifierCmd.Flags().String("name", "", "A custom name for the custom data identifier.")
	macie2_createCustomDataIdentifierCmd.Flags().String("regex", "", "The regular expression (*regex*) that defines the pattern to match.")
	macie2_createCustomDataIdentifierCmd.Flags().String("severity-levels", "", "The severity to assign to findings that the custom data identifier produces, based on the number of occurrences of text that match the custom data identifier's detection criteria.")
	macie2_createCustomDataIdentifierCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the custom data identifier.")
	macie2_createCustomDataIdentifierCmd.MarkFlagRequired("name")
	macie2_createCustomDataIdentifierCmd.MarkFlagRequired("regex")
	macie2Cmd.AddCommand(macie2_createCustomDataIdentifierCmd)
}
