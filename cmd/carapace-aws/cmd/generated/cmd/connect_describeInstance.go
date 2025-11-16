package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeInstanceCmd = &cobra.Command{
	Use:   "describe-instance",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeInstanceCmd).Standalone()

	connect_describeInstanceCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeInstanceCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeInstanceCmd)
}
