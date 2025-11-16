package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createAvailabilityConfigurationCmd = &cobra.Command{
	Use:   "create-availability-configuration",
	Short: "Creates an `AvailabilityConfiguration` for the given WorkMail organization and domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createAvailabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_createAvailabilityConfigurationCmd).Standalone()

		workmail_createAvailabilityConfigurationCmd.Flags().String("client-token", "", "An idempotent token that ensures that an API request is executed only once.")
		workmail_createAvailabilityConfigurationCmd.Flags().String("domain-name", "", "The domain to which the provider applies.")
		workmail_createAvailabilityConfigurationCmd.Flags().String("ews-provider", "", "Exchange Web Services (EWS) availability provider definition.")
		workmail_createAvailabilityConfigurationCmd.Flags().String("lambda-provider", "", "Lambda availability provider definition.")
		workmail_createAvailabilityConfigurationCmd.Flags().String("organization-id", "", "The WorkMail organization for which the `AvailabilityConfiguration` will be created.")
		workmail_createAvailabilityConfigurationCmd.MarkFlagRequired("domain-name")
		workmail_createAvailabilityConfigurationCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_createAvailabilityConfigurationCmd)
}
