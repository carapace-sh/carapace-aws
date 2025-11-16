package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_deleteLocationCmd = &cobra.Command{
	Use:   "delete-location",
	Short: "Deletes a transfer location resource from DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_deleteLocationCmd).Standalone()

	datasync_deleteLocationCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the location to delete.")
	datasync_deleteLocationCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_deleteLocationCmd)
}
