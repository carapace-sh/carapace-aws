package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceDependenciesCmd = &cobra.Command{
	Use:   "list-service-dependencies",
	Short: "Returns a list of service dependencies of the service that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceDependenciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listServiceDependenciesCmd).Standalone()

		applicationSignals_listServiceDependenciesCmd.Flags().String("end-time", "", "The end of the time period to retrieve information about.")
		applicationSignals_listServiceDependenciesCmd.Flags().String("key-attributes", "", "Use this field to specify which service you want to retrieve information for.")
		applicationSignals_listServiceDependenciesCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		applicationSignals_listServiceDependenciesCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service dependencies.")
		applicationSignals_listServiceDependenciesCmd.Flags().String("start-time", "", "The start of the time period to retrieve information about.")
		applicationSignals_listServiceDependenciesCmd.MarkFlagRequired("end-time")
		applicationSignals_listServiceDependenciesCmd.MarkFlagRequired("key-attributes")
		applicationSignals_listServiceDependenciesCmd.MarkFlagRequired("start-time")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceDependenciesCmd)
}
