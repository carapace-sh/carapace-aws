package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getOperationCmd = &cobra.Command{
	Use:   "get-operation",
	Short: "Gets the details of an operation by specifying the operation ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_getOperationCmd).Standalone()

		ssmSap_getOperationCmd.Flags().String("operation-id", "", "The ID of the operation.")
		ssmSap_getOperationCmd.MarkFlagRequired("operation-id")
	})
	ssmSapCmd.AddCommand(ssmSap_getOperationCmd)
}
