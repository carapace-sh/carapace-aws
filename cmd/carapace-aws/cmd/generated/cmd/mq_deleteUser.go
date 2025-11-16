package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes an ActiveMQ user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_deleteUserCmd).Standalone()

	mq_deleteUserCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
	mq_deleteUserCmd.Flags().String("username", "", "The username of the ActiveMQ user.")
	mq_deleteUserCmd.MarkFlagRequired("broker-id")
	mq_deleteUserCmd.MarkFlagRequired("username")
	mqCmd.AddCommand(mq_deleteUserCmd)
}
