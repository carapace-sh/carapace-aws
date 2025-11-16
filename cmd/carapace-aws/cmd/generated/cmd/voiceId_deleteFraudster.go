package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_deleteFraudsterCmd = &cobra.Command{
	Use:   "delete-fraudster",
	Short: "Deletes the specified fraudster from Voice ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_deleteFraudsterCmd).Standalone()

	voiceId_deleteFraudsterCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster.")
	voiceId_deleteFraudsterCmd.Flags().String("fraudster-id", "", "The identifier of the fraudster you want to delete.")
	voiceId_deleteFraudsterCmd.MarkFlagRequired("domain-id")
	voiceId_deleteFraudsterCmd.MarkFlagRequired("fraudster-id")
	voiceIdCmd.AddCommand(voiceId_deleteFraudsterCmd)
}
