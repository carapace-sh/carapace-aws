package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createContactFlowModuleCmd = &cobra.Command{
	Use:   "create-contact-flow-module",
	Short: "Creates a flow module for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createContactFlowModuleCmd).Standalone()

	connect_createContactFlowModuleCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createContactFlowModuleCmd.Flags().String("content", "", "The JSON string that represents the content of the flow.")
	connect_createContactFlowModuleCmd.Flags().String("description", "", "The description of the flow module.")
	connect_createContactFlowModuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createContactFlowModuleCmd.Flags().String("name", "", "The name of the flow module.")
	connect_createContactFlowModuleCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createContactFlowModuleCmd.MarkFlagRequired("content")
	connect_createContactFlowModuleCmd.MarkFlagRequired("instance-id")
	connect_createContactFlowModuleCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_createContactFlowModuleCmd)
}
