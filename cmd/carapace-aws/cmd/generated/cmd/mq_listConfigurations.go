package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_listConfigurationsCmd = &cobra.Command{
	Use:   "list-configurations",
	Short: "Returns a list of all configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_listConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_listConfigurationsCmd).Standalone()

		mq_listConfigurationsCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
		mq_listConfigurationsCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
	})
	mqCmd.AddCommand(mq_listConfigurationsCmd)
}
