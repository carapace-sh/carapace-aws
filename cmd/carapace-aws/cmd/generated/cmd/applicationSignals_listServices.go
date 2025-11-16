package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Returns a list of services that have been discovered by Application Signals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServicesCmd).Standalone()

	applicationSignals_listServicesCmd.Flags().String("aws-account-id", "", "Amazon Web Services Account ID.")
	applicationSignals_listServicesCmd.Flags().String("end-time", "", "The end of the time period to retrieve information about.")
	applicationSignals_listServicesCmd.Flags().Bool("include-linked-accounts", false, "If you are using this operation in a monitoring account, specify `true` to include services from source accounts in the returned data.")
	applicationSignals_listServicesCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	applicationSignals_listServicesCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of services.")
	applicationSignals_listServicesCmd.Flags().Bool("no-include-linked-accounts", false, "If you are using this operation in a monitoring account, specify `true` to include services from source accounts in the returned data.")
	applicationSignals_listServicesCmd.Flags().String("start-time", "", "The start of the time period to retrieve information about.")
	applicationSignals_listServicesCmd.MarkFlagRequired("end-time")
	applicationSignals_listServicesCmd.Flag("no-include-linked-accounts").Hidden = true
	applicationSignals_listServicesCmd.MarkFlagRequired("start-time")
	applicationSignalsCmd.AddCommand(applicationSignals_listServicesCmd)
}
