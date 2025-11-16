package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listVaultsCmd = &cobra.Command{
	Use:   "list-vaults",
	Short: "This operation lists all vaults owned by the calling user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listVaultsCmd).Standalone()

	glacier_listVaultsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID.")
	glacier_listVaultsCmd.Flags().String("limit", "", "The maximum number of vaults to be returned.")
	glacier_listVaultsCmd.Flags().String("marker", "", "A string used for pagination.")
	glacier_listVaultsCmd.MarkFlagRequired("account-id")
	glacierCmd.AddCommand(glacier_listVaultsCmd)
}
