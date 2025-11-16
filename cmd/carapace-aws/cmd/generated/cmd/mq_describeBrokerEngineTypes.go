package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeBrokerEngineTypesCmd = &cobra.Command{
	Use:   "describe-broker-engine-types",
	Short: "Describe available engine types and versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeBrokerEngineTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_describeBrokerEngineTypesCmd).Standalone()

		mq_describeBrokerEngineTypesCmd.Flags().String("engine-type", "", "Filter response by engine type.")
		mq_describeBrokerEngineTypesCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
		mq_describeBrokerEngineTypesCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
	})
	mqCmd.AddCommand(mq_describeBrokerEngineTypesCmd)
}
