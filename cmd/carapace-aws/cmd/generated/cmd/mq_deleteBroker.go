package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_deleteBrokerCmd = &cobra.Command{
	Use:   "delete-broker",
	Short: "Deletes a broker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_deleteBrokerCmd).Standalone()

	mq_deleteBrokerCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
	mq_deleteBrokerCmd.MarkFlagRequired("broker-id")
	mqCmd.AddCommand(mq_deleteBrokerCmd)
}
