package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeBrokerInstanceOptionsCmd = &cobra.Command{
	Use:   "describe-broker-instance-options",
	Short: "Describe available broker instance options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeBrokerInstanceOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_describeBrokerInstanceOptionsCmd).Standalone()

		mq_describeBrokerInstanceOptionsCmd.Flags().String("engine-type", "", "Filter response by engine type.")
		mq_describeBrokerInstanceOptionsCmd.Flags().String("host-instance-type", "", "Filter response by host instance type.")
		mq_describeBrokerInstanceOptionsCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
		mq_describeBrokerInstanceOptionsCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
		mq_describeBrokerInstanceOptionsCmd.Flags().String("storage-type", "", "Filter response by storage type.")
	})
	mqCmd.AddCommand(mq_describeBrokerInstanceOptionsCmd)
}
