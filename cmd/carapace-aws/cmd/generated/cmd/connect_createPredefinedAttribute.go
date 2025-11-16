package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createPredefinedAttributeCmd = &cobra.Command{
	Use:   "create-predefined-attribute",
	Short: "Creates a new predefined attribute for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createPredefinedAttributeCmd).Standalone()

	connect_createPredefinedAttributeCmd.Flags().String("attribute-configuration", "", "Custom metadata that is associated to predefined attributes to control behavior in upstream services, such as controlling how a predefined attribute should be displayed in the Amazon Connect admin website.")
	connect_createPredefinedAttributeCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createPredefinedAttributeCmd.Flags().String("name", "", "The name of the predefined attribute.")
	connect_createPredefinedAttributeCmd.Flags().String("purposes", "", "Values that enable you to categorize your predefined attributes.")
	connect_createPredefinedAttributeCmd.Flags().String("values", "", "The values of the predefined attribute.")
	connect_createPredefinedAttributeCmd.MarkFlagRequired("instance-id")
	connect_createPredefinedAttributeCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_createPredefinedAttributeCmd)
}
