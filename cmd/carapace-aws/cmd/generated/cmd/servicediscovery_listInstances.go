package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "Lists summary information about the instances that you registered by using a specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_listInstancesCmd).Standalone()

	servicediscovery_listInstancesCmd.Flags().String("max-results", "", "The maximum number of instances that you want Cloud Map to return in the response to a `ListInstances` request.")
	servicediscovery_listInstancesCmd.Flags().String("next-token", "", "For the first `ListInstances` request, omit this value.")
	servicediscovery_listInstancesCmd.Flags().String("service-id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to list instances for.")
	servicediscovery_listInstancesCmd.MarkFlagRequired("service-id")
	servicediscoveryCmd.AddCommand(servicediscovery_listInstancesCmd)
}
