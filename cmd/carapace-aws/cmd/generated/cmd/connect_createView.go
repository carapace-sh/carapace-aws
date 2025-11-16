package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createViewCmd = &cobra.Command{
	Use:   "create-view",
	Short: "Creates a new view with the possible status of `SAVED` or `PUBLISHED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createViewCmd).Standalone()

		connect_createViewCmd.Flags().String("client-token", "", "A unique Id for each create view request to avoid duplicate view creation.")
		connect_createViewCmd.Flags().String("content", "", "View content containing all content necessary to render a view except for runtime input data.")
		connect_createViewCmd.Flags().String("description", "", "The description of the view.")
		connect_createViewCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createViewCmd.Flags().String("name", "", "The name of the view.")
		connect_createViewCmd.Flags().String("status", "", "Indicates the view status as either `SAVED` or `PUBLISHED`.")
		connect_createViewCmd.Flags().String("tags", "", "The tags associated with the view resource (not specific to view version).These tags can be used to organize, track, or control access for this resource.")
		connect_createViewCmd.MarkFlagRequired("content")
		connect_createViewCmd.MarkFlagRequired("instance-id")
		connect_createViewCmd.MarkFlagRequired("name")
		connect_createViewCmd.MarkFlagRequired("status")
	})
	connectCmd.AddCommand(connect_createViewCmd)
}
