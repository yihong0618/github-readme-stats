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
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/oauth2"
)

var (
	githubUserName string
	telegramID     int64
	telegramToken  string
	staredNumber   int
	reposNumber    int
	withStared     bool
	showAllPR      bool
	githubToken    string
)

var baseURL = "https://github.com/"
var myTitle = "# My GitHub Status\n"
var myCreatedTitle = "## The repos I created\n"
var myContributedTitle = "## The repos I contributed to\n"

func init() {
	flag.Int64Var(&telegramID, "tgid", 0, "telegram room id")
	flag.IntVar(&staredNumber, "stared", 10, "stared number random")
	flag.IntVar(&reposNumber, "repos", 0, "number of personal repos to show")
	flag.StringVar(&telegramToken, "tgtoken", "", "token from telegram")
	flag.StringVar(&githubUserName, "username", "", "github user name")
	flag.StringVar(&githubToken, "ghtoken", "", "token from github")
	flag.BoolVar(&withStared, "withstared", false, "if with stared repos")
	flag.BoolVar(&showAllPR, "showallpr", true, "if you want to show all prs included closed")
}

type myRepoInfo struct {
	star     int
	name     string
	HTMLURL  string
	create   string
	update   string
	language string
}

func (r *myRepoInfo) mdName() string {
	return "[" + r.name + "]" + "(" + r.HTMLURL + ")"
}

type myPrInfo struct {
	name      string
	repoURL   string
	firstDate string
	lastDate  string
	language  string
	prCount   int
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

func getRepoNameAndOwner(RepositoryURL string) (string, string) {
	q := strings.Split(RepositoryURL, "/")
	return q[len(q)-1], q[len(q)-2]
}

// 添加新的结构体来存储所有数据
type AllData struct {
	repos  []*github.Repository
	issues []*github.Issue
	stars  []*github.StarredRepository
}

func fetchAllCreatedRepos(username string, client *github.Client) []*github.Repository {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	var mutex sync.Mutex
	var wg sync.WaitGroup

	repos, resp, err := client.Repositories.List(context.Background(), username, opt)
	if err != nil {
		fmt.Println("Something wrong to get repos", err)
		return nil
	}
	allRepos = append(allRepos, repos...)

	totalPages := resp.LastPage
	for page := 2; page <= totalPages; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			opt.Page = pageNum
			repos, _, err := client.Repositories.List(context.Background(), username, opt)
			if err != nil {
				fmt.Println("Error fetching page", pageNum, err)
				return
			}
			mutex.Lock()
			allRepos = append(allRepos, repos...)
			mutex.Unlock()
		}(page)
	}
	wg.Wait()
	return allRepos
}

func fetchAllPrIssues(username string, client *github.Client) []*github.Issue {
	var allIssues []*github.Issue
	var mutex sync.Mutex
	var wg sync.WaitGroup

	filterContext := fmt.Sprintf("is:pr author:%s is:closed is:merged -user:%s", username, username)
	if showAllPR {
		filterContext = fmt.Sprintf("is:pr author:%s -user:%s", username, username)
	}

	semaphore := make(chan struct{}, 5)

	for page := 1; page <= 10; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			opt := &github.SearchOptions{ListOptions: github.ListOptions{Page: pageNum, PerPage: 100}}
			result, _, err := client.Search.Issues(context.Background(), filterContext, opt)
			if err != nil {
				fmt.Println("Error fetching issues page", pageNum, err)
				return
			}

			mutex.Lock()
			allIssues = append(allIssues, result.Issues...)
			mutex.Unlock()
		}(page)
	}
	wg.Wait()
	return allIssues
}

func fetchAllStared(username string, client *github.Client) []*github.StarredRepository {
	var allStared []*github.StarredRepository
	var mutex sync.Mutex
	var wg sync.WaitGroup

	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{Page: 1, PerPage: 100},
	}

	repos, resp, err := client.Activity.ListStarred(context.Background(), username, opt)
	if err != nil {
		fmt.Println("Something wrong to get stared")
		return nil
	}
	allStared = append(allStared, repos...)

	totalPages := resp.LastPage
	for page := 2; page <= totalPages; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			opt.Page = pageNum
			repos, _, err := client.Activity.ListStarred(context.Background(), username, opt)
			if err != nil {
				fmt.Println("Error fetching starred page", pageNum, err)
				return
			}
			mutex.Lock()
			allStared = append(allStared, repos...)
			mutex.Unlock()
		}(page)
	}
	wg.Wait()
	return allStared
}

