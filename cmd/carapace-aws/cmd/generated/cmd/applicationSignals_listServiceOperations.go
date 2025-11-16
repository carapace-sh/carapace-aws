package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceOperationsCmd = &cobra.Command{
	Use:   "list-service-operations",
	Short: "Returns a list of the *operations* of this service that have been discovered by Application Signals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listServiceOperationsCmd).Standalone()

		applicationSignals_listServiceOperationsCmd.Flags().String("end-time", "", "The end of the time period to retrieve information about.")
		applicationSignals_listServiceOperationsCmd.Flags().String("key-attributes", "", "Use this field to specify which service you want to retrieve information for.")
		applicationSignals_listServiceOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		applicationSignals_listServiceOperationsCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service operations.")
		applicationSignals_listServiceOperationsCmd.Flags().String("start-time", "", "The start of the time period to retrieve information about.")
		applicationSignals_listServiceOperationsCmd.MarkFlagRequired("end-time")
		applicationSignals_listServiceOperationsCmd.MarkFlagRequired("key-attributes")
		applicationSignals_listServiceOperationsCmd.MarkFlagRequired("start-time")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceOperationsCmd)
}
