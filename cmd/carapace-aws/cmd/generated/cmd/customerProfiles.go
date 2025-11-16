package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfilesCmd = &cobra.Command{
	Use:   "customer-profiles",
	Short: "Amazon Connect Customer Profiles\n\n- [Customer Profiles actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Customer_Profiles.html)\n- [Customer Profiles data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Connect_Customer_Profiles.html)\n\nAmazon Connect Customer Profiles is a unified customer profile for your contact center that has pre-built connectors powered by AppFlow that make it easy to combine customer information from third party applications, such as Salesforce (CRM), ServiceNow (ITSM), and your enterprise resource planning (ERP), with contact history from your Amazon Connect contact center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfilesCmd).Standalone()

	})
	rootCmd.AddCommand(customerProfilesCmd)
}
