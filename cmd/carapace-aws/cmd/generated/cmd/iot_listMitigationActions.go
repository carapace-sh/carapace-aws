package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listMitigationActionsCmd = &cobra.Command{
	Use:   "list-mitigation-actions",
	Short: "Gets a list of all mitigation actions that match the specified filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listMitigationActionsCmd).Standalone()

	iot_listMitigationActionsCmd.Flags().String("action-type", "", "Specify a value to limit the result to mitigation actions with a specific action type.")
	iot_listMitigationActionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listMitigationActionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iotCmd.AddCommand(iot_listMitigationActionsCmd)
}
