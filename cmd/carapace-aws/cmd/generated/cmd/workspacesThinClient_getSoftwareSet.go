package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_getSoftwareSetCmd = &cobra.Command{
	Use:   "get-software-set",
	Short: "Returns information for a software set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_getSoftwareSetCmd).Standalone()

	workspacesThinClient_getSoftwareSetCmd.Flags().String("id", "", "The ID of the software set for which to return information.")
	workspacesThinClient_getSoftwareSetCmd.MarkFlagRequired("id")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_getSoftwareSetCmd)
}
