package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates an ActiveMQ user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_createUserCmd).Standalone()

		mq_createUserCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_createUserCmd.Flags().String("console-access", "", "Enables access to the ActiveMQ Web Console for the ActiveMQ user.")
		mq_createUserCmd.Flags().String("groups", "", "The list of groups (20 maximum) to which the ActiveMQ user belongs.")
		mq_createUserCmd.Flags().String("password", "", "Required.")
		mq_createUserCmd.Flags().String("replication-user", "", "Defines if this user is intended for CRDR replication purposes.")
		mq_createUserCmd.Flags().String("username", "", "The username of the ActiveMQ user.")
		mq_createUserCmd.MarkFlagRequired("broker-id")
		mq_createUserCmd.MarkFlagRequired("password")
		mq_createUserCmd.MarkFlagRequired("username")
	})
	mqCmd.AddCommand(mq_createUserCmd)
}
