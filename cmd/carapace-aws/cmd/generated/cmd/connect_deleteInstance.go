package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteInstanceCmd = &cobra.Command{
	Use:   "delete-instance",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteInstanceCmd).Standalone()

		connect_deleteInstanceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_deleteInstanceCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteInstanceCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteInstanceCmd)
}
