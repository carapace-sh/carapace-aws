package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_putHypervisorPropertyMappingsCmd = &cobra.Command{
	Use:   "put-hypervisor-property-mappings",
	Short: "This action sets the property mappings for the specified hypervisor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_putHypervisorPropertyMappingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_putHypervisorPropertyMappingsCmd).Standalone()

		backupGateway_putHypervisorPropertyMappingsCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor.")
		backupGateway_putHypervisorPropertyMappingsCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
		backupGateway_putHypervisorPropertyMappingsCmd.Flags().String("vmware-to-aws-tag-mappings", "", "This action requests the mappings of on-premises VMware tags to the Amazon Web Services tags.")
		backupGateway_putHypervisorPropertyMappingsCmd.MarkFlagRequired("hypervisor-arn")
		backupGateway_putHypervisorPropertyMappingsCmd.MarkFlagRequired("iam-role-arn")
		backupGateway_putHypervisorPropertyMappingsCmd.MarkFlagRequired("vmware-to-aws-tag-mappings")
	})
	backupGatewayCmd.AddCommand(backupGateway_putHypervisorPropertyMappingsCmd)
}
