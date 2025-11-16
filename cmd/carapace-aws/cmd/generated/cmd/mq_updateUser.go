package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates the information for an ActiveMQ user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_updateUserCmd).Standalone()

	mq_updateUserCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
	mq_updateUserCmd.Flags().String("console-access", "", "Enables access to the the ActiveMQ Web Console for the ActiveMQ user.")
	mq_updateUserCmd.Flags().String("groups", "", "The list of groups (20 maximum) to which the ActiveMQ user belongs.")
	mq_updateUserCmd.Flags().String("password", "", "The password of the user.")
	mq_updateUserCmd.Flags().String("replication-user", "", "Defines whether the user is intended for data replication.")
	mq_updateUserCmd.Flags().String("username", "", "The username of the ActiveMQ user.")
	mq_updateUserCmd.MarkFlagRequired("broker-id")
	mq_updateUserCmd.MarkFlagRequired("username")
	mqCmd.AddCommand(mq_updateUserCmd)
}
