package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteAvailabilityConfigurationCmd = &cobra.Command{
	Use:   "delete-availability-configuration",
	Short: "Deletes the `AvailabilityConfiguration` for the given WorkMail organization and domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteAvailabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deleteAvailabilityConfigurationCmd).Standalone()

		workmail_deleteAvailabilityConfigurationCmd.Flags().String("domain-name", "", "The domain for which the `AvailabilityConfiguration` will be deleted.")
		workmail_deleteAvailabilityConfigurationCmd.Flags().String("organization-id", "", "The WorkMail organization for which the `AvailabilityConfiguration` will be deleted.")
		workmail_deleteAvailabilityConfigurationCmd.MarkFlagRequired("domain-name")
		workmail_deleteAvailabilityConfigurationCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_deleteAvailabilityConfigurationCmd)
}
