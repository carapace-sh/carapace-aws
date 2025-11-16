package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeHealthServiceStatusForOrganizationCmd = &cobra.Command{
	Use:   "describe-health-service-status-for-organization",
	Short: "This operation provides status information on enabling or disabling Health to work with your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeHealthServiceStatusForOrganizationCmd).Standalone()

	healthCmd.AddCommand(health_describeHealthServiceStatusForOrganizationCmd)
}
