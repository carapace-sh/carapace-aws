package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes a project in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteProjectCmd).Standalone()

		datazone_deleteProjectCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the project is deleted.")
		datazone_deleteProjectCmd.Flags().String("identifier", "", "The identifier of the project that is to be deleted.")
		datazone_deleteProjectCmd.Flags().Bool("no-skip-deletion-check", false, "Specifies the optional flag to delete all child entities within the project.")
		datazone_deleteProjectCmd.Flags().Bool("skip-deletion-check", false, "Specifies the optional flag to delete all child entities within the project.")
		datazone_deleteProjectCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteProjectCmd.MarkFlagRequired("identifier")
		datazone_deleteProjectCmd.Flag("no-skip-deletion-check").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_deleteProjectCmd)
}
