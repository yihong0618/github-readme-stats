package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
)

var (
	githubUserName string
	telegramID     int64
	telegramToken  string
	staredNumber   int
	withStared     bool
)

var baseURL = "https://github.com/"
var myTitle = "# My GitHub Status\n"
var myCreatedTitle = "## The repos I created\n"
var myContributedTitle = "## The repos I contributed to\n"

func init() {
	flag.Int64Var(&telegramID, "tgid", 0, "telegram room id")
	flag.IntVar(&staredNumber, "stared", 10, "stared number random")
	flag.StringVar(&telegramToken, "tgtoken", "", "token from telegram")
	flag.StringVar(&githubUserName, "username", "", "github user name")
	flag.BoolVar(&withStared, "withstared", true, "if with stared repos")
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
	return "[" + p.name + "]" + "(" + p.repoLink() + ")"
}

func (p *myPrInfo) repoLink() string {
	q := strings.Split(p.repoURL, "/")
	return baseURL + q[len(q)-2] + "/" + q[len(q)-1]
}

func getAllPrLinks(p myPrInfo) string {
	url := fmt.Sprintf("%s/pulls?q=is:pr+author:%s", p.repoLink(), githubUserName)
	return "https://" + strings.ReplaceAll(strings.Split(url, "https://")[1], ":", "%3A")
}

type myStaredInfo struct {
	staredDate string
	desc       string
	myRepoInfo
}

func getRepoName(RepositoryURL string) string {
	q := strings.Split(RepositoryURL, "/")
	return q[len(q)-1]
}

func fetchAllCreatedRepos(username string, client *github.Client) []*github.Repository {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), username, opt)
		if err != nil {
			fmt.Println("Something wrong to get repos")
			continue
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos
}

func makeCreatedRepos(repos []*github.Repository) ([]myRepoInfo, int, int) {
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
	return myRepos, totalCount, longest
}

func fetchAllPrIssues(username string, client *github.Client) []*github.Issue {
	nowPage := 100
	opt := &github.SearchOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	var allIssues []*github.Issue
	for {
		result, _, err := client.Search.Issues(context.Background(), fmt.Sprintf("is:pr author:%s is:closed is:merged", username), opt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allIssues = append(allIssues, result.Issues...)
		if nowPage >= result.GetTotal() {
			break
		}
		opt.Page = opt.Page + 1
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

func fetchAllStared(username string, client *github.Client) []*github.StarredRepository {
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{Page: 1, PerPage: 100},
	}
	var allStared []*github.StarredRepository
	for {
		repos, resp, err := client.Activity.ListStarred(context.Background(), username, opt)
		if err != nil {
			fmt.Println("Something wrong to get stared")
		}
		allStared = append(allStared, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allStared
}

func makeStaredRepos(stars []*github.StarredRepository) []myStaredInfo {
	myStars := []myStaredInfo{}
	for _, star := range stars {
		repo := *star.Repository
		lauguage := "md"
		if repo.Language != nil {
			lauguage = *repo.Language
		}
		desc := ""
		if repo.Description != nil {
			desc = *repo.Description
		}

		myStars = append(myStars, myStaredInfo{
			staredDate: (*star.StarredAt).String()[:10],
			desc:       desc,
			myRepoInfo: myRepoInfo{
				name:     *repo.Name,
				create:   (*repo.CreatedAt).String()[:10],
				update:   (*repo.UpdatedAt).String()[:10],
				lauguage: lauguage,
				HTMLURL:  *repo.HTMLURL,
			},
		})
	}
	// shffle to get random array
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(myStars), func(i, j int) { myStars[i], myStars[j] = myStars[j], myStars[i] })
	return myStars
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

func send2Telegram(tgToken string, roomID int64, message string) bool {
	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return false
	}

	s := tgbotapi.NewMessage(roomID, message)
	s.ParseMode = "MarkdownV2"
	bot.Send(s)
	return true
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

func makeCreatedString(repos []myRepoInfo) string {
	starsData := [][]string{}
	for i, repo := range repos {
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), repo.create, repo.update, repo.lauguage, strconv.Itoa(repo.star)})
	}
	myStarsString := makeMdTable(starsData, []string{"ID", "Repo", "Start", "Update", "Lauguage", "Stars"})
	return myCreatedTitle + myStarsString + "\n"
}

func makeContributedString(myPRs []myPrInfo) string {
	prsData := [][]string{}
	for i, pr := range myPRs {
		prsData = append(prsData, []string{strconv.Itoa(i + 1), pr.mdName(), pr.fisrstDate, pr.lasteDate, fmt.Sprintf("[%d](%s)", pr.prCount, getAllPrLinks(pr))})
	}
	myPrString := makeMdTable(prsData, []string{"ID", "Repo", "firstDate", "lasteDate", "prCount"})
	return myContributedTitle + myPrString + "\n"
}

func makeStaredString(myStars []myStaredInfo, starNumber int) string {
	myStaredTitle := fmt.Sprintf("## The repos I stared (random %s)", strconv.Itoa(starNumber)) + "\n"
	starsData := [][]string{}
	// maybe a better way in golang?
	if (len(myStars)) < starNumber {
		starNumber = len(myStars)
	}
	for i, star := range myStars[:starNumber] {
		repo := star.myRepoInfo
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), star.staredDate, repo.lauguage, repo.update})
	}
	myStaredString := makeMdTable(starsData, []string{"ID", "Repo", "staredDate", "Lauguage", "LatestUpdate"})
	return myStaredTitle + myStaredString + "\n"
}

func main() {
	flag.Parse()
	client := github.NewClient(nil)
	repos := fetchAllCreatedRepos(githubUserName, client)
	myRepos, totalCount, longest := makeCreatedRepos(repos)
	// change sort logic here
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	issues := fetchAllPrIssues(githubUserName, client)
	myPRs := makePrRepos(issues)
	// change sort logic here
	sort.Slice(myPRs[:], func(i, j int) bool {
		return myPRs[j].prCount < myPRs[i].prCount
	})
	myStaredString := ""
	if withStared {
		stars := fetchAllStared(githubUserName, client)
		myStared := makeStaredRepos(stars)
		myStaredString = makeStaredString(myStared, staredNumber)
	}

	if telegramToken != "" {
		totalMessage := genTgMessage(myRepos, totalCount, longest)
		send2Telegram(telegramToken, telegramID, totalMessage)
	}

	myCreatedString := makeCreatedString(myRepos)
	myPrString := makeContributedString(myPRs)

	readMeFile := path.Join(os.Getenv("GITHUB_WORKSPACE"), "README.md")
	readMeContent, err := ioutil.ReadFile(readMeFile)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`(?s)(<!--START_SECTION:my_github-->)(.*)(<!--END_SECTION:my_github-->)`)
	newContentString := myCreatedString + myPrString
	if withStared {
		newContentString = newContentString + myStaredString
	}
	newContent := []byte(re.ReplaceAllString(string(readMeContent), `$1`+"\n"+newContentString+`$3`))
	err = ioutil.WriteFile(readMeFile, newContent, 0644)
	if err != nil {
		panic(err)
	}
}
