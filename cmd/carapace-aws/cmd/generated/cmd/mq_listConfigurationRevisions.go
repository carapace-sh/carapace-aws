package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_listConfigurationRevisionsCmd = &cobra.Command{
	Use:   "list-configuration-revisions",
	Short: "Returns a list of all revisions for the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_listConfigurationRevisionsCmd).Standalone()

	mq_listConfigurationRevisionsCmd.Flags().String("configuration-id", "", "The unique ID that Amazon MQ generates for the configuration.")
	mq_listConfigurationRevisionsCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
	mq_listConfigurationRevisionsCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
	mq_listConfigurationRevisionsCmd.MarkFlagRequired("configuration-id")
	mqCmd.AddCommand(mq_listConfigurationRevisionsCmd)
}
