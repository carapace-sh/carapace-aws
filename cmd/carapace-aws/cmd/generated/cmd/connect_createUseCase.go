package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createUseCaseCmd = &cobra.Command{
	Use:   "create-use-case",
	Short: "Creates a use case for an integration association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createUseCaseCmd).Standalone()

	connect_createUseCaseCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createUseCaseCmd.Flags().String("integration-association-id", "", "The identifier for the integration association.")
	connect_createUseCaseCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createUseCaseCmd.Flags().String("use-case-type", "", "The type of use case to associate to the integration association.")
	connect_createUseCaseCmd.MarkFlagRequired("instance-id")
	connect_createUseCaseCmd.MarkFlagRequired("integration-association-id")
	connect_createUseCaseCmd.MarkFlagRequired("use-case-type")
	connectCmd.AddCommand(connect_createUseCaseCmd)
}
