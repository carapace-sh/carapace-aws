package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getReservationUtilizationCmd = &cobra.Command{
	Use:   "get-reservation-utilization",
	Short: "Retrieves the reservation utilization for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getReservationUtilizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getReservationUtilizationCmd).Standalone()

		ce_getReservationUtilizationCmd.Flags().String("filter", "", "Filters utilization data by dimensions.")
		ce_getReservationUtilizationCmd.Flags().String("granularity", "", "If `GroupBy` is set, `Granularity` can't be set.")
		ce_getReservationUtilizationCmd.Flags().String("group-by", "", "Groups only by `SUBSCRIPTION_ID`.")
		ce_getReservationUtilizationCmd.Flags().String("max-results", "", "The maximum number of objects that you returned for this request.")
		ce_getReservationUtilizationCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
		ce_getReservationUtilizationCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
		ce_getReservationUtilizationCmd.Flags().String("time-period", "", "Sets the start and end dates for retrieving Reserved Instance (RI) utilization.")
		ce_getReservationUtilizationCmd.MarkFlagRequired("time-period")
	})
	ceCmd.AddCommand(ce_getReservationUtilizationCmd)
}
