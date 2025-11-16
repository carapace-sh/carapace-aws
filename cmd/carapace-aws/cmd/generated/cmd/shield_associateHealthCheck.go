package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_associateHealthCheckCmd = &cobra.Command{
	Use:   "associate-health-check",
	Short: "Adds health-based detection to the Shield Advanced protection for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_associateHealthCheckCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_associateHealthCheckCmd).Standalone()

		shield_associateHealthCheckCmd.Flags().String("health-check-arn", "", "The Amazon Resource Name (ARN) of the health check to associate with the protection.")
		shield_associateHealthCheckCmd.Flags().String("protection-id", "", "The unique identifier (ID) for the [Protection]() object to add the health check association to.")
		shield_associateHealthCheckCmd.MarkFlagRequired("health-check-arn")
		shield_associateHealthCheckCmd.MarkFlagRequired("protection-id")
	})
	shieldCmd.AddCommand(shield_associateHealthCheckCmd)
}
