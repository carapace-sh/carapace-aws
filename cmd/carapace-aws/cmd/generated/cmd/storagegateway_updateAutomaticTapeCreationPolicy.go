package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateAutomaticTapeCreationPolicyCmd = &cobra.Command{
	Use:   "update-automatic-tape-creation-policy",
	Short: "Updates the automatic tape creation policy of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateAutomaticTapeCreationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateAutomaticTapeCreationPolicyCmd).Standalone()

		storagegateway_updateAutomaticTapeCreationPolicyCmd.Flags().String("automatic-tape-creation-rules", "", "An automatic tape creation policy consists of a list of automatic tape creation rules.")
		storagegateway_updateAutomaticTapeCreationPolicyCmd.Flags().String("gateway-arn", "", "")
		storagegateway_updateAutomaticTapeCreationPolicyCmd.MarkFlagRequired("automatic-tape-creation-rules")
		storagegateway_updateAutomaticTapeCreationPolicyCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateAutomaticTapeCreationPolicyCmd)
}
