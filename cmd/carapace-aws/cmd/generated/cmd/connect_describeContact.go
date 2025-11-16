package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeContactCmd = &cobra.Command{
	Use:   "describe-contact",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeContactCmd).Standalone()

	connect_describeContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_describeContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeContactCmd.MarkFlagRequired("contact-id")
	connect_describeContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeContactCmd)
}
