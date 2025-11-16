package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disassociateAppBlockBuilderAppBlockCmd = &cobra.Command{
	Use:   "disassociate-app-block-builder-app-block",
	Short: "Disassociates a specified app block builder from a specified app block.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disassociateAppBlockBuilderAppBlockCmd).Standalone()

	appstream_disassociateAppBlockBuilderAppBlockCmd.Flags().String("app-block-arn", "", "The ARN of the app block.")
	appstream_disassociateAppBlockBuilderAppBlockCmd.Flags().String("app-block-builder-name", "", "The name of the app block builder.")
	appstream_disassociateAppBlockBuilderAppBlockCmd.MarkFlagRequired("app-block-arn")
	appstream_disassociateAppBlockBuilderAppBlockCmd.MarkFlagRequired("app-block-builder-name")
	appstreamCmd.AddCommand(appstream_disassociateAppBlockBuilderAppBlockCmd)
}
