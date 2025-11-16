package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_disassociateIpGroupsCmd = &cobra.Command{
	Use:   "disassociate-ip-groups",
	Short: "Disassociates the specified IP access control group from the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_disassociateIpGroupsCmd).Standalone()

	workspaces_disassociateIpGroupsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	workspaces_disassociateIpGroupsCmd.Flags().String("group-ids", "", "The identifiers of one or more IP access control groups.")
	workspaces_disassociateIpGroupsCmd.MarkFlagRequired("directory-id")
	workspaces_disassociateIpGroupsCmd.MarkFlagRequired("group-ids")
	workspacesCmd.AddCommand(workspaces_disassociateIpGroupsCmd)
}
