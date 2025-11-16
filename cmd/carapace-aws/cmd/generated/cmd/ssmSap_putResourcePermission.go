package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_putResourcePermissionCmd = &cobra.Command{
	Use:   "put-resource-permission",
	Short: "Adds permissions to the target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_putResourcePermissionCmd).Standalone()

	ssmSap_putResourcePermissionCmd.Flags().String("action-type", "", "")
	ssmSap_putResourcePermissionCmd.Flags().String("resource-arn", "", "")
	ssmSap_putResourcePermissionCmd.Flags().String("source-resource-arn", "", "")
	ssmSap_putResourcePermissionCmd.MarkFlagRequired("action-type")
	ssmSap_putResourcePermissionCmd.MarkFlagRequired("resource-arn")
	ssmSap_putResourcePermissionCmd.MarkFlagRequired("source-resource-arn")
	ssmSapCmd.AddCommand(ssmSap_putResourcePermissionCmd)
}
