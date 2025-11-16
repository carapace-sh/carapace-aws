package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeOptOutListsCmd = &cobra.Command{
	Use:   "describe-opt-out-lists",
	Short: "Describes the specified opt-out list or all opt-out lists in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeOptOutListsCmd).Standalone()

	pinpointSmsVoiceV2_describeOptOutListsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeOptOutListsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeOptOutListsCmd.Flags().String("opt-out-list-names", "", "The OptOutLists to show the details of.")
	pinpointSmsVoiceV2_describeOptOutListsCmd.Flags().String("owner", "", "Use `SELF` to filter the list of Opt-Out List to ones your account owns or use `SHARED` to filter on Opt-Out List shared with your account.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeOptOutListsCmd)
}
