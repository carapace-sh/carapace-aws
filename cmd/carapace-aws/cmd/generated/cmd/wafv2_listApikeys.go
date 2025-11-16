package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listApikeysCmd = &cobra.Command{
	Use:   "list-apikeys",
	Short: "Retrieves a list of the API keys that you've defined for the specified scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listApikeysCmd).Standalone()

	wafv2_listApikeysCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listApikeysCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listApikeysCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_listApikeysCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_listApikeysCmd)
}
