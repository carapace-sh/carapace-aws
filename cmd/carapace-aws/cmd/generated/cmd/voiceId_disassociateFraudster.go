package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_disassociateFraudsterCmd = &cobra.Command{
	Use:   "disassociate-fraudster",
	Short: "Disassociates the fraudsters from the watchlist specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_disassociateFraudsterCmd).Standalone()

	voiceId_disassociateFraudsterCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster.")
	voiceId_disassociateFraudsterCmd.Flags().String("fraudster-id", "", "The identifier of the fraudster to be disassociated from the watchlist.")
	voiceId_disassociateFraudsterCmd.Flags().String("watchlist-id", "", "The identifier of the watchlist that you want to disassociate from the fraudster.")
	voiceId_disassociateFraudsterCmd.MarkFlagRequired("domain-id")
	voiceId_disassociateFraudsterCmd.MarkFlagRequired("fraudster-id")
	voiceId_disassociateFraudsterCmd.MarkFlagRequired("watchlist-id")
	voiceIdCmd.AddCommand(voiceId_disassociateFraudsterCmd)
}
