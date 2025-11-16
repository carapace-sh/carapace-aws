package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listCloudConnectorsCmd = &cobra.Command{
	Use:   "list-cloud-connectors",
	Short: "Returns a list of connectors filtered by its Lambda Amazon Resource Name (ARN) and `type`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listCloudConnectorsCmd).Standalone()

	iotManagedIntegrations_listCloudConnectorsCmd.Flags().String("lambda-arn", "", "The Amazon Resource Name (ARN) of the Lambda function to filter cloud connectors by.")
	iotManagedIntegrations_listCloudConnectorsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotManagedIntegrations_listCloudConnectorsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	iotManagedIntegrations_listCloudConnectorsCmd.Flags().String("type", "", "The type of cloud connectors to filter by when listing available connectors.")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listCloudConnectorsCmd)
}
