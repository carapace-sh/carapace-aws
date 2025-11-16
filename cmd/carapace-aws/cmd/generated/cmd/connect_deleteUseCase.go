package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteUseCaseCmd = &cobra.Command{
	Use:   "delete-use-case",
	Short: "Deletes a use case from an integration association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteUseCaseCmd).Standalone()

	connect_deleteUseCaseCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteUseCaseCmd.Flags().String("integration-association-id", "", "The identifier for the integration association.")
	connect_deleteUseCaseCmd.Flags().String("use-case-id", "", "The identifier for the use case.")
	connect_deleteUseCaseCmd.MarkFlagRequired("instance-id")
	connect_deleteUseCaseCmd.MarkFlagRequired("integration-association-id")
	connect_deleteUseCaseCmd.MarkFlagRequired("use-case-id")
	connectCmd.AddCommand(connect_deleteUseCaseCmd)
}
