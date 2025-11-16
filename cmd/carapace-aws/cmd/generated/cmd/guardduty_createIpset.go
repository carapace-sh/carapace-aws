package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createIpsetCmd = &cobra.Command{
	Use:   "create-ipset",
	Short: "Creates a new IPSet, which is called a trusted IP list in the console user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createIpsetCmd).Standalone()

	guardduty_createIpsetCmd.Flags().Bool("activate", false, "A Boolean value that indicates whether GuardDuty is to start using the uploaded IPSet.")
	guardduty_createIpsetCmd.Flags().String("client-token", "", "The idempotency token for the create request.")
	guardduty_createIpsetCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account for which you want to create an IPSet.")
	guardduty_createIpsetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
	guardduty_createIpsetCmd.Flags().String("format", "", "The format of the file that contains the IPSet.")
	guardduty_createIpsetCmd.Flags().String("location", "", "The URI of the file that contains the IPSet.")
	guardduty_createIpsetCmd.Flags().String("name", "", "The user-friendly name to identify the IPSet.")
	guardduty_createIpsetCmd.Flags().Bool("no-activate", false, "A Boolean value that indicates whether GuardDuty is to start using the uploaded IPSet.")
	guardduty_createIpsetCmd.Flags().String("tags", "", "The tags to be added to a new IP set resource.")
	guardduty_createIpsetCmd.MarkFlagRequired("activate")
	guardduty_createIpsetCmd.MarkFlagRequired("detector-id")
	guardduty_createIpsetCmd.MarkFlagRequired("format")
	guardduty_createIpsetCmd.MarkFlagRequired("location")
	guardduty_createIpsetCmd.MarkFlagRequired("name")
	guardduty_createIpsetCmd.Flag("no-activate").Hidden = true
	guardduty_createIpsetCmd.MarkFlagRequired("no-activate")
	guarddutyCmd.AddCommand(guardduty_createIpsetCmd)
}
