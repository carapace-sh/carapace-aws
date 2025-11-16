package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updatePredefinedAttributeCmd = &cobra.Command{
	Use:   "update-predefined-attribute",
	Short: "Updates a predefined attribute for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updatePredefinedAttributeCmd).Standalone()

	connect_updatePredefinedAttributeCmd.Flags().String("attribute-configuration", "", "Custom metadata that is associated to predefined attributes to control behavior in upstream services, such as controlling how a predefined attribute should be displayed in the Amazon Connect admin website.")
	connect_updatePredefinedAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updatePredefinedAttributeCmd.Flags().String("name", "", "The name of the predefined attribute.")
	connect_updatePredefinedAttributeCmd.Flags().String("purposes", "", "Values that enable you to categorize your predefined attributes.")
	connect_updatePredefinedAttributeCmd.Flags().String("values", "", "The values of the predefined attribute.")
	connect_updatePredefinedAttributeCmd.MarkFlagRequired("instance-id")
	connect_updatePredefinedAttributeCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_updatePredefinedAttributeCmd)
}
