package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipantCmd = &cobra.Command{
	Use:   "connectparticipant",
	Short: "- [Participant Service actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Participant_Service.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipantCmd).Standalone()

	})
	rootCmd.AddCommand(connectparticipantCmd)
}
