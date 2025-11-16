package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putRemediationExceptionsCmd = &cobra.Command{
	Use:   "put-remediation-exceptions",
	Short: "A remediation exception is when a specified resource is no longer considered for auto-remediation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putRemediationExceptionsCmd).Standalone()

	config_putRemediationExceptionsCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want to create remediation exception.")
	config_putRemediationExceptionsCmd.Flags().String("expiration-time", "", "The exception is automatically deleted after the expiration date.")
	config_putRemediationExceptionsCmd.Flags().String("message", "", "The message contains an explanation of the exception.")
	config_putRemediationExceptionsCmd.Flags().String("resource-keys", "", "An exception list of resource exception keys to be processed with the current request.")
	config_putRemediationExceptionsCmd.MarkFlagRequired("config-rule-name")
	config_putRemediationExceptionsCmd.MarkFlagRequired("resource-keys")
	configCmd.AddCommand(config_putRemediationExceptionsCmd)
}
