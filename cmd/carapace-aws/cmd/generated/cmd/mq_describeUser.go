package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Returns information about an ActiveMQ user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_describeUserCmd).Standalone()

		mq_describeUserCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_describeUserCmd.Flags().String("username", "", "The username of the ActiveMQ user.")
		mq_describeUserCmd.MarkFlagRequired("broker-id")
		mq_describeUserCmd.MarkFlagRequired("username")
	})
	mqCmd.AddCommand(mq_describeUserCmd)
}
