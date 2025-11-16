package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getCustomDataIdentifierCmd = &cobra.Command{
	Use:   "get-custom-data-identifier",
	Short: "Retrieves the criteria and other settings for a custom data identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getCustomDataIdentifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getCustomDataIdentifierCmd).Standalone()

		macie2_getCustomDataIdentifierCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_getCustomDataIdentifierCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_getCustomDataIdentifierCmd)
}
