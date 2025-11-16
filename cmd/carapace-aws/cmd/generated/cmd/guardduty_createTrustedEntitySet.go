package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createTrustedEntitySetCmd = &cobra.Command{
	Use:   "create-trusted-entity-set",
	Short: "Creates a new trusted entity set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createTrustedEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_createTrustedEntitySetCmd).Standalone()

		guardduty_createTrustedEntitySetCmd.Flags().Bool("activate", false, "A boolean value that indicates whether GuardDuty is to start using the uploaded trusted entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().String("client-token", "", "The idempotency token for the create request.")
		guardduty_createTrustedEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account for which you want to create a trusted entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
		guardduty_createTrustedEntitySetCmd.Flags().String("format", "", "The format of the file that contains the trusted entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().String("location", "", "The URI of the file that contains the threat entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().String("name", "", "A user-friendly name to identify the trusted entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().Bool("no-activate", false, "A boolean value that indicates whether GuardDuty is to start using the uploaded trusted entity set.")
		guardduty_createTrustedEntitySetCmd.Flags().String("tags", "", "The tags to be added to a new trusted entity set resource.")
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("activate")
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("format")
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("location")
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("name")
		guardduty_createTrustedEntitySetCmd.Flag("no-activate").Hidden = true
		guardduty_createTrustedEntitySetCmd.MarkFlagRequired("no-activate")
	})
	guarddutyCmd.AddCommand(guardduty_createTrustedEntitySetCmd)
}
