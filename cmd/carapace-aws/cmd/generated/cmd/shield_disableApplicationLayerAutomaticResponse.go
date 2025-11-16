package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_disableApplicationLayerAutomaticResponseCmd = &cobra.Command{
	Use:   "disable-application-layer-automatic-response",
	Short: "Disable the Shield Advanced automatic application layer DDoS mitigation feature for the protected resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_disableApplicationLayerAutomaticResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_disableApplicationLayerAutomaticResponseCmd).Standalone()

		shield_disableApplicationLayerAutomaticResponseCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the protected resource.")
		shield_disableApplicationLayerAutomaticResponseCmd.MarkFlagRequired("resource-arn")
	})
	shieldCmd.AddCommand(shield_disableApplicationLayerAutomaticResponseCmd)
}
