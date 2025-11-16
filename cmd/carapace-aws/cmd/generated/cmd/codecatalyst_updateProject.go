package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Changes one or more values for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_updateProjectCmd).Standalone()

		codecatalyst_updateProjectCmd.Flags().String("description", "", "The description of the project.")
		codecatalyst_updateProjectCmd.Flags().String("name", "", "The name of the project.")
		codecatalyst_updateProjectCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_updateProjectCmd.MarkFlagRequired("name")
		codecatalyst_updateProjectCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_updateProjectCmd)
}
