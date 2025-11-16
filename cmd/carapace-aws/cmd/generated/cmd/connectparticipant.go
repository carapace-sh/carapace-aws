package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipantCmd = &cobra.Command{
	Use:   "connectparticipant",
	Short: "- [Participant Service actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Participant_Service.html)\n- [Participant Service data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Connect_Participant_Service.html)\n\nAmazon Connect is an easy-to-use omnichannel cloud contact center service that enables companies of any size to deliver superior customer service at a lower cost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipantCmd).Standalone()

	})
	rootCmd.AddCommand(connectparticipantCmd)
}
