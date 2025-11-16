package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_startConfigurationPolicyDisassociationCmd = &cobra.Command{
	Use:   "start-configuration-policy-disassociation",
	Short: "Disassociates a target account, organizational unit, or the root from a specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_startConfigurationPolicyDisassociationCmd).Standalone()

	securityhub_startConfigurationPolicyDisassociationCmd.Flags().String("configuration-policy-identifier", "", "The Amazon Resource Name (ARN) of a configuration policy, the universally unique identifier (UUID) of a configuration policy, or a value of `SELF_MANAGED_SECURITY_HUB` for a self-managed configuration.")
	securityhub_startConfigurationPolicyDisassociationCmd.Flags().String("target", "", "The identifier of the target account, organizational unit, or the root to disassociate from the specified configuration.")
	securityhub_startConfigurationPolicyDisassociationCmd.MarkFlagRequired("configuration-policy-identifier")
	securityhubCmd.AddCommand(securityhub_startConfigurationPolicyDisassociationCmd)
}
