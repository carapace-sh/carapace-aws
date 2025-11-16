package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_deleteHostCmd = &cobra.Command{
	Use:   "delete-host",
	Short: "The host to be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_deleteHostCmd).Standalone()

	codeconnections_deleteHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the host to be deleted.")
	codeconnections_deleteHostCmd.MarkFlagRequired("host-arn")
	codeconnectionsCmd.AddCommand(codeconnections_deleteHostCmd)
}
