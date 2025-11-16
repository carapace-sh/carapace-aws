package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_disableHealthServiceAccessForOrganizationCmd = &cobra.Command{
	Use:   "disable-health-service-access-for-organization",
	Short: "Disables Health from working with Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_disableHealthServiceAccessForOrganizationCmd).Standalone()

	healthCmd.AddCommand(health_disableHealthServiceAccessForOrganizationCmd)
}
