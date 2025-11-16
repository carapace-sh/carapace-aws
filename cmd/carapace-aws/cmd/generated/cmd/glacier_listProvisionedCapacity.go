package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listProvisionedCapacityCmd = &cobra.Command{
	Use:   "list-provisioned-capacity",
	Short: "This operation lists the provisioned capacity units for the specified AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listProvisionedCapacityCmd).Standalone()

	glacier_listProvisionedCapacityCmd.Flags().String("account-id", "", "The AWS account ID of the account that owns the vault.")
	glacier_listProvisionedCapacityCmd.MarkFlagRequired("account-id")
	glacierCmd.AddCommand(glacier_listProvisionedCapacityCmd)
}
