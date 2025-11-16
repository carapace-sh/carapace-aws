package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_createCaseCmd = &cobra.Command{
	Use:   "create-case",
	Short: "Creates a new case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_createCaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_createCaseCmd).Standalone()

		securityIr_createCaseCmd.Flags().String("client-token", "", "The `clientToken` field is an idempotency key used to ensure that repeated attempts for a single action will be ignored by the server during retries.")
		securityIr_createCaseCmd.Flags().String("description", "", "Required element used in combination with CreateCase")
		securityIr_createCaseCmd.Flags().String("engagement-type", "", "Required element used in combination with CreateCase to provide an engagement type for the new cases.")
		securityIr_createCaseCmd.Flags().String("impacted-accounts", "", "Required element used in combination with CreateCase to provide a list of impacted accounts.")
		securityIr_createCaseCmd.Flags().String("impacted-aws-regions", "", "An optional element used in combination with CreateCase to provide a list of impacted regions.")
		securityIr_createCaseCmd.Flags().String("impacted-services", "", "An optional element used in combination with CreateCase to provide a list of services impacted.")
		securityIr_createCaseCmd.Flags().String("reported-incident-start-date", "", "Required element used in combination with CreateCase to provide an initial start date for the unauthorized activity.")
		securityIr_createCaseCmd.Flags().String("resolver-type", "", "Required element used in combination with CreateCase to identify the resolver type.")
		securityIr_createCaseCmd.Flags().String("tags", "", "An optional element used in combination with CreateCase to add customer specified tags to a case.")
		securityIr_createCaseCmd.Flags().String("threat-actor-ip-addresses", "", "An optional element used in combination with CreateCase to provide a list of suspicious internet protocol addresses associated with unauthorized activity.")
		securityIr_createCaseCmd.Flags().String("title", "", "Required element used in combination with CreateCase to provide a title for the new case.")
		securityIr_createCaseCmd.Flags().String("watchers", "", "Required element used in combination with CreateCase to provide a list of entities to receive notifications for case updates.")
		securityIr_createCaseCmd.MarkFlagRequired("description")
		securityIr_createCaseCmd.MarkFlagRequired("engagement-type")
		securityIr_createCaseCmd.MarkFlagRequired("impacted-accounts")
		securityIr_createCaseCmd.MarkFlagRequired("reported-incident-start-date")
		securityIr_createCaseCmd.MarkFlagRequired("resolver-type")
		securityIr_createCaseCmd.MarkFlagRequired("title")
		securityIr_createCaseCmd.MarkFlagRequired("watchers")
	})
	securityIrCmd.AddCommand(securityIr_createCaseCmd)
}
