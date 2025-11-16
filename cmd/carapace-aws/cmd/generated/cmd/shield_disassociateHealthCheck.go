package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_disassociateHealthCheckCmd = &cobra.Command{
	Use:   "disassociate-health-check",
	Short: "Removes health-based detection from the Shield Advanced protection for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_disassociateHealthCheckCmd).Standalone()

	shield_disassociateHealthCheckCmd.Flags().String("health-check-arn", "", "The Amazon Resource Name (ARN) of the health check that is associated with the protection.")
	shield_disassociateHealthCheckCmd.Flags().String("protection-id", "", "The unique identifier (ID) for the [Protection]() object to remove the health check association from.")
	shield_disassociateHealthCheckCmd.MarkFlagRequired("health-check-arn")
	shield_disassociateHealthCheckCmd.MarkFlagRequired("protection-id")
	shieldCmd.AddCommand(shield_disassociateHealthCheckCmd)
}
