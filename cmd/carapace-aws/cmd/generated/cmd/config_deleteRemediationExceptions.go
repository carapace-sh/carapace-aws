package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteRemediationExceptionsCmd = &cobra.Command{
	Use:   "delete-remediation-exceptions",
	Short: "Deletes one or more remediation exceptions mentioned in the resource keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteRemediationExceptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteRemediationExceptionsCmd).Standalone()

		config_deleteRemediationExceptionsCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want to delete remediation exception configuration.")
		config_deleteRemediationExceptionsCmd.Flags().String("resource-keys", "", "An exception list of resource exception keys to be processed with the current request.")
		config_deleteRemediationExceptionsCmd.MarkFlagRequired("config-rule-name")
		config_deleteRemediationExceptionsCmd.MarkFlagRequired("resource-keys")
	})
	configCmd.AddCommand(config_deleteRemediationExceptionsCmd)
}
