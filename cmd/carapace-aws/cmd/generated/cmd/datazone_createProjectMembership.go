package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createProjectMembershipCmd = &cobra.Command{
	Use:   "create-project-membership",
	Short: "Creates a project membership in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createProjectMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createProjectMembershipCmd).Standalone()

		datazone_createProjectMembershipCmd.Flags().String("designation", "", "The designation of the project membership.")
		datazone_createProjectMembershipCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which project membership is created.")
		datazone_createProjectMembershipCmd.Flags().String("member", "", "The project member whose project membership was created.")
		datazone_createProjectMembershipCmd.Flags().String("project-identifier", "", "The ID of the project for which this project membership was created.")
		datazone_createProjectMembershipCmd.MarkFlagRequired("designation")
		datazone_createProjectMembershipCmd.MarkFlagRequired("domain-identifier")
		datazone_createProjectMembershipCmd.MarkFlagRequired("member")
		datazone_createProjectMembershipCmd.MarkFlagRequired("project-identifier")
	})
	datazoneCmd.AddCommand(datazone_createProjectMembershipCmd)
}
