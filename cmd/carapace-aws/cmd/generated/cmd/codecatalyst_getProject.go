package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getProjectCmd = &cobra.Command{
	Use:   "get-project",
	Short: "Returns information about a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getProjectCmd).Standalone()

	codecatalyst_getProjectCmd.Flags().String("name", "", "The name of the project in the space.")
	codecatalyst_getProjectCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_getProjectCmd.MarkFlagRequired("name")
	codecatalyst_getProjectCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_getProjectCmd)
}
