package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "Amazon MQ is a managed message broker service for Apache ActiveMQ and RabbitMQ that makes it easy to set up and operate message brokers in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mqCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mqCmd).Standalone()

	})
	rootCmd.AddCommand(mqCmd)
}
