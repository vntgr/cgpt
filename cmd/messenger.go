package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"git.mysticmode.net/mysticmode/kily/pkg/chatgpt"
	"git.mysticmode.net/mysticmode/kily/pkg/kvstore"
)

var kv kvstore.KVOps

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

		requesterData := []chatgpt.RequesterData{
			{
				Role:    "user",
				Content: inputMsg,
			},
		}

		client := chatgpt.NewChatGPTClient(openAIApiKey)
		req := &chatgpt.ChatGPTRequest{
			Messages: requesterData,
		}

		data, err := chatgpt.PostMessage(client, req)
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		fmt.Println(data.Choices[0].Message.Content)

		err = kv.Put([]byte(data.Choices[0].Message.Role), []byte(data.Choices[0].Message.Content))
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(messengerCmd)
	kv, _ = kvstore.InitializeKVStore()
}
