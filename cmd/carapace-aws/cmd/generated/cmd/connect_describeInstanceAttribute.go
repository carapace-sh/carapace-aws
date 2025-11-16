package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeInstanceAttributeCmd = &cobra.Command{
	Use:   "describe-instance-attribute",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeInstanceAttributeCmd).Standalone()

	connect_describeInstanceAttributeCmd.Flags().String("attribute-type", "", "The type of attribute.")
	connect_describeInstanceAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeInstanceAttributeCmd.MarkFlagRequired("attribute-type")
	connect_describeInstanceAttributeCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeInstanceAttributeCmd)
}
