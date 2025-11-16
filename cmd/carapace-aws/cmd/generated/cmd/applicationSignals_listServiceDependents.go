package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceDependentsCmd = &cobra.Command{
	Use:   "list-service-dependents",
	Short: "Returns the list of dependents that invoked the specified service during the provided time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceDependentsCmd).Standalone()

	applicationSignals_listServiceDependentsCmd.Flags().String("end-time", "", "The end of the time period to retrieve information about.")
	applicationSignals_listServiceDependentsCmd.Flags().String("key-attributes", "", "Use this field to specify which service you want to retrieve information for.")
	applicationSignals_listServiceDependentsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	applicationSignals_listServiceDependentsCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service dependents.")
	applicationSignals_listServiceDependentsCmd.Flags().String("start-time", "", "The start of the time period to retrieve information about.")
	applicationSignals_listServiceDependentsCmd.MarkFlagRequired("end-time")
	applicationSignals_listServiceDependentsCmd.MarkFlagRequired("key-attributes")
	applicationSignals_listServiceDependentsCmd.MarkFlagRequired("start-time")
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceDependentsCmd)
}
