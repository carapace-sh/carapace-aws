package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_deleteSpaceCmd = &cobra.Command{
	Use:   "delete-space",
	Short: "Deletes a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_deleteSpaceCmd).Standalone()

	codecatalyst_deleteSpaceCmd.Flags().String("name", "", "The name of the space.")
	codecatalyst_deleteSpaceCmd.MarkFlagRequired("name")
	codecatalystCmd.AddCommand(codecatalyst_deleteSpaceCmd)
}
