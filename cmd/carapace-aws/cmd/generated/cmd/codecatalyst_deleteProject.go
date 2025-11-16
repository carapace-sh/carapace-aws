package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes a project in a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_deleteProjectCmd).Standalone()

		codecatalyst_deleteProjectCmd.Flags().String("name", "", "The name of the project in the space.")
		codecatalyst_deleteProjectCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_deleteProjectCmd.MarkFlagRequired("name")
		codecatalyst_deleteProjectCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_deleteProjectCmd)
}
