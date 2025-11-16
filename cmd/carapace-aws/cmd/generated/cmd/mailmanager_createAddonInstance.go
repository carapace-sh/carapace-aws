package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createAddonInstanceCmd = &cobra.Command{
	Use:   "create-addon-instance",
	Short: "Creates an Add On instance for the subscription indicated in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createAddonInstanceCmd).Standalone()

	mailmanager_createAddonInstanceCmd.Flags().String("addon-subscription-id", "", "The unique ID of a previously created subscription that an Add On instance is created for.")
	mailmanager_createAddonInstanceCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createAddonInstanceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createAddonInstanceCmd.MarkFlagRequired("addon-subscription-id")
	mailmanagerCmd.AddCommand(mailmanager_createAddonInstanceCmd)
}
