package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_listBrokersCmd = &cobra.Command{
	Use:   "list-brokers",
	Short: "Returns a list of all brokers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_listBrokersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_listBrokersCmd).Standalone()

		mq_listBrokersCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
		mq_listBrokersCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
	})
	mqCmd.AddCommand(mq_listBrokersCmd)
}
