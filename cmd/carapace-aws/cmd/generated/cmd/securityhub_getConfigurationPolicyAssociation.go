package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getConfigurationPolicyAssociationCmd = &cobra.Command{
	Use:   "get-configuration-policy-association",
	Short: "Returns the association between a configuration and a target account, organizational unit, or the root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getConfigurationPolicyAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getConfigurationPolicyAssociationCmd).Standalone()

		securityhub_getConfigurationPolicyAssociationCmd.Flags().String("target", "", "The target account ID, organizational unit ID, or the root ID to retrieve the association for.")
		securityhub_getConfigurationPolicyAssociationCmd.MarkFlagRequired("target")
	})
	securityhubCmd.AddCommand(securityhub_getConfigurationPolicyAssociationCmd)
}
