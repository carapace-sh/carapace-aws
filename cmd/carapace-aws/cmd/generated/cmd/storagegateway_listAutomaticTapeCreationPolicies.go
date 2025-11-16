package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listAutomaticTapeCreationPoliciesCmd = &cobra.Command{
	Use:   "list-automatic-tape-creation-policies",
	Short: "Lists the automatic tape creation policies for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listAutomaticTapeCreationPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listAutomaticTapeCreationPoliciesCmd).Standalone()

		storagegateway_listAutomaticTapeCreationPoliciesCmd.Flags().String("gateway-arn", "", "")
	})
	storagegatewayCmd.AddCommand(storagegateway_listAutomaticTapeCreationPoliciesCmd)
}
