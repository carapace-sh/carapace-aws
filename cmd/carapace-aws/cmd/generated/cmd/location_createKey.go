package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createKeyCmd = &cobra.Command{
	Use:   "create-key",
	Short: "Creates an API key resource in your Amazon Web Services account, which lets you grant actions for Amazon Location resources to the API key bearer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_createKeyCmd).Standalone()

		location_createKeyCmd.Flags().String("description", "", "An optional description for the API key resource.")
		location_createKeyCmd.Flags().String("expire-time", "", "The optional timestamp for when the API key resource will expire in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ`.")
		location_createKeyCmd.Flags().String("key-name", "", "A custom name for the API key resource.")
		location_createKeyCmd.Flags().Bool("no-expiry", false, "Optionally set to `true` to set no expiration time for the API key.")
		location_createKeyCmd.Flags().Bool("no-no-expiry", false, "Optionally set to `true` to set no expiration time for the API key.")
		location_createKeyCmd.Flags().String("restrictions", "", "The API key restrictions for the API key resource.")
		location_createKeyCmd.Flags().String("tags", "", "Applies one or more tags to the map resource.")
		location_createKeyCmd.MarkFlagRequired("key-name")
		location_createKeyCmd.Flag("no-no-expiry").Hidden = true
		location_createKeyCmd.MarkFlagRequired("restrictions")
	})
	locationCmd.AddCommand(location_createKeyCmd)
}
