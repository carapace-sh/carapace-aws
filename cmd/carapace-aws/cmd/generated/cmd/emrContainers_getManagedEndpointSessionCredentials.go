package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_getManagedEndpointSessionCredentialsCmd = &cobra.Command{
	Use:   "get-managed-endpoint-session-credentials",
	Short: "Generate a session token to connect to a managed endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_getManagedEndpointSessionCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_getManagedEndpointSessionCredentialsCmd).Standalone()

		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("client-token", "", "The client idempotency token of the job run request.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("credential-type", "", "Type of the token requested.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("duration-in-seconds", "", "Duration in seconds for which the session token is valid.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("endpoint-identifier", "", "The ARN of the managed endpoint for which the request is submitted.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("execution-role-arn", "", "The IAM Execution Role ARN that will be used by the job run.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("log-context", "", "String identifier used to separate sections of the execution logs uploaded to S3.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.Flags().String("virtual-cluster-identifier", "", "The ARN of the Virtual Cluster which the Managed Endpoint belongs to.")
		emrContainers_getManagedEndpointSessionCredentialsCmd.MarkFlagRequired("credential-type")
		emrContainers_getManagedEndpointSessionCredentialsCmd.MarkFlagRequired("endpoint-identifier")
		emrContainers_getManagedEndpointSessionCredentialsCmd.MarkFlagRequired("execution-role-arn")
		emrContainers_getManagedEndpointSessionCredentialsCmd.MarkFlagRequired("virtual-cluster-identifier")
	})
	emrContainersCmd.AddCommand(emrContainers_getManagedEndpointSessionCredentialsCmd)
}
