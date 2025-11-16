package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getSpaceCmd = &cobra.Command{
	Use:   "get-space",
	Short: "Returns information about an space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getSpaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_getSpaceCmd).Standalone()

		codecatalyst_getSpaceCmd.Flags().String("name", "", "The name of the space.")
		codecatalyst_getSpaceCmd.MarkFlagRequired("name")
	})
	codecatalystCmd.AddCommand(codecatalyst_getSpaceCmd)
}
