package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_purchaseProvisionedCapacityCmd = &cobra.Command{
	Use:   "purchase-provisioned-capacity",
	Short: "This operation purchases a provisioned capacity unit for an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_purchaseProvisionedCapacityCmd).Standalone()

	glacier_purchaseProvisionedCapacityCmd.Flags().String("account-id", "", "The AWS account ID of the account that owns the vault.")
	glacier_purchaseProvisionedCapacityCmd.MarkFlagRequired("account-id")
	glacierCmd.AddCommand(glacier_purchaseProvisionedCapacityCmd)
}
