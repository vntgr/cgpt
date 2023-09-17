package cmd

import (
	"fmt"
	"os"
	"strings"

	"git.mysticmode.net/mysticmode/kily/pkg/chatgpt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// messengerCmd represents the messenger command
var messengerCmd = &cobra.Command{
	Use:   "messenger <message>",
	Short: "Send a message to ChatGPT",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputMsg := strings.Join(args, " ")

		openAIApiKey := viper.GetString("openai.api_key")
		if openAIApiKey == "" {
			pterm.Error.Println("openai.com API key not set. Please run 'kily init' to set it.")
			os.Exit(1)
		}

		client := chatgpt.NewChatGPTClient(openAIApiKey)
		req := &chatgpt.ChatGPTRequest{
			Messages: []chatgpt.RequesterData{
				{
					Content: inputMsg,
					Role:    "user",
				},
			},
		}

		data, err := chatgpt.PostMessage(client, req)
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(messengerCmd)
}
