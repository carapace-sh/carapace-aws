package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_executeActionCmd = &cobra.Command{
	Use:   "execute-action",
	Short: "Executes an action on a target resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_executeActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_executeActionCmd).Standalone()

		iotsitewise_executeActionCmd.Flags().String("action-definition-id", "", "The ID of the action definition.")
		iotsitewise_executeActionCmd.Flags().String("action-payload", "", "The JSON payload of the action.")
		iotsitewise_executeActionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_executeActionCmd.Flags().String("resolve-to", "", "The detailed resource this action resolves to.")
		iotsitewise_executeActionCmd.Flags().String("target-resource", "", "The resource the action will be taken on.")
		iotsitewise_executeActionCmd.MarkFlagRequired("action-definition-id")
		iotsitewise_executeActionCmd.MarkFlagRequired("action-payload")
		iotsitewise_executeActionCmd.MarkFlagRequired("target-resource")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_executeActionCmd)
}
