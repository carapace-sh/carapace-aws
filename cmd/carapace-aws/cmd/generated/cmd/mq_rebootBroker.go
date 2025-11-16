package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_rebootBrokerCmd = &cobra.Command{
	Use:   "reboot-broker",
	Short: "Reboots a broker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_rebootBrokerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_rebootBrokerCmd).Standalone()

		mq_rebootBrokerCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_rebootBrokerCmd.MarkFlagRequired("broker-id")
	})
	mqCmd.AddCommand(mq_rebootBrokerCmd)
}
