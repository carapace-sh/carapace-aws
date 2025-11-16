package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getMapSpritesCmd = &cobra.Command{
	Use:   "get-map-sprites",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getMapSpritesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getMapSpritesCmd).Standalone()

		location_getMapSpritesCmd.Flags().String("file-name", "", "The name of the sprite ﬁle.")
		location_getMapSpritesCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
		location_getMapSpritesCmd.Flags().String("map-name", "", "The map resource associated with the sprite ﬁle.")
		location_getMapSpritesCmd.MarkFlagRequired("file-name")
		location_getMapSpritesCmd.MarkFlagRequired("map-name")
	})
	locationCmd.AddCommand(location_getMapSpritesCmd)
}
