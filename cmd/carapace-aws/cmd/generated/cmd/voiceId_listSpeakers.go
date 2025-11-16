package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listSpeakersCmd = &cobra.Command{
	Use:   "list-speakers",
	Short: "Lists all speakers in a specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listSpeakersCmd).Standalone()

	voiceId_listSpeakersCmd.Flags().String("domain-id", "", "The identifier of the domain.")
	voiceId_listSpeakersCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	voiceId_listSpeakersCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
	voiceId_listSpeakersCmd.MarkFlagRequired("domain-id")
	voiceIdCmd.AddCommand(voiceId_listSpeakersCmd)
}
