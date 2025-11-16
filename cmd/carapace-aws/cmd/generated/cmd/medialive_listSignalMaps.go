package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listSignalMapsCmd = &cobra.Command{
	Use:   "list-signal-maps",
	Short: "Lists signal maps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listSignalMapsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listSignalMapsCmd).Standalone()

		medialive_listSignalMapsCmd.Flags().String("cloud-watch-alarm-template-group-identifier", "", "A cloudwatch alarm template group's identifier.")
		medialive_listSignalMapsCmd.Flags().String("event-bridge-rule-template-group-identifier", "", "An eventbridge rule template group's identifier.")
		medialive_listSignalMapsCmd.Flags().String("max-results", "", "")
		medialive_listSignalMapsCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results in paginated list responses.")
	})
	medialiveCmd.AddCommand(medialive_listSignalMapsCmd)
}
