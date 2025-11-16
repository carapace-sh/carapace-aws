package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_createResourceGroupCmd = &cobra.Command{
	Use:   "create-resource-group",
	Short: "Creates a resource group using the specified set of tags (key and value pairs) that are used to select the EC2 instances to be included in an Amazon Inspector assessment target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_createResourceGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_createResourceGroupCmd).Standalone()

		inspector_createResourceGroupCmd.Flags().String("resource-group-tags", "", "A collection of keys and an array of possible values, '\\[{\"key\":\"key1\",\"values\":\\[\"Value1\",\"Value2\"]},{\"key\":\"Key2\",\"values\":\\[\"Value3\"]}]'.")
		inspector_createResourceGroupCmd.MarkFlagRequired("resource-group-tags")
	})
	inspectorCmd.AddCommand(inspector_createResourceGroupCmd)
}
