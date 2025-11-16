package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateSmblocalGroupsCmd = &cobra.Command{
	Use:   "update-smblocal-groups",
	Short: "Updates the list of Active Directory users and groups that have special permissions for SMB file shares on the gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateSmblocalGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateSmblocalGroupsCmd).Standalone()

		storagegateway_updateSmblocalGroupsCmd.Flags().String("gateway-arn", "", "")
		storagegateway_updateSmblocalGroupsCmd.Flags().String("smblocal-groups", "", "A list of Active Directory users and groups that you want to grant special permissions for SMB file shares on the gateway.")
		storagegateway_updateSmblocalGroupsCmd.MarkFlagRequired("gateway-arn")
		storagegateway_updateSmblocalGroupsCmd.MarkFlagRequired("smblocal-groups")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateSmblocalGroupsCmd)
}
