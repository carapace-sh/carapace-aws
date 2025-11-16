package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteTypeCmd = &cobra.Command{
	Use:   "delete-type",
	Short: "Deletes a `Type` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteTypeCmd).Standalone()

	appsync_deleteTypeCmd.Flags().String("api-id", "", "The API ID.")
	appsync_deleteTypeCmd.Flags().String("type-name", "", "The type name.")
	appsync_deleteTypeCmd.MarkFlagRequired("api-id")
	appsync_deleteTypeCmd.MarkFlagRequired("type-name")
	appsyncCmd.AddCommand(appsync_deleteTypeCmd)
}
