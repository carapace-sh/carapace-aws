package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_deleteResourcePermissionCmd = &cobra.Command{
	Use:   "delete-resource-permission",
	Short: "Removes permissions associated with the target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_deleteResourcePermissionCmd).Standalone()

	ssmSap_deleteResourcePermissionCmd.Flags().String("action-type", "", "Delete or restore the permissions on the target database.")
	ssmSap_deleteResourcePermissionCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	ssmSap_deleteResourcePermissionCmd.Flags().String("source-resource-arn", "", "The Amazon Resource Name (ARN) of the source resource.")
	ssmSap_deleteResourcePermissionCmd.MarkFlagRequired("resource-arn")
	ssmSapCmd.AddCommand(ssmSap_deleteResourcePermissionCmd)
}
