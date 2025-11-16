package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_createManagedEndpointCmd = &cobra.Command{
	Use:   "create-managed-endpoint",
	Short: "Creates a managed endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_createManagedEndpointCmd).Standalone()

	emrContainers_createManagedEndpointCmd.Flags().String("certificate-arn", "", "The certificate ARN provided by users for the managed endpoint.")
	emrContainers_createManagedEndpointCmd.Flags().String("client-token", "", "The client idempotency token for this create call.")
	emrContainers_createManagedEndpointCmd.Flags().String("configuration-overrides", "", "The configuration settings that will be used to override existing configurations.")
	emrContainers_createManagedEndpointCmd.Flags().String("execution-role-arn", "", "The ARN of the execution role.")
	emrContainers_createManagedEndpointCmd.Flags().String("name", "", "The name of the managed endpoint.")
	emrContainers_createManagedEndpointCmd.Flags().String("release-label", "", "The Amazon EMR release version.")
	emrContainers_createManagedEndpointCmd.Flags().String("tags", "", "The tags of the managed endpoint.")
	emrContainers_createManagedEndpointCmd.Flags().String("type", "", "The type of the managed endpoint.")
	emrContainers_createManagedEndpointCmd.Flags().String("virtual-cluster-id", "", "The ID of the virtual cluster for which a managed endpoint is created.")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("client-token")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("execution-role-arn")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("name")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("release-label")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("type")
	emrContainers_createManagedEndpointCmd.MarkFlagRequired("virtual-cluster-id")
	emrContainersCmd.AddCommand(emrContainers_createManagedEndpointCmd)
}
