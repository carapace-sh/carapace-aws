package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createPromptCmd = &cobra.Command{
	Use:   "create-prompt",
	Short: "Creates a prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createPromptCmd).Standalone()

	connect_createPromptCmd.Flags().String("description", "", "The description of the prompt.")
	connect_createPromptCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createPromptCmd.Flags().String("name", "", "The name of the prompt.")
	connect_createPromptCmd.Flags().String("s3-uri", "", "The URI for the S3 bucket where the prompt is stored.")
	connect_createPromptCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createPromptCmd.MarkFlagRequired("instance-id")
	connect_createPromptCmd.MarkFlagRequired("name")
	connect_createPromptCmd.MarkFlagRequired("s3-uri")
	connectCmd.AddCommand(connect_createPromptCmd)
}