func makeCreatedRepos(repos []*github.Repository) ([]myRepoInfo, int, int) {
	totalCount := 0
	longest := 0
	myRepos := []myRepoInfo{}
	for _, repo := range repos {

		// support fork if this fork stars > 5
		if !*repo.Fork || *repo.StargazersCount >= 5 {
			create := (*repo.CreatedAt).String()[:10]
			update := (*repo.UpdatedAt).String()[:10]
			language := "md"
			if repo.Language != nil {
				language = *repo.Language
			}
			myRepos = append(myRepos, myRepoInfo{
				star:     *repo.StargazersCount,
				name:     *repo.Name,
				create:   create,
				update:   update,
				language: language,
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

func makePrRepos(issues []*github.Issue, client *github.Client) ([]myPrInfo, int) {
	prMap := make(map[string]map[string]interface{})
	totalCount := 0
	for _, issue := range issues {
		if *issue.AuthorAssociation == "OWNER" {
			continue
		}
		repoName, owner := getRepoNameAndOwner(*issue.RepositoryURL)
		repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
		if err != nil {
			fmt.Println(repoName, "Something wrong to get repo language", err)
			continue
		}
		if *repo.Private == true {
			continue
		}
		if len(prMap[repoName]) == 0 {
			prMap[repoName] = make(map[string]interface{})
			prMap[repoName]["prCount"] = 1
			prMap[repoName]["firstDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["lastDate"] = (*issue.CreatedAt).String()[:10]
			prMap[repoName]["repoURL"] = *issue.RepositoryURL
			prMap[repoName]["firstHTML"] = *issue.HTMLURL
			prMap[repoName]["lastHTML"] = *issue.HTMLURL
			language := "md"
			if repo.Language != nil {
				language = *repo.Language
			}
			prMap[repoName]["language"] = language
		} else {
			prMap[repoName]["prCount"] = prMap[repoName]["prCount"].(int) + 1
			if prMap[repoName]["firstDate"].(string) > (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["firstDate"] = (*issue.CreatedAt).String()[:10]
				prMap[repoName]["firstHTML"] = *issue.HTMLURL
			}
			if prMap[repoName]["lastDate"].(string) < (*issue.CreatedAt).String()[:10] {
				prMap[repoName]["lastDate"] = (*issue.CreatedAt).String()[:10]
				prMap[repoName]["lastHTML"] = *issue.HTMLURL
			}
		}
		totalCount++
	}
	myPrs := []myPrInfo{}
	for k, v := range prMap {
		myPrs = append(myPrs, myPrInfo{
			name:      k,
			repoURL:   v["repoURL"].(string),
			firstDate: "[" + v["firstDate"].(string) + "]" + "(" + v["firstHTML"].(string) + ")", // markdown format -> []()
			lastDate:  "[" + v["lastDate"].(string) + "]" + "(" + v["lastHTML"].(string) + ")",   // markdown format -> []()
			language:  v["language"].(string),
			prCount:   v["prCount"].(int),
		})
	}
	return myPrs, totalCount
}

func makeStaredRepos(stars []*github.StarredRepository) []myStaredInfo {
	myStars := []myStaredInfo{}
	for _, star := range stars {
		repo := *star.Repository
		language := "md"
		if repo.Language != nil {
			language = *repo.Language
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
				language: language,
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

func makeCreatedString(repos []myRepoInfo, total int, reposNumber int) string {
	if reposNumber > 0 {
		repos = repos[:reposNumber]
	}
	starsData := [][]string{}
	for i, repo := range repos {
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), repo.create, repo.update, repo.language, strconv.Itoa(repo.star)})
	}
	starsData = append(starsData, []string{"sum", "", "", "", "", strconv.Itoa(total)})
	myStarsString := makeMdTable(starsData, []string{"ID", "Repo", "Start", "Update", "Language", "Stars"})
	return myCreatedTitle + myStarsString + "\n"
}

func makeContributedString(myPRs []myPrInfo, total int) string {
	prsData := [][]string{}
	for i, pr := range myPRs {
		prsData = append(prsData, []string{strconv.Itoa(i + 1), pr.mdName(), pr.firstDate, pr.lastDate, pr.language, fmt.Sprintf("[%d](%s)", pr.prCount, getAllPrLinks(pr))})
	}
	prsData = append(prsData, []string{"sum", "", "", "", "", strconv.Itoa(total)})
	myPrString := makeMdTable(prsData, []string{"ID", "Repo", "firstDate", "lastDate", "Language", "prCount"})
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
		starsData = append(starsData, []string{strconv.Itoa(i + 1), repo.mdName(), star.staredDate, repo.language, repo.update})
	}
	myStaredString := makeMdTable(starsData, []string{"ID", "Repo", "staredDate", "Language", "LatestUpdate"})
	return myStaredTitle + myStaredString + "\n"
}

func main() {
	flag.Parse()
	start_time := time.Now()
	defer func() {
		fmt.Println("Total time: ", time.Since(start_time))
	}()
	client := github.NewClient(nil)
	if githubToken != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
		ctx := context.Background()
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	var wg sync.WaitGroup
	data := &AllData{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		data.repos = fetchAllCreatedRepos(githubUserName, client)
	}()

	go func() {
		defer wg.Done()
		data.issues = fetchAllPrIssues(githubUserName, client)
	}()

	go func() {
		defer wg.Done()
		if withStared {
			data.stars = fetchAllStared(githubUserName, client)
		}
	}()

	wg.Wait()

	myRepos, totalStarsCount, longest := makeCreatedRepos(data.repos)
	sort.Slice(myRepos[:], func(i, j int) bool {
		return myRepos[j].star < myRepos[i].star
	})

	myPRs, totalPrCount := makePrRepos(data.issues, client)
	sort.Slice(myPRs[:], func(i, j int) bool {
		return myPRs[j].prCount < myPRs[i].prCount
	})
	myStaredString := ""
	if withStared {
		myStared := makeStaredRepos(data.stars)
		myStaredString = makeStaredString(myStared, staredNumber)
	}

	if telegramToken != "" {
		totalMessage := genTgMessage(myRepos, totalStarsCount, longest)
		send2Telegram(telegramToken, telegramID, totalMessage)
	}

	myCreatedString := makeCreatedString(myRepos, totalStarsCount, reposNumber)
	myPrString := makeContributedString(myPRs, totalPrCount)

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
