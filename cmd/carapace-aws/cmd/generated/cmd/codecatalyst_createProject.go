package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a project in a specified space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_createProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_createProjectCmd).Standalone()

		codecatalyst_createProjectCmd.Flags().String("description", "", "The description of the project.")
		codecatalyst_createProjectCmd.Flags().String("display-name", "", "The friendly name of the project that will be displayed to users.")
		codecatalyst_createProjectCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_createProjectCmd.MarkFlagRequired("display-name")
		codecatalyst_createProjectCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_createProjectCmd)
}
