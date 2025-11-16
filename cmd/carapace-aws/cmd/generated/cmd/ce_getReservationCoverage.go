package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getReservationCoverageCmd = &cobra.Command{
	Use:   "get-reservation-coverage",
	Short: "Retrieves the reservation coverage for your account, which you can use to see how much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon Relational Database Service, or Amazon Redshift usage is covered by a reservation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getReservationCoverageCmd).Standalone()

	ce_getReservationCoverageCmd.Flags().String("filter", "", "Filters utilization data by dimensions.")
	ce_getReservationCoverageCmd.Flags().String("granularity", "", "The granularity of the Amazon Web Services cost data for the reservation.")
	ce_getReservationCoverageCmd.Flags().String("group-by", "", "You can group the data by the following attributes:")
	ce_getReservationCoverageCmd.Flags().String("max-results", "", "The maximum number of objects that you returned for this request.")
	ce_getReservationCoverageCmd.Flags().String("metrics", "", "The measurement that you want your reservation coverage reported in.")
	ce_getReservationCoverageCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getReservationCoverageCmd.Flags().String("sort-by", "", "The value by which you want to sort the data.")
	ce_getReservationCoverageCmd.Flags().String("time-period", "", "The start and end dates of the period that you want to retrieve data about reservation coverage for.")
	ce_getReservationCoverageCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getReservationCoverageCmd)
}
