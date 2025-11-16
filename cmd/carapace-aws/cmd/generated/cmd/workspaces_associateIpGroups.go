package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_associateIpGroupsCmd = &cobra.Command{
	Use:   "associate-ip-groups",
	Short: "Associates the specified IP access control group with the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_associateIpGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_associateIpGroupsCmd).Standalone()

		workspaces_associateIpGroupsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		workspaces_associateIpGroupsCmd.Flags().String("group-ids", "", "The identifiers of one or more IP access control groups.")
		workspaces_associateIpGroupsCmd.MarkFlagRequired("directory-id")
		workspaces_associateIpGroupsCmd.MarkFlagRequired("group-ids")
	})
	workspacesCmd.AddCommand(workspaces_associateIpGroupsCmd)
}
