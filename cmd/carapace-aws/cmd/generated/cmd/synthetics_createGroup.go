package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a group which you can use to associate canaries with each other, including cross-Region canaries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_createGroupCmd).Standalone()

	synthetics_createGroupCmd.Flags().String("name", "", "The name for the group.")
	synthetics_createGroupCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the group.")
	synthetics_createGroupCmd.MarkFlagRequired("name")
	syntheticsCmd.AddCommand(synthetics_createGroupCmd)
}
