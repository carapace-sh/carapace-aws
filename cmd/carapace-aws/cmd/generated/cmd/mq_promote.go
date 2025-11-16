package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promotes a data replication replica broker to the primary broker role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_promoteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_promoteCmd).Standalone()

		mq_promoteCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_promoteCmd.Flags().String("mode", "", "The Promote mode requested.")
		mq_promoteCmd.MarkFlagRequired("broker-id")
		mq_promoteCmd.MarkFlagRequired("mode")
	})
	mqCmd.AddCommand(mq_promoteCmd)
}
