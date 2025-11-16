package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_deleteGroupCmd).Standalone()

	synthetics_deleteGroupCmd.Flags().String("group-identifier", "", "Specifies which group to delete.")
	synthetics_deleteGroupCmd.MarkFlagRequired("group-identifier")
	syntheticsCmd.AddCommand(synthetics_deleteGroupCmd)
}
