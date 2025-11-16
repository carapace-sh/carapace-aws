package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Describes the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeUserCmd).Standalone()

	connect_describeUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeUserCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_describeUserCmd.MarkFlagRequired("instance-id")
	connect_describeUserCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_describeUserCmd)
}
