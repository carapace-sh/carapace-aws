package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_getServiceCmd = &cobra.Command{
	Use:   "get-service",
	Short: "Returns information about a service discovered by Application Signals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_getServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_getServiceCmd).Standalone()

		applicationSignals_getServiceCmd.Flags().String("end-time", "", "The end of the time period to retrieve information about.")
		applicationSignals_getServiceCmd.Flags().String("key-attributes", "", "Use this field to specify which service you want to retrieve information for.")
		applicationSignals_getServiceCmd.Flags().String("start-time", "", "The start of the time period to retrieve information about.")
		applicationSignals_getServiceCmd.MarkFlagRequired("end-time")
		applicationSignals_getServiceCmd.MarkFlagRequired("key-attributes")
		applicationSignals_getServiceCmd.MarkFlagRequired("start-time")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_getServiceCmd)
}
