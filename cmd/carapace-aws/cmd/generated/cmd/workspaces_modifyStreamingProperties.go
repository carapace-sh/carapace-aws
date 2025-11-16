package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyStreamingPropertiesCmd = &cobra.Command{
	Use:   "modify-streaming-properties",
	Short: "Modifies the specified streaming properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyStreamingPropertiesCmd).Standalone()

	workspaces_modifyStreamingPropertiesCmd.Flags().String("resource-id", "", "The identifier of the resource.")
	workspaces_modifyStreamingPropertiesCmd.Flags().String("streaming-properties", "", "The streaming properties to configure.")
	workspaces_modifyStreamingPropertiesCmd.MarkFlagRequired("resource-id")
	workspacesCmd.AddCommand(workspaces_modifyStreamingPropertiesCmd)
}
