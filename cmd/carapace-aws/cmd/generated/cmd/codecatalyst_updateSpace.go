package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_updateSpaceCmd = &cobra.Command{
	Use:   "update-space",
	Short: "Changes one or more values for a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_updateSpaceCmd).Standalone()

	codecatalyst_updateSpaceCmd.Flags().String("description", "", "The description of the space.")
	codecatalyst_updateSpaceCmd.Flags().String("name", "", "The name of the space.")
	codecatalyst_updateSpaceCmd.MarkFlagRequired("name")
	codecatalystCmd.AddCommand(codecatalyst_updateSpaceCmd)
}
