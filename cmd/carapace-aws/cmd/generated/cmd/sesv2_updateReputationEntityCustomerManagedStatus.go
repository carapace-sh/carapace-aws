package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateReputationEntityCustomerManagedStatusCmd = &cobra.Command{
	Use:   "update-reputation-entity-customer-managed-status",
	Short: "Update the customer-managed sending status for a reputation entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateReputationEntityCustomerManagedStatusCmd).Standalone()

	sesv2_updateReputationEntityCustomerManagedStatusCmd.Flags().String("reputation-entity-reference", "", "The unique identifier for the reputation entity.")
	sesv2_updateReputationEntityCustomerManagedStatusCmd.Flags().String("reputation-entity-type", "", "The type of reputation entity.")
	sesv2_updateReputationEntityCustomerManagedStatusCmd.Flags().String("sending-status", "", "The new customer-managed sending status for the reputation entity.")
	sesv2_updateReputationEntityCustomerManagedStatusCmd.MarkFlagRequired("reputation-entity-reference")
	sesv2_updateReputationEntityCustomerManagedStatusCmd.MarkFlagRequired("reputation-entity-type")
	sesv2_updateReputationEntityCustomerManagedStatusCmd.MarkFlagRequired("sending-status")
	sesv2Cmd.AddCommand(sesv2_updateReputationEntityCustomerManagedStatusCmd)
}
