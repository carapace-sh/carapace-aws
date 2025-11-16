package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listAlertsCmd = &cobra.Command{
	Use:   "list-alerts",
	Short: "Lists the alerts that are associated with a MediaTailor channel assembly resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listAlertsCmd).Standalone()

	mediatailor_listAlertsCmd.Flags().String("max-results", "", "The maximum number of alerts that you want MediaTailor to return in response to the current request.")
	mediatailor_listAlertsCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
	mediatailor_listAlertsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	mediatailor_listAlertsCmd.MarkFlagRequired("resource-arn")
	mediatailorCmd.AddCommand(mediatailor_listAlertsCmd)
}
