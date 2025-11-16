package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteOpsItemCmd = &cobra.Command{
	Use:   "delete-ops-item",
	Short: "Delete an OpsItem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteOpsItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteOpsItemCmd).Standalone()

		ssm_deleteOpsItemCmd.Flags().String("ops-item-id", "", "The ID of the OpsItem that you want to delete.")
		ssm_deleteOpsItemCmd.MarkFlagRequired("ops-item-id")
	})
	ssmCmd.AddCommand(ssm_deleteOpsItemCmd)
}
