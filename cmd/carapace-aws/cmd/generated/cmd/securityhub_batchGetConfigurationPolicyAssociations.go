package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchGetConfigurationPolicyAssociationsCmd = &cobra.Command{
	Use:   "batch-get-configuration-policy-associations",
	Short: "Returns associations between an Security Hub configuration and a batch of target accounts, organizational units, or the root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchGetConfigurationPolicyAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchGetConfigurationPolicyAssociationsCmd).Standalone()

		securityhub_batchGetConfigurationPolicyAssociationsCmd.Flags().String("configuration-policy-association-identifiers", "", "Specifies one or more target account IDs, organizational unit (OU) IDs, or the root ID to retrieve associations for.")
		securityhub_batchGetConfigurationPolicyAssociationsCmd.MarkFlagRequired("configuration-policy-association-identifiers")
	})
	securityhubCmd.AddCommand(securityhub_batchGetConfigurationPolicyAssociationsCmd)
}
