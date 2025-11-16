package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeViewCmd = &cobra.Command{
	Use:   "describe-view",
	Short: "Retrieves the view for the specified Amazon Connect instance and view identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeViewCmd).Standalone()

	connect_describeViewCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeViewCmd.Flags().String("view-id", "", "The ViewId of the view.")
	connect_describeViewCmd.MarkFlagRequired("instance-id")
	connect_describeViewCmd.MarkFlagRequired("view-id")
	connectCmd.AddCommand(connect_describeViewCmd)
}
