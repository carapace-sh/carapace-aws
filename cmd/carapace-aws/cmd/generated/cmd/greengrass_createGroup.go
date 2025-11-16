package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createGroupCmd).Standalone()

	greengrass_createGroupCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createGroupCmd.Flags().String("initial-version", "", "Information about the initial version of the group.")
	greengrass_createGroupCmd.Flags().String("name", "", "The name of the group.")
	greengrass_createGroupCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	greengrass_createGroupCmd.MarkFlagRequired("name")
	greengrassCmd.AddCommand(greengrass_createGroupCmd)
}
