package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_enableHealthServiceAccessForOrganizationCmd = &cobra.Command{
	Use:   "enable-health-service-access-for-organization",
	Short: "Enables Health to work with Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_enableHealthServiceAccessForOrganizationCmd).Standalone()

	healthCmd.AddCommand(health_enableHealthServiceAccessForOrganizationCmd)
}
