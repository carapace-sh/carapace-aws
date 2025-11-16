package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getCalendarStateCmd = &cobra.Command{
	Use:   "get-calendar-state",
	Short: "Gets the state of a Amazon Web Services Systems Manager change calendar at the current time or a specified time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getCalendarStateCmd).Standalone()

	ssm_getCalendarStateCmd.Flags().String("at-time", "", "(Optional) The specific time for which you want to get calendar state information, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.")
	ssm_getCalendarStateCmd.Flags().String("calendar-names", "", "The names of Amazon Resource Names (ARNs) of the Systems Manager documents (SSM documents) that represent the calendar entries for which you want to get the state.")
	ssm_getCalendarStateCmd.MarkFlagRequired("calendar-names")
	ssmCmd.AddCommand(ssm_getCalendarStateCmd)
}
