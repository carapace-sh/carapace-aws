package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_createOrganizationalUnitCmd = &cobra.Command{
	Use:   "create-organizational-unit",
	Short: "Creates an organizational unit (OU) within a root or parent OU.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createOrganizationalUnitCmd).Standalone()

	organizations_createOrganizationalUnitCmd.Flags().String("name", "", "The friendly name to assign to the new OU.")
	organizations_createOrganizationalUnitCmd.Flags().String("parent-id", "", "The unique identifier (ID) of the parent root or OU that you want to create the new OU in.")
	organizations_createOrganizationalUnitCmd.Flags().String("tags", "", "A list of tags that you want to attach to the newly created OU.")
	organizations_createOrganizationalUnitCmd.MarkFlagRequired("name")
	organizations_createOrganizationalUnitCmd.MarkFlagRequired("parent-id")
	organizationsCmd.AddCommand(organizations_createOrganizationalUnitCmd)
}
