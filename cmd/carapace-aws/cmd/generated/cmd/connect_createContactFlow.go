package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createContactFlowCmd = &cobra.Command{
	Use:   "create-contact-flow",
	Short: "Creates a flow for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createContactFlowCmd).Standalone()

	connect_createContactFlowCmd.Flags().String("content", "", "The JSON string that represents the content of the flow.")
	connect_createContactFlowCmd.Flags().String("description", "", "The description of the flow.")
	connect_createContactFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createContactFlowCmd.Flags().String("name", "", "The name of the flow.")
	connect_createContactFlowCmd.Flags().String("status", "", "Indicates the flow status as either `SAVED` or `PUBLISHED`.")
	connect_createContactFlowCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createContactFlowCmd.Flags().String("type", "", "The type of the flow.")
	connect_createContactFlowCmd.MarkFlagRequired("content")
	connect_createContactFlowCmd.MarkFlagRequired("instance-id")
	connect_createContactFlowCmd.MarkFlagRequired("name")
	connect_createContactFlowCmd.MarkFlagRequired("type")
	connectCmd.AddCommand(connect_createContactFlowCmd)
}
