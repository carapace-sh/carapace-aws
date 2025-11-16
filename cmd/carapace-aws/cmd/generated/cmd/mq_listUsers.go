package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Returns a list of all ActiveMQ users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_listUsersCmd).Standalone()

	mq_listUsersCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
	mq_listUsersCmd.Flags().String("max-results", "", "The maximum number of brokers that Amazon MQ can return per page (20 by default).")
	mq_listUsersCmd.Flags().String("next-token", "", "The token that specifies the next page of results Amazon MQ should return.")
	mq_listUsersCmd.MarkFlagRequired("broker-id")
	mqCmd.AddCommand(mq_listUsersCmd)
}
