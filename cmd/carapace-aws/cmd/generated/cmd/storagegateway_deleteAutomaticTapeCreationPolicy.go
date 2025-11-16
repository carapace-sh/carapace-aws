package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteAutomaticTapeCreationPolicyCmd = &cobra.Command{
	Use:   "delete-automatic-tape-creation-policy",
	Short: "Deletes the automatic tape creation policy of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteAutomaticTapeCreationPolicyCmd).Standalone()

	storagegateway_deleteAutomaticTapeCreationPolicyCmd.Flags().String("gateway-arn", "", "")
	storagegateway_deleteAutomaticTapeCreationPolicyCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_deleteAutomaticTapeCreationPolicyCmd)
}
