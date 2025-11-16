package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceStatesCmd = &cobra.Command{
	Use:   "list-service-states",
	Short: "Retrieves the current state information for services monitored by Application Signals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceStatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listServiceStatesCmd).Standalone()

		applicationSignals_listServiceStatesCmd.Flags().String("attribute-filters", "", "An array of attribute filters to narrow down the service states returned.")
		applicationSignals_listServiceStatesCmd.Flags().String("aws-account-id", "", "The AWS account ID to filter service states.")
		applicationSignals_listServiceStatesCmd.Flags().String("end-time", "", "The end time for the service states query.")
		applicationSignals_listServiceStatesCmd.Flags().Bool("include-linked-accounts", false, "Specifies whether to include service states from linked AWS accounts in the results.")
		applicationSignals_listServiceStatesCmd.Flags().String("max-results", "", "The maximum number of service states to return in a single request.")
		applicationSignals_listServiceStatesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		applicationSignals_listServiceStatesCmd.Flags().Bool("no-include-linked-accounts", false, "Specifies whether to include service states from linked AWS accounts in the results.")
		applicationSignals_listServiceStatesCmd.Flags().String("start-time", "", "The start time for the service states query.")
		applicationSignals_listServiceStatesCmd.MarkFlagRequired("end-time")
		applicationSignals_listServiceStatesCmd.Flag("no-include-linked-accounts").Hidden = true
		applicationSignals_listServiceStatesCmd.MarkFlagRequired("start-time")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceStatesCmd)
}
