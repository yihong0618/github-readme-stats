package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/go-github/github"
)

var (
	githubUserName string
	telegramID     int64
	telegramToken  string
)

func init() {
	flag.Int64Var(&telegramID, "tgid", 0, "telegram room id")
	flag.StringVar(&telegramToken, "tgtoken", "", "token from telegram")
	flag.StringVar(&githubUserName, "username", "", "github user name")
}

type myRepoInfo struct {
	star int
	name string
}

func fetchAllRepos(username string) []*github.Repository {
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), username, opt)
		if err != nil {
			fmt.Println("wrong")
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos
}

func genTgMessage(myRepos []myRepoInfo, totalCount int, longest int) string {
	totalMessage := fmt.Sprintf("Total stars: %d", totalCount)
	totalMessage = totalMessage + "\n```"
	totalMessage = totalMessage + "\n"
	for _, repo := range myRepos {
		spaceTimes := longest - len(repo.name) + 4
		totalMessage = totalMessage + fmt.Sprintf("%s%s ---> %d \n", repo.name, strings.Repeat(" ", spaceTimes), repo.star)
	}
	totalMessage = totalMessage + "```\n"
	return totalMessage
}

func send2Telegram(tgToken string, roomID int64, message string) {
	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	ss := tgbotapi.NewMessage(roomID, message)
	ss.ParseMode = "MarkdownV2"
	bot.Send(ss)
}

func main() {
	flag.Parse()
	repos := fetchAllRepos(githubUserName)
	totalCount := 0
	longest := 0
	myRepos := []myRepoInfo{}
	for _, repo := range repos {
		if !*repo.Fork {
			myRepos = append(myRepos, myRepoInfo{star: *repo.StargazersCount, name: *repo.Name})
			totalCount = totalCount + *repo.StargazersCount
			if len(*repo.Name) > longest {
				longest = len(*repo.Name)
			}
		}
	}
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	totalMessage := genTgMessage(myRepos, totalCount, longest)
	send2Telegram(telegramToken, telegramID, totalMessage)
}
