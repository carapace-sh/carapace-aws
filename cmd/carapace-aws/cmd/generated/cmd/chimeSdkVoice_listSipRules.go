package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listSipRulesCmd = &cobra.Command{
	Use:   "list-sip-rules",
	Short: "Lists the SIP rules under the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listSipRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_listSipRulesCmd).Standalone()

		chimeSdkVoice_listSipRulesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkVoice_listSipRulesCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
		chimeSdkVoice_listSipRulesCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listSipRulesCmd)
}
