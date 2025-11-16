package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describePredefinedAttributeCmd = &cobra.Command{
	Use:   "describe-predefined-attribute",
	Short: "Describes a predefined attribute for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describePredefinedAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describePredefinedAttributeCmd).Standalone()

		connect_describePredefinedAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describePredefinedAttributeCmd.Flags().String("name", "", "The name of the predefined attribute.")
		connect_describePredefinedAttributeCmd.MarkFlagRequired("instance-id")
		connect_describePredefinedAttributeCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_describePredefinedAttributeCmd)
}
