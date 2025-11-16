package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createSipRuleCmd = &cobra.Command{
	Use:   "create-sip-rule",
	Short: "Creates a SIP rule, which can be used to run a SIP media application as a target for a specific trigger type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createSipRuleCmd).Standalone()

	chimeSdkVoice_createSipRuleCmd.Flags().String("disabled", "", "Disables or enables a SIP rule.")
	chimeSdkVoice_createSipRuleCmd.Flags().String("name", "", "The name of the SIP rule.")
	chimeSdkVoice_createSipRuleCmd.Flags().String("target-applications", "", "List of SIP media applications, with priority and AWS Region.")
	chimeSdkVoice_createSipRuleCmd.Flags().String("trigger-type", "", "The type of trigger assigned to the SIP rule in `TriggerValue`, currently `RequestUriHostname` or `ToPhoneNumber`.")
	chimeSdkVoice_createSipRuleCmd.Flags().String("trigger-value", "", "If `TriggerType` is `RequestUriHostname`, the value can be the outbound host name of a Voice Connector.")
	chimeSdkVoice_createSipRuleCmd.MarkFlagRequired("name")
	chimeSdkVoice_createSipRuleCmd.MarkFlagRequired("trigger-type")
	chimeSdkVoice_createSipRuleCmd.MarkFlagRequired("trigger-value")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createSipRuleCmd)
}
