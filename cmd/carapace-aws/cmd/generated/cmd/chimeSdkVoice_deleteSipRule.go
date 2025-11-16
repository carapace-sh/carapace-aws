package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteSipRuleCmd = &cobra.Command{
	Use:   "delete-sip-rule",
	Short: "Deletes a SIP rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteSipRuleCmd).Standalone()

	chimeSdkVoice_deleteSipRuleCmd.Flags().String("sip-rule-id", "", "The SIP rule ID.")
	chimeSdkVoice_deleteSipRuleCmd.MarkFlagRequired("sip-rule-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteSipRuleCmd)
}
