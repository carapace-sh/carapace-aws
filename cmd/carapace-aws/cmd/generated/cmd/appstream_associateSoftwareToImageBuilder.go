package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_associateSoftwareToImageBuilderCmd = &cobra.Command{
	Use:   "associate-software-to-image-builder",
	Short: "Associates license included application(s) with an existing image builder instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_associateSoftwareToImageBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_associateSoftwareToImageBuilderCmd).Standalone()

		appstream_associateSoftwareToImageBuilderCmd.Flags().String("image-builder-name", "", "The name of the target image builder instance.")
		appstream_associateSoftwareToImageBuilderCmd.Flags().String("software-names", "", "The list of license included applications to associate with the image builder.")
		appstream_associateSoftwareToImageBuilderCmd.MarkFlagRequired("image-builder-name")
		appstream_associateSoftwareToImageBuilderCmd.MarkFlagRequired("software-names")
	})
	appstreamCmd.AddCommand(appstream_associateSoftwareToImageBuilderCmd)
}
