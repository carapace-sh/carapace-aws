package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_updateAvailabilityOptionsCmd = &cobra.Command{
	Use:   "update-availability-options",
	Short: "Configures the availability options for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_updateAvailabilityOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_updateAvailabilityOptionsCmd).Standalone()

		cloudsearch_updateAvailabilityOptionsCmd.Flags().String("domain-name", "", "")
		cloudsearch_updateAvailabilityOptionsCmd.Flags().Bool("multi-az", false, "You expand an existing search domain to a second Availability Zone by setting the Multi-AZ option to true.")
		cloudsearch_updateAvailabilityOptionsCmd.Flags().Bool("no-multi-az", false, "You expand an existing search domain to a second Availability Zone by setting the Multi-AZ option to true.")
		cloudsearch_updateAvailabilityOptionsCmd.MarkFlagRequired("domain-name")
		cloudsearch_updateAvailabilityOptionsCmd.MarkFlagRequired("multi-az")
		cloudsearch_updateAvailabilityOptionsCmd.Flag("no-multi-az").Hidden = true
		cloudsearch_updateAvailabilityOptionsCmd.MarkFlagRequired("no-multi-az")
	})
	cloudsearchCmd.AddCommand(cloudsearch_updateAvailabilityOptionsCmd)
}
