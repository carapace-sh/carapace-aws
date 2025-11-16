package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_createParticipantTokenCmd = &cobra.Command{
	Use:   "create-participant-token",
	Short: "Creates an additional token for a specified stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_createParticipantTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_createParticipantTokenCmd).Standalone()

		ivsRealtime_createParticipantTokenCmd.Flags().String("attributes", "", "Application-provided attributes to encode into the token and attach to a stage.")
		ivsRealtime_createParticipantTokenCmd.Flags().String("capabilities", "", "Set of capabilities that the user is allowed to perform in the stage.")
		ivsRealtime_createParticipantTokenCmd.Flags().String("duration", "", "Duration (in minutes), after which the token expires.")
		ivsRealtime_createParticipantTokenCmd.Flags().String("stage-arn", "", "ARN of the stage to which this token is scoped.")
		ivsRealtime_createParticipantTokenCmd.Flags().String("user-id", "", "Name that can be specified to help identify the token.")
		ivsRealtime_createParticipantTokenCmd.MarkFlagRequired("stage-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_createParticipantTokenCmd)
}
