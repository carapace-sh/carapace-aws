package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getSipRuleCmd = &cobra.Command{
	Use:   "get-sip-rule",
	Short: "Retrieves the details of a SIP rule, such as the rule ID, name, triggers, and target endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getSipRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getSipRuleCmd).Standalone()

		chimeSdkVoice_getSipRuleCmd.Flags().String("sip-rule-id", "", "The SIP rule ID.")
		chimeSdkVoice_getSipRuleCmd.MarkFlagRequired("sip-rule-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getSipRuleCmd)
}
