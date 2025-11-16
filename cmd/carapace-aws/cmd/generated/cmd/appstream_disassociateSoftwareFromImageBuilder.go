package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disassociateSoftwareFromImageBuilderCmd = &cobra.Command{
	Use:   "disassociate-software-from-image-builder",
	Short: "Removes license included application(s) association(s) from an image builder instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disassociateSoftwareFromImageBuilderCmd).Standalone()

	appstream_disassociateSoftwareFromImageBuilderCmd.Flags().String("image-builder-name", "", "The name of the target image builder instance.")
	appstream_disassociateSoftwareFromImageBuilderCmd.Flags().String("software-names", "", "The list of license included applications to disassociate from the image builder.")
	appstream_disassociateSoftwareFromImageBuilderCmd.MarkFlagRequired("image-builder-name")
	appstream_disassociateSoftwareFromImageBuilderCmd.MarkFlagRequired("software-names")
	appstreamCmd.AddCommand(appstream_disassociateSoftwareFromImageBuilderCmd)
}
