package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteApiCmd = &cobra.Command{
	Use:   "delete-api",
	Short: "Deletes an `Api` object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteApiCmd).Standalone()

	appsync_deleteApiCmd.Flags().String("api-id", "", "The `Api` ID.")
	appsync_deleteApiCmd.MarkFlagRequired("api-id")
	appsyncCmd.AddCommand(appsync_deleteApiCmd)
}
