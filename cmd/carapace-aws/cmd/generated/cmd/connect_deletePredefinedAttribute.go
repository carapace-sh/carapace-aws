package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deletePredefinedAttributeCmd = &cobra.Command{
	Use:   "delete-predefined-attribute",
	Short: "Deletes a predefined attribute from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deletePredefinedAttributeCmd).Standalone()

	connect_deletePredefinedAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deletePredefinedAttributeCmd.Flags().String("name", "", "The name of the predefined attribute.")
	connect_deletePredefinedAttributeCmd.MarkFlagRequired("instance-id")
	connect_deletePredefinedAttributeCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_deletePredefinedAttributeCmd)
}
