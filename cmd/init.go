package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the kily CLI configuration",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Println("initializing kily configuration")

		// Set the openai.com API key
		openAIApiKey, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your OpenAI API key")
		if err != nil {
			pterm.Error.Println("error parsing OpenAI API key")
		}
		viper.Set("openai.api_key", openAIApiKey)

		// Save the configuration to kily.yml
		viper.WriteConfigAs(viper.ConfigFileUsed())
		if err != nil {
			pterm.Error.Println("error writing kily configuration")
		}

		pterm.Println()
		pterm.Success.Println("kily configuration initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
