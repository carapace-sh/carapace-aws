package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateSipRuleCmd = &cobra.Command{
	Use:   "update-sip-rule",
	Short: "Updates the details of the specified SIP rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateSipRuleCmd).Standalone()

	chimeSdkVoice_updateSipRuleCmd.Flags().String("disabled", "", "The new value that indicates whether the rule is disabled.")
	chimeSdkVoice_updateSipRuleCmd.Flags().String("name", "", "The new name for the specified SIP rule.")
	chimeSdkVoice_updateSipRuleCmd.Flags().String("sip-rule-id", "", "The SIP rule ID.")
	chimeSdkVoice_updateSipRuleCmd.Flags().String("target-applications", "", "The new list of target applications.")
	chimeSdkVoice_updateSipRuleCmd.MarkFlagRequired("name")
	chimeSdkVoice_updateSipRuleCmd.MarkFlagRequired("sip-rule-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateSipRuleCmd)
}
