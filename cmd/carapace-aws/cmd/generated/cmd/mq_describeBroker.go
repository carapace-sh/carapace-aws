package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeBrokerCmd = &cobra.Command{
	Use:   "describe-broker",
	Short: "Returns information about the specified broker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeBrokerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_describeBrokerCmd).Standalone()

		mq_describeBrokerCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_describeBrokerCmd.MarkFlagRequired("broker-id")
	})
	mqCmd.AddCommand(mq_describeBrokerCmd)
}
