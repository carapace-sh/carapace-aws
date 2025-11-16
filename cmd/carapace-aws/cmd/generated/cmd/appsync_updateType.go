package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateTypeCmd = &cobra.Command{
	Use:   "update-type",
	Short: "Updates a `Type` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateTypeCmd).Standalone()

	appsync_updateTypeCmd.Flags().String("api-id", "", "The API ID.")
	appsync_updateTypeCmd.Flags().String("definition", "", "The new definition.")
	appsync_updateTypeCmd.Flags().String("format", "", "The new type format: SDL or JSON.")
	appsync_updateTypeCmd.Flags().String("type-name", "", "The new type name.")
	appsync_updateTypeCmd.MarkFlagRequired("api-id")
	appsync_updateTypeCmd.MarkFlagRequired("format")
	appsync_updateTypeCmd.MarkFlagRequired("type-name")
	appsyncCmd.AddCommand(appsync_updateTypeCmd)
}
