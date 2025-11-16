package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_getHypervisorPropertyMappingsCmd = &cobra.Command{
	Use:   "get-hypervisor-property-mappings",
	Short: "This action retrieves the property mappings for the specified hypervisor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_getHypervisorPropertyMappingsCmd).Standalone()

	backupGateway_getHypervisorPropertyMappingsCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor.")
	backupGateway_getHypervisorPropertyMappingsCmd.MarkFlagRequired("hypervisor-arn")
	backupGatewayCmd.AddCommand(backupGateway_getHypervisorPropertyMappingsCmd)
}
