package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Lists summary information for all the services that are associated with one or more namespaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_listServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_listServicesCmd).Standalone()

		servicediscovery_listServicesCmd.Flags().String("filters", "", "A complex type that contains specifications for the namespaces that you want to list services for.")
		servicediscovery_listServicesCmd.Flags().String("max-results", "", "The maximum number of services that you want Cloud Map to return in the response to a `ListServices` request.")
		servicediscovery_listServicesCmd.Flags().String("next-token", "", "For the first `ListServices` request, omit this value.")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_listServicesCmd)
}
