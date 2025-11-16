package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_startConfigurationPolicyAssociationCmd = &cobra.Command{
	Use:   "start-configuration-policy-association",
	Short: "Associates a target account, organizational unit, or the root with a specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_startConfigurationPolicyAssociationCmd).Standalone()

	securityhub_startConfigurationPolicyAssociationCmd.Flags().String("configuration-policy-identifier", "", "The Amazon Resource Name (ARN) of a configuration policy, the universally unique identifier (UUID) of a configuration policy, or a value of `SELF_MANAGED_SECURITY_HUB` for a self-managed configuration.")
	securityhub_startConfigurationPolicyAssociationCmd.Flags().String("target", "", "The identifier of the target account, organizational unit, or the root to associate with the specified configuration.")
	securityhub_startConfigurationPolicyAssociationCmd.MarkFlagRequired("configuration-policy-identifier")
	securityhub_startConfigurationPolicyAssociationCmd.MarkFlagRequired("target")
	securityhubCmd.AddCommand(securityhub_startConfigurationPolicyAssociationCmd)
}
