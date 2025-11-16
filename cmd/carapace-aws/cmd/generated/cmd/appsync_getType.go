package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getTypeCmd = &cobra.Command{
	Use:   "get-type",
	Short: "Retrieves a `Type` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getTypeCmd).Standalone()

		appsync_getTypeCmd.Flags().String("api-id", "", "The API ID.")
		appsync_getTypeCmd.Flags().String("format", "", "The type format: SDL or JSON.")
		appsync_getTypeCmd.Flags().String("type-name", "", "The type name.")
		appsync_getTypeCmd.MarkFlagRequired("api-id")
		appsync_getTypeCmd.MarkFlagRequired("format")
		appsync_getTypeCmd.MarkFlagRequired("type-name")
	})
	appsyncCmd.AddCommand(appsync_getTypeCmd)
}
