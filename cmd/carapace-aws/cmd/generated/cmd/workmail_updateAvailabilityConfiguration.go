package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateAvailabilityConfigurationCmd = &cobra.Command{
	Use:   "update-availability-configuration",
	Short: "Updates an existing `AvailabilityConfiguration` for the given WorkMail organization and domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateAvailabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_updateAvailabilityConfigurationCmd).Standalone()

		workmail_updateAvailabilityConfigurationCmd.Flags().String("domain-name", "", "The domain to which the provider applies the availability configuration.")
		workmail_updateAvailabilityConfigurationCmd.Flags().String("ews-provider", "", "The EWS availability provider definition.")
		workmail_updateAvailabilityConfigurationCmd.Flags().String("lambda-provider", "", "The Lambda availability provider definition.")
		workmail_updateAvailabilityConfigurationCmd.Flags().String("organization-id", "", "The WorkMail organization for which the `AvailabilityConfiguration` will be updated.")
		workmail_updateAvailabilityConfigurationCmd.MarkFlagRequired("domain-name")
		workmail_updateAvailabilityConfigurationCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_updateAvailabilityConfigurationCmd)
}
