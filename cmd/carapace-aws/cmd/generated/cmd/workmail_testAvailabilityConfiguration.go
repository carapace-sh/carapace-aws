package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_testAvailabilityConfigurationCmd = &cobra.Command{
	Use:   "test-availability-configuration",
	Short: "Performs a test on an availability provider to ensure that access is allowed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_testAvailabilityConfigurationCmd).Standalone()

	workmail_testAvailabilityConfigurationCmd.Flags().String("domain-name", "", "The domain to which the provider applies.")
	workmail_testAvailabilityConfigurationCmd.Flags().String("ews-provider", "", "")
	workmail_testAvailabilityConfigurationCmd.Flags().String("lambda-provider", "", "")
	workmail_testAvailabilityConfigurationCmd.Flags().String("organization-id", "", "The WorkMail organization where the availability provider will be tested.")
	workmail_testAvailabilityConfigurationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_testAvailabilityConfigurationCmd)
}
