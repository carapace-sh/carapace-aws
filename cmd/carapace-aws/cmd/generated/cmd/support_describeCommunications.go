package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeCommunicationsCmd = &cobra.Command{
	Use:   "describe-communications",
	Short: "Returns communications and attachments for one or more support cases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeCommunicationsCmd).Standalone()

	support_describeCommunicationsCmd.Flags().String("after-time", "", "The start date for a filtered date search on support case communications.")
	support_describeCommunicationsCmd.Flags().String("before-time", "", "The end date for a filtered date search on support case communications.")
	support_describeCommunicationsCmd.Flags().String("case-id", "", "The support case ID requested or returned in the call.")
	support_describeCommunicationsCmd.Flags().String("max-results", "", "The maximum number of results to return before paginating.")
	support_describeCommunicationsCmd.Flags().String("next-token", "", "A resumption point for pagination.")
	support_describeCommunicationsCmd.MarkFlagRequired("case-id")
	supportCmd.AddCommand(support_describeCommunicationsCmd)
}
