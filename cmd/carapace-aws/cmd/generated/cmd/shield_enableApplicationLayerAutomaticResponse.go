package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_enableApplicationLayerAutomaticResponseCmd = &cobra.Command{
	Use:   "enable-application-layer-automatic-response",
	Short: "Enable the Shield Advanced automatic application layer DDoS mitigation for the protected resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_enableApplicationLayerAutomaticResponseCmd).Standalone()

	shield_enableApplicationLayerAutomaticResponseCmd.Flags().String("action", "", "Specifies the action setting that Shield Advanced should use in the WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.")
	shield_enableApplicationLayerAutomaticResponseCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the protected resource.")
	shield_enableApplicationLayerAutomaticResponseCmd.MarkFlagRequired("action")
	shield_enableApplicationLayerAutomaticResponseCmd.MarkFlagRequired("resource-arn")
	shieldCmd.AddCommand(shield_enableApplicationLayerAutomaticResponseCmd)
}
