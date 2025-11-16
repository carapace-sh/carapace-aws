package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getResourcePermissionCmd = &cobra.Command{
	Use:   "get-resource-permission",
	Short: "Gets permissions associated with the target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getResourcePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_getResourcePermissionCmd).Standalone()

		ssmSap_getResourcePermissionCmd.Flags().String("action-type", "", "")
		ssmSap_getResourcePermissionCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		ssmSap_getResourcePermissionCmd.MarkFlagRequired("resource-arn")
	})
	ssmSapCmd.AddCommand(ssmSap_getResourcePermissionCmd)
}
