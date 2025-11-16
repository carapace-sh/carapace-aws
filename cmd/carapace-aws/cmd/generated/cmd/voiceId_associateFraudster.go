package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_associateFraudsterCmd = &cobra.Command{
	Use:   "associate-fraudster",
	Short: "Associates the fraudsters with the watchlist specified in the same domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_associateFraudsterCmd).Standalone()

	voiceId_associateFraudsterCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster.")
	voiceId_associateFraudsterCmd.Flags().String("fraudster-id", "", "The identifier of the fraudster to be associated with the watchlist.")
	voiceId_associateFraudsterCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist you want to associate with the fraudster.")
	voiceId_associateFraudsterCmd.MarkFlagRequired("domain-id")
	voiceId_associateFraudsterCmd.MarkFlagRequired("fraudster-id")
	voiceId_associateFraudsterCmd.MarkFlagRequired("watchlist-id")
	voiceIdCmd.AddCommand(voiceId_associateFraudsterCmd)
}
