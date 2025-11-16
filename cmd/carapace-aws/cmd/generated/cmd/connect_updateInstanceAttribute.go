package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateInstanceAttributeCmd = &cobra.Command{
	Use:   "update-instance-attribute",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateInstanceAttributeCmd).Standalone()

	connect_updateInstanceAttributeCmd.Flags().String("attribute-type", "", "The type of attribute.")
	connect_updateInstanceAttributeCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_updateInstanceAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateInstanceAttributeCmd.Flags().String("value", "", "The value for the attribute.")
	connect_updateInstanceAttributeCmd.MarkFlagRequired("attribute-type")
	connect_updateInstanceAttributeCmd.MarkFlagRequired("instance-id")
	connect_updateInstanceAttributeCmd.MarkFlagRequired("value")
	connectCmd.AddCommand(connect_updateInstanceAttributeCmd)
}
