package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listOriginationNumbersCmd = &cobra.Command{
	Use:   "list-origination-numbers",
	Short: "Lists the calling Amazon Web Services account's dedicated origination numbers and their metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listOriginationNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_listOriginationNumbersCmd).Standalone()

		sns_listOriginationNumbersCmd.Flags().String("max-results", "", "The maximum number of origination numbers to return.")
		sns_listOriginationNumbersCmd.Flags().String("next-token", "", "Token that the previous `ListOriginationNumbers` request returns.")
	})
	snsCmd.AddCommand(sns_listOriginationNumbersCmd)
}
