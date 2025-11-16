package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listAvailabilityConfigurationsCmd = &cobra.Command{
	Use:   "list-availability-configurations",
	Short: "List all the `AvailabilityConfiguration`'s for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listAvailabilityConfigurationsCmd).Standalone()

	workmail_listAvailabilityConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listAvailabilityConfigurationsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmail_listAvailabilityConfigurationsCmd.Flags().String("organization-id", "", "The WorkMail organization for which the `AvailabilityConfiguration`'s will be listed.")
	workmail_listAvailabilityConfigurationsCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listAvailabilityConfigurationsCmd)
}
