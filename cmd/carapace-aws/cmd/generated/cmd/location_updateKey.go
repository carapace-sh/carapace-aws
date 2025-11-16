package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updateKeyCmd = &cobra.Command{
	Use:   "update-key",
	Short: "Updates the specified properties of a given API key resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updateKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_updateKeyCmd).Standalone()

		location_updateKeyCmd.Flags().String("description", "", "Updates the description for the API key resource.")
		location_updateKeyCmd.Flags().String("expire-time", "", "Updates the timestamp for when the API key resource will expire in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ`.")
		location_updateKeyCmd.Flags().Bool("force-update", false, "The boolean flag to be included for updating `ExpireTime` or `Restrictions` details.")
		location_updateKeyCmd.Flags().String("key-name", "", "The name of the API key resource to update.")
		location_updateKeyCmd.Flags().Bool("no-expiry", false, "Whether the API key should expire.")
		location_updateKeyCmd.Flags().Bool("no-force-update", false, "The boolean flag to be included for updating `ExpireTime` or `Restrictions` details.")
		location_updateKeyCmd.Flags().Bool("no-no-expiry", false, "Whether the API key should expire.")
		location_updateKeyCmd.Flags().String("restrictions", "", "Updates the API key restrictions for the API key resource.")
		location_updateKeyCmd.MarkFlagRequired("key-name")
		location_updateKeyCmd.Flag("no-force-update").Hidden = true
		location_updateKeyCmd.Flag("no-no-expiry").Hidden = true
	})
	locationCmd.AddCommand(location_updateKeyCmd)
}
