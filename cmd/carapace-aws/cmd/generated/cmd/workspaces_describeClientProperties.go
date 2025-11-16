package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeClientPropertiesCmd = &cobra.Command{
	Use:   "describe-client-properties",
	Short: "Retrieves a list that describes one or more specified Amazon WorkSpaces clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeClientPropertiesCmd).Standalone()

	workspaces_describeClientPropertiesCmd.Flags().String("resource-ids", "", "The resource identifier, in the form of directory IDs.")
	workspaces_describeClientPropertiesCmd.MarkFlagRequired("resource-ids")
	workspacesCmd.AddCommand(workspaces_describeClientPropertiesCmd)
}
