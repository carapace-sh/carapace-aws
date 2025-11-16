package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_batchGetProfileCmd = &cobra.Command{
	Use:   "batch-get-profile",
	Short: "Get a batch of profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_batchGetProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_batchGetProfileCmd).Standalone()

		customerProfiles_batchGetProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_batchGetProfileCmd.Flags().String("profile-ids", "", "List of unique identifiers for customer profiles to retrieve.")
		customerProfiles_batchGetProfileCmd.MarkFlagRequired("domain-name")
		customerProfiles_batchGetProfileCmd.MarkFlagRequired("profile-ids")
	})
	customerProfilesCmd.AddCommand(customerProfiles_batchGetProfileCmd)
}
