package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getOpsItemCmd = &cobra.Command{
	Use:   "get-ops-item",
	Short: "Get information about an OpsItem by using the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getOpsItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getOpsItemCmd).Standalone()

		ssm_getOpsItemCmd.Flags().String("ops-item-arn", "", "The OpsItem Amazon Resource Name (ARN).")
		ssm_getOpsItemCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem that you want to get.")
		ssm_getOpsItemCmd.MarkFlagRequired("ops-item-id")
	})
	ssmCmd.AddCommand(ssm_getOpsItemCmd)
}
