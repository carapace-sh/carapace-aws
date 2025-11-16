package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listSipMediaApplicationsCmd = &cobra.Command{
	Use:   "list-sip-media-applications",
	Short: "Lists the SIP media applications under the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listSipMediaApplicationsCmd).Standalone()

	chimeSdkVoice_listSipMediaApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkVoice_listSipMediaApplicationsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listSipMediaApplicationsCmd)
}
