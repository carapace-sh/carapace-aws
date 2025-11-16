package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteProjectMembershipCmd = &cobra.Command{
	Use:   "delete-project-membership",
	Short: "Deletes project membership in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteProjectMembershipCmd).Standalone()

	datazone_deleteProjectMembershipCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where project membership is deleted.")
	datazone_deleteProjectMembershipCmd.Flags().String("member", "", "The project member whose project membership is deleted.")
	datazone_deleteProjectMembershipCmd.Flags().String("project-identifier", "", "The ID of the Amazon DataZone project the membership to which is deleted.")
	datazone_deleteProjectMembershipCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteProjectMembershipCmd.MarkFlagRequired("member")
	datazone_deleteProjectMembershipCmd.MarkFlagRequired("project-identifier")
	datazoneCmd.AddCommand(datazone_deleteProjectMembershipCmd)
}
