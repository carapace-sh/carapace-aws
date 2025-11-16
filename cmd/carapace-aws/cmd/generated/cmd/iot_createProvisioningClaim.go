package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createProvisioningClaimCmd = &cobra.Command{
	Use:   "create-provisioning-claim",
	Short: "Creates a provisioning claim.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createProvisioningClaimCmd).Standalone()

	iot_createProvisioningClaimCmd.Flags().String("template-name", "", "The name of the provisioning template to use.")
	iot_createProvisioningClaimCmd.MarkFlagRequired("template-name")
	iotCmd.AddCommand(iot_createProvisioningClaimCmd)
}
