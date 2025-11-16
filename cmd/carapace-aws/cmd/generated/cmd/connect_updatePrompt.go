package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updatePromptCmd = &cobra.Command{
	Use:   "update-prompt",
	Short: "Updates a prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updatePromptCmd).Standalone()

	connect_updatePromptCmd.Flags().String("description", "", "A description of the prompt.")
	connect_updatePromptCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updatePromptCmd.Flags().String("name", "", "The name of the prompt.")
	connect_updatePromptCmd.Flags().String("prompt-id", "", "A unique identifier for the prompt.")
	connect_updatePromptCmd.Flags().String("s3-uri", "", "The URI for the S3 bucket where the prompt is stored.")
	connect_updatePromptCmd.MarkFlagRequired("instance-id")
	connect_updatePromptCmd.MarkFlagRequired("prompt-id")
	connectCmd.AddCommand(connect_updatePromptCmd)
}
