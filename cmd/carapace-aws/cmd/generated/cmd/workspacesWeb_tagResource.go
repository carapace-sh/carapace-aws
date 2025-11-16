package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites one or more tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_tagResourceCmd).Standalone()

	workspacesWeb_tagResourceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	workspacesWeb_tagResourceCmd.Flags().String("tags", "", "The tags of the resource.")
	workspacesWeb_tagResourceCmd.MarkFlagRequired("resource-arn")
	workspacesWeb_tagResourceCmd.MarkFlagRequired("tags")
	workspacesWebCmd.AddCommand(workspacesWeb_tagResourceCmd)
}
