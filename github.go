package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
)

var (
	githubUserName string
	telegramID     int64
	telegramToken  string
)

var baseURL = "https://github.com/"
var myTitle = "# My GitHub Status\n"
var myStatsChartTitle = `<img align="middle" src="https://github-readme-stats-1.yihong0618.vercel.app/api?username=yihong0618&show_icons=true&&&hide_title=true" />`

func init() {
	flag.Int64Var(&telegramID, "tgid", 0, "telegram room id")
	flag.StringVar(&telegramToken, "tgtoken", "", "token from telegram")
	flag.StringVar(&githubUserName, "username", "", "github user name")
}

type myRepoInfo struct {
	star     int
	name     string
	HTMLURL  string
	create   string
	update   string
	lauguage string
}

func (r *myRepoInfo) mdName() string {
	return "[" + r.name + "]" + "(" + r.HTMLURL + ")"
}

type myPrInfo struct {
	name       string
	repoURL    string
	fisrstDate string
	lasteDate  string
	prCount    int
}

func (p *myPrInfo) mdName() string {
	q := strings.Split(p.repoURL, "/")
	// md format
	return "[" + p.name + "]" + "(" + baseURL + q[len(q)-2] + "/" + q[len(q)-1] + ")"
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

func getRepoName(RepositoryURL string) string {
	q := strings.Split(RepositoryURL, "/")
	return q[len(q)-1]
}

func fetchAllPrIssues(username string) []*github.Issue {
	client := github.NewClient(nil)
	opt := &github.SearchOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	nowPage := 100
	var allIssues []*github.Issue
	for {
		result, _, err := client.Search.Issues(context.Background(), fmt.Sprintf("is:pr author:%s is:closed is:merged", username), opt)
		if err != nil {
			fmt.Println(err)
		}
		allIssues = append(allIssues, result.Issues...)
		if nowPage >= result.GetTotal() {
			break
		}
		nowPage = nowPage + 100
	}
	return allIssues
}

func makePrRepos(issues []*github.Issue) []myPrInfo {
	prMap := make(map[string]map[string]interface{})
	for _, issue := range issues {
		if *issue.AuthorAssociation == "OWNER" {
			continue
		}
		repoName := getRepoName(*issue.RepositoryURL)
		if len(prMap[repoName]) == 0 {
			prMap[repoName] = make(map[string]interface{})
			prMap[repoName]["prCount"] = 1
			prMap[repoName]["fisrstDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["lasteDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["repoURL"] = *issue.RepositoryURL
		} else {
			prMap[repoName]["prCount"] = prMap[repoName]["prCount"].(int) + 1
			if prMap[repoName]["fisrstDate"].(string) > (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["fisrstDate"] = (*issue.CreatedAt).String()[:10]
			}
			if prMap[repoName]["lasteDate"].(string) < (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["lasteDate"] = (*issue.CreatedAt).String()[:10]
			}
		}
	}
	myPrs := []myPrInfo{}
	for k, v := range prMap {
		myPrs = append(myPrs, myPrInfo{
			name:       k,
			repoURL:    v["repoURL"].(string),
			fisrstDate: v["fisrstDate"].(string),
			lasteDate:  v["lasteDate"].(string),
			prCount:    v["prCount"].(int),
		})
	}
	return myPrs
}

func genTgMessage(myRepos []myRepoInfo, totalCount int, longest int) string {
	totalMessage := fmt.Sprintf("Total stars: %d", totalCount)
	totalMessage = totalMessage + "\n```"
	totalMessage = totalMessage + "\n"
	for _, repo := range myRepos {
		spaceTimes := longest - len(repo.name) + 2
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

	s := tgbotapi.NewMessage(roomID, message)
	s.ParseMode = "MarkdownV2"
	bot.Send(s)
}

func makeMdTable(data [][]string, header []string) string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return tableString.String()
}

func main() {
	flag.Parse()
	repos := fetchAllRepos(githubUserName)
	totalCount := 0
	longest := 0
	myRepos := []myRepoInfo{}
	for _, repo := range repos {
		if !*repo.Fork {
			create := (*repo.CreatedAt).String()[:10]
			update := (*repo.UpdatedAt).String()[:10]
			lauguage := "md"
			if repo.Language != nil {
				lauguage = *repo.Language
			}
			myRepos = append(myRepos, myRepoInfo{
				star:     *repo.StargazersCount,
				name:     *repo.Name,
				create:   create,
				update:   update,
				lauguage: lauguage,
				HTMLURL:  *repo.HTMLURL,
			})
			totalCount = totalCount + *repo.StargazersCount
			if len(*repo.Name) > longest {
				longest = len(*repo.Name)
			}
		}
	}
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	issues := fetchAllPrIssues(githubUserName)
	myPRs := makePrRepos(issues)
	sort.Slice(myPRs[:], func(i, j int) bool {
		return myPRs[j].prCount < myPRs[i].prCount
	})

	fmt.Println(myPRs)
	// totalMessage := genTgMessage(myRepos, totalCount, longest)
	// send2Telegram(telegramToken, telegramID, totalMessage)
	starsData := [][]string{}
	for _, repo := range myRepos {
		starsData = append(starsData, []string{repo.mdName(), repo.create, repo.update, repo.lauguage, strconv.Itoa(repo.star)})
	}
	myStarsString := makeMdTable(starsData, []string{"Repo", "Start", "Update", "Lauguage", "Stars"})

	prsData := [][]string{}
	for _, pr := range myPRs {
		prsData = append(prsData, []string{pr.mdName(), pr.fisrstDate, pr.lasteDate, strconv.Itoa(pr.prCount)})
	}
	myPrString := makeMdTable(prsData, []string{"Repo", "firstDate", "lasteDate", "prCount"})

	content := []byte(myTitle + "\n" + myStarsString + "\n" + "\n" + "# My PRs\n" + "\n" + myPrString)
	err := ioutil.WriteFile("test.md", content, 0644)
	fmt.Println(myStatsChartTitle)
	if err != nil {
		panic(err)
	}
}
